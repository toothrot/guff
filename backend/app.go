package main

import (
	"context"
	"log"
	"mime"
	"net"
	"net/http"
	"os"
	"path"
	"path/filepath"
	"time"

	"github.com/golang/glog"
	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"github.com/grpc-ecosystem/go-grpc-middleware"
	"github.com/improbable-eng/grpc-web/go/grpcweb"
	"google.golang.org/grpc"

	"github.com/toothrot/guff/backend/auth"
	"github.com/toothrot/guff/backend/core"
	"github.com/toothrot/guff/backend/generated"
	"github.com/toothrot/guff/backend/services"

	_ "google.golang.org/grpc/grpclog/glogger"
)

type guffApp struct {
	Config *core.Config

	router     *mux.Router
	server     *http.Server
	grpcServer *grpc.Server
	grpcWeb    *grpcweb.WrappedGrpcServer
}

func (g *guffApp) Serve(ctx context.Context) {
	g.router = mux.NewRouter()

	am, err := auth.NewAuthMiddleware(ctx, g.Config.OAuthConfig)
	if err != nil {
		glog.Fatalf("Error creating auth middleware: %q", err)
	}

	g.server = &http.Server{Addr: net.JoinHostPort("", *port), Handler: handlers.CompressHandler(g.router)}
	g.grpcServer = grpc.NewServer(grpc.UnaryInterceptor(grpc_middleware.ChainUnaryServer(am.ServerInterceptor)))
	guff_proto.RegisterUsersServiceServer(g.grpcServer, &services.Users{Config: g.Config})
	g.grpcWeb = grpcweb.WrapServer(g.grpcServer)

	g.registerRoutes()

	go func() {
		glog.Infof("Listening on port %q", *port)
		if err := g.server.ListenAndServe(); err != http.ErrServerClosed {
			glog.Errorf("http.ListenAndServe() = %q", err)
		}
	}()
	<-ctx.Done()

	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()

	if err := g.server.Shutdown(ctx); err != nil {
		glog.Errorf("Error shutting down http server: %q", err)
	}
}

func (g *guffApp) registerRoutes() {
	g.router.PathPrefix("/guff.proto").Handler(g.grpcWeb)
	g.router.PathPrefix("/").Handler(&fileServer{*webRoot})
}

type fileServer struct {
	webRoot string
}

func (f *fileServer) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	abs, err := filepath.Abs(f.webRoot)
	if err != nil {
		log.Fatalf("Error parsing absolute path from %q", *webRoot)
	}
	glog.Infof("Serving from root %q", abs)
	if _, err := os.Stat(path.Join(abs, r.URL.Path)); os.IsNotExist(err) {
		// Fall-back to angular's router.
		glog.Infof("welp %q %q %q", r.URL.Path, err, path.Join(abs, "/index.html"))
		http.ServeFile(w, r, path.Join(abs, "/index.html"))
		return
	}

	glog.Infof("%q: %q", r.Method, r.URL.Path)
	w.Header().Set("Content-Type", mime.TypeByExtension(path.Ext(r.URL.Path)))
	w.Header().Set("Cache-Control", "no-cache, private, max-age=0")

	fs := http.FileServer(http.Dir(abs))
	fs.ServeHTTP(w, r)
}
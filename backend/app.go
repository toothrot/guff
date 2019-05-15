package main

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"mime"
	"net"
	"net/http"
	"net/url"
	"os"
	"path"
	"path/filepath"
	"strings"
	"time"

	"github.com/golang/glog"
	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	grpc_middleware "github.com/grpc-ecosystem/go-grpc-middleware"
	"github.com/improbable-eng/grpc-web/go/grpcweb"
	"google.golang.org/grpc"

	_ "google.golang.org/grpc/grpclog/glogger"

	"github.com/toothrot/guff/backend/auth"
	"github.com/toothrot/guff/backend/core"
	guff_proto "github.com/toothrot/guff/backend/generated"
	"github.com/toothrot/guff/backend/models"
	"github.com/toothrot/guff/backend/services"
)

func newGuffApp(ctx context.Context, config *core.Config) *guffApp {
	g := &guffApp{config: config}
	g.router = mux.NewRouter()
	p := &models.DBPersist{DB: setupDB(config)}
	am, err := auth.NewMiddleware(ctx, g.config.OAuthConfig, p)
	if err != nil {
		glog.Fatalf("Error creating auth middleware: %q", err)
	}
	g.server = &http.Server{Addr: net.JoinHostPort("", *port), Handler: handlers.CompressHandler(g.router)}
	g.grpcServer = grpc.NewServer(grpc.UnaryInterceptor(grpc_middleware.ChainUnaryServer(am.ServerInterceptor)))
	guff_proto.RegisterUsersServiceServer(g.grpcServer, &services.Users{Config: g.config})
	guff_proto.RegisterAdminServiceServer(g.grpcServer, &services.Admin{Config: g.config, Persist: p})
	guff_proto.RegisterDivisionsServiceServer(g.grpcServer, &services.Divisions{Persist: p})
	g.grpcWeb = grpcweb.WrapServer(g.grpcServer)
	g.registerRoutes()
	return g
}

func setupDB(config *core.Config) *sql.DB {
	connstring := fmt.Sprintf("dbname=%s password=%s", config.DBName, config.DBPassword)
	if config.DBURL != "" {
		connstring = strings.TrimSpace(config.DBURL)
	}
	db, err := sql.Open("postgres", connstring)
	if err != nil {
		glog.Fatalf("sql.Open(%q, %q) = _, %v, wanted no error", "postgres", fmt.Sprintf("dbname=%q", config.DBName), err)
	}
	if err := models.Migrate(db); err != nil {
		glog.Fatalf("Migrate(%v) = %v, wanted no error", db, err)
	}
	return db
}

type guffApp struct {
	config *core.Config

	context    context.Context
	router     *mux.Router
	server     *http.Server
	grpcServer *grpc.Server
	grpcWeb    *grpcweb.WrappedGrpcServer
}

func (g *guffApp) Serve(ctx context.Context) {
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
	if g.config.RequireHTTPS != "" {
		g.router.Use(httpsMiddleware)
	}
}

type fileServer struct {
	webRoot string
}

func (f *fileServer) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	glog.Infof("h: %#v", r.Header)
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

var xForwardedProto = http.CanonicalHeaderKey("X-Forwarded-Proto")

func httpsMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if strings.ToLower(r.Header.Get(xForwardedProto)) != "https" {
			u := &url.URL{
				Scheme:   "https",
				Host:     r.Host,
				Path:     r.URL.Path,
				RawQuery: r.URL.RawQuery,
			}
			http.Redirect(w, r, u.String(), http.StatusFound)
			return
		}
		next.ServeHTTP(w, r)
	})
}

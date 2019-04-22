package main

import (
	"context"
	"encoding/json"
	"log"
	"mime"
	"net"
	"net/http"
	"net/url"
	"path"
	"path/filepath"
	"time"

	"github.com/golang/glog"
	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"github.com/improbable-eng/grpc-web/go/grpcweb"
	"google.golang.org/api/oauth2/v2"
	"google.golang.org/grpc"

	"github.com/toothrot/guff/go/core"
	"github.com/toothrot/guff/go/generated"
	"github.com/toothrot/guff/go/services"

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

	g.server = &http.Server{Addr: net.JoinHostPort("", *port), Handler: handlers.CompressHandler(g.router)}
	g.grpcServer = grpc.NewServer()
	guff_proto.RegisterUsersServiceServer(g.grpcServer, &services.UsersService{Config: g.Config})
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
	g.router.PathPrefix("/login").HandlerFunc(g.loginHandler)
	g.router.PathPrefix("/oauth2callback").HandlerFunc(g.oauth2callback)
	g.router.PathPrefix("/").Handler(&fileServer{*webRoot})
}

func (g *guffApp) loginHandler(w http.ResponseWriter, r *http.Request) {
	v := url.Values{}
	v.Set("scope", "email profile")
	v.Set("client_id", g.Config.OAuthConfig.ClientID)
	v.Set("redirect_uri", "http://localhost:8080/oauth2callback")
	v.Set("response_type", "code")

	uri := url.URL{
		Scheme:   "https",
		Host:     "accounts.google.com",
		Path:     "/o/oauth2/v2/auth",
		RawQuery: v.Encode(),
	}
	http.Redirect(w, r, uri.String(), http.StatusFound)
}

func (g *guffApp) oauth2callback(w http.ResponseWriter, r *http.Request) {
	if err := r.ParseForm(); err != nil {
		glog.Errorf("parseform: %q", err)
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	token, err := g.Config.OAuthConfig.Exchange(context.Background(), r.Form.Get("code"))
	if err != nil {
		glog.Errorf("exchange: %q", err)
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}
	resp, err := http.Get("https://www.googleapis.com/oauth2/v2/userinfo?access_token=" + token.AccessToken)
	if err != nil {
		glog.Errorf("userinfo: %q", err)
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	d := json.NewDecoder(resp.Body)
	user := &oauth2.Userinfoplus{}
	if err := d.Decode(user); err != nil {
		glog.Errorf("decode: %q", err)
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	session, err := g.Config.CookieStore.Get(r, "guff")
	if err != nil {
		glog.Errorf("cookiestore get error %q", err)
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	session.Values["email"] = user.Email
	if err := session.Save(r, w); err != nil {
		glog.Errorf("session save: %q", err)
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	http.Redirect(w, r, "/", http.StatusFound)
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

	glog.Infof("%q: %q", r.Method, r.URL.Path)
	w.Header().Set("Content-Type", mime.TypeByExtension(path.Ext(r.URL.Path)))
	w.Header().Set("Cache-Control", "no-cache, private, max-age=0")

	fs := http.FileServer(http.Dir(abs))
	fs.ServeHTTP(w, r)
}

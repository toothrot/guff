package main

import (
  "context"
  "flag"
  "log"
  "mime"
  "net"
  "net/http"
  "os"
  "os/signal"
  "path"
  "path/filepath"
  "syscall"
  "time"

  "github.com/golang/glog"
  "github.com/gorilla/handlers"
  "github.com/gorilla/mux"
)

var (
	webRoot = flag.String("web_root", "", "Path from which to serve web files.")
  port = flag.String("port", "8080", "Port to listen for HTTP requests.")
)

func main() {
  defer glog.Info("later, gator.")
	flag.Parse()
	glog.Info("Don't take any guff from these swine.")

  ctx, cancel := context.WithCancel(context.Background())
  defer cancel()
  ctx = handleSigs(ctx)

	r := mux.NewRouter()
	r.PathPrefix("/").HandlerFunc(fileServerFunc(*webRoot))

	s := &http.Server{Addr: net.JoinHostPort("", *port), Handler: handlers.CompressHandler(r)}

	go func() {
    glog.Infof("Listening on port %q", *port)
    if err :=  s.ListenAndServe(); err != http.ErrServerClosed {
      glog.Errorf("http.ListenAndServe() = %q", err)
    }
  }()
  <-ctx.Done()

  ctx, cancel = context.WithTimeout(context.Background(), 2 * time.Second)
  defer cancel()

  if err := s.Shutdown(ctx); err != nil {
    glog.Errorf("Error shutting down http server: %q", err)
  }
}

func fileServerFunc(root string) http.HandlerFunc {
	abs, err := filepath.Abs(root)
	if err != nil {
		log.Fatalf("Error parsing absolute path from %q", *webRoot)
	}
	glog.Infof("Serving from root %q", abs)

	return func(w http.ResponseWriter, r *http.Request) {
		glog.Infof("%q: %q", r.Method, r.URL.Path)
		w.Header().Set("Content-Type", mime.TypeByExtension(path.Ext(r.URL.Path)))
		fs := http.FileServer(http.Dir(abs))
		fs.ServeHTTP(w, r)
	}
}

func handleSigs(ctx context.Context) context.Context {
  sigint := make(chan os.Signal, 1)
  signal.Notify(sigint, os.Interrupt)
  signal.Notify(sigint, syscall.SIGTERM)

  ctx, cancel := context.WithCancel(ctx)

  go func() {
    select {
    case s := <-sigint:
      glog.Infof("Handling signal %#v", s)
      cancel()
    case <-ctx.Done():
    }
  }()

  return ctx
}

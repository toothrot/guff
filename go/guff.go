package main

import (
	"flag"
	"log"
	"mime"
	"net/http"
	"path"
	"path/filepath"

	"github.com/golang/glog"
	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
)

var (
	webRoot = flag.String("web_root", "", "Path from which to serve web files.")
)

func main() {
	flag.Parse()
	glog.Info("Don't take any guff from these swine.")

	r := mux.NewRouter()
	r.PathPrefix("/").HandlerFunc(fileServerFunc(*webRoot))
	err := http.ListenAndServe(":8080", handlers.CompressHandler(r))
	if err != nil {
		glog.Errorf("http.ListenAndServe() = %q", err)
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

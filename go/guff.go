package main

import (
	"context"
	"flag"
	"os"
	"os/signal"
	"syscall"

	"github.com/golang/glog"
)

var (
	webRoot = flag.String("web_root", "", "Path from which to serve web files.")
	port    = flag.String("port", "8080", "Port to listen for HTTP requests.")
)

func main() {
	defer glog.Info("later, gator.")
	flag.Parse()
	glog.Info("Don't take any guff from these swine.")

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	ctx = handleSigs(ctx)

	g := guffApp{}
	g.Serve(ctx)
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

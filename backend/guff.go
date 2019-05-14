package main

import (
	"context"
	"flag"
	"io/ioutil"
	"os"
	"os/signal"
	"syscall"

	"github.com/golang/glog"
	"github.com/gorilla/sessions"
	"golang.org/x/oauth2/google"

	"github.com/toothrot/guff/backend/core"
)

var (
	webRoot         = flag.String("web_root", "", "Path from which to serve web files.")
	port            = flag.String("port", "8080", "Port to listen for HTTP requests.")
	sessionKeyPath  = flag.String("session_key", "/run/secrets/session-key", "Session key secret file path")
	oauthConfigPath = flag.String("oauth_config", "/run/secrets/oauth.json", "OAuth config JSON file path (see http://golang.org/x/oauth2/google#ConfigFromJSON)")
	dbPassPath      = flag.String("db_pass_path", "/run/secrets/postgres-guff-password", "Database password secret file path")
	divisionsURL    = flag.String("divisions_url", "https://royalpalmsshuffle.leagueapps.com/leagues?state=LIVE&locationId=&seasonId=&days=&levelId=", "URL of Divisions page")
	dbname          = flag.String("dbname", "guff_dev", "Postgres database name")
)

func main() {
	defer glog.Info("later, gator.")
	flag.Parse()
	glog.Info("Don't take any guff from these swine.")

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	ctx = handleSigs(ctx)

	b, err := getSecret(*sessionKeyPath)
	if err != nil {
		glog.Fatalf("error getting secret %q: %q", *sessionKeyPath, err)
	}
	store := sessions.NewCookieStore(b)

	b, err = getSecret(*oauthConfigPath)
	if err != nil {
		glog.Fatalf("error getting secret %q: %q", *oauthConfigPath, err)
	}
	oc, err := google.ConfigFromJSON(b)
	if err != nil {
		glog.Errorf("google.ConfigFromJSON() returned error %q", err)
	}
	dbpass, err := getSecret(*dbPassPath)
	if err != nil {
		glog.Fatalf("error getting secret %q: %q", *oauthConfigPath, err)
	}
	oc.Scopes = []string{"email", "profile"}
	conf := &core.Config{
		OAuthConfig: oc,
		CookieStore: store,
		ProgramsURL: *divisionsURL,
		DBName:      *dbname,
		DBPassword:  string(dbpass),
	}
	g := newGuffApp(ctx, conf)

	g.Serve(ctx)
}

func getSecret(path string) ([]byte, error) {
	f, err := os.Open(path)
	if err != nil {
		return []byte{}, err
	}
	return ioutil.ReadAll(f)
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

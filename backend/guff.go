package main

import (
	"context"
	"encoding/base64"
	"flag"
	"io/ioutil"
	"os"
	"os/signal"
	"syscall"

	"cloud.google.com/go/kms/apiv1"
	"github.com/golang/glog"
	"github.com/gorilla/sessions"
	"golang.org/x/oauth2/google"
	kmspb "google.golang.org/genproto/googleapis/cloud/kms/v1"

	"github.com/toothrot/guff/backend/core"
)

var (
	webRoot      = flag.String("web_root", "", "Path from which to serve web files.")
	port         = flag.String("port", "8080", "Port to listen for HTTP requests.")
	divisionsURL = flag.String("divisions_url", "https://royalpalmsshuffle.leagueapps.com/leagues?state=LIVE&locationId=&seasonId=&days=&levelId=", "URL of Divisions page")
	dbname       = flag.String("dbname", "guff_dev", "Postgres database name")

	// Secrets
	sessionKeyPath  = flag.String("session_key", "/run/secrets/session-key", "Session key secret file path")
	oauthConfigPath = flag.String("oauth_config", "/run/secrets/oauth.json", "OAuth config JSON file path (see http://golang.org/x/oauth2/google#ConfigFromJSON)")
	dbPassPath      = flag.String("db_pass_path", "/run/secrets/postgres-guff-password", "Database password secret file path")

	// ENV Secrets
	// example keyName: "projects/PROJECT_ID/locations/global/keyRings/RING_ID/cryptoKeys/KEY_ID"
	kmsKey         = os.Getenv("GUFF_KMS_KEY")
	sessionKeyEnc  = os.Getenv("GUFF_SESSION_KEY_ENC")
	oauthConfigEnc = os.Getenv("GUFF_OAUTH_CONFIG_ENC")
	dbURLEnc       = os.Getenv("GUFF_DB_URL_ENC")

	requireHTTPS = os.Getenv("GUFF_REQUIRE_HTTPS")
)

func main() {
	defer glog.Info("later, gator.")
	flag.Parse()
	glog.Info("Don't take any guff from these swine.")

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	ctx = handleSigs(ctx)

	if lis := os.Getenv("PORT"); lis != "" {
		*port = lis
		glog.Infof("Will listen on %q", *port)
	}

	kc, err := kms.NewKeyManagementClient(ctx)
	if err != nil {
		glog.Errorf("kms.NewKeyManagementClient() = %v", err)
	}
	store := sessions.NewCookieStore(sessionKey(ctx, kc))

	oc, err := google.ConfigFromJSON(oauthConfig(ctx, kc))
	if err != nil {
		glog.Errorf("google.ConfigFromJSON() returned error %q", err)
	}

	dbpass, err := getSecret(*dbPassPath)
	if err != nil {
		glog.Errorf("error getting secret %q: %q", *dbPassPath, err)
		dbpass = []byte{}
	}
	oc.Scopes = []string{"email", "profile"}
	conf := &core.Config{
		OAuthConfig:  oc,
		CookieStore:  store,
		ProgramsURL:  *divisionsURL,
		DBName:       *dbname,
		DBPassword:   string(dbpass),
		DBURL:        string(getKMSSecret(ctx, kc, dbURLEnc)),
		RequireHTTPS: requireHTTPS,
	}
	g := newGuffApp(ctx, conf)

	g.Serve(ctx)
}

func sessionKey(ctx context.Context, kc *kms.KeyManagementClient) []byte {
	if sessionKeyEnc != "" {
		return getKMSSecret(ctx, kc, sessionKeyEnc)
	}
	b, err := getSecret(*sessionKeyPath)
	if err != nil {
		glog.Fatalf("error getting secret %q: %q", *sessionKeyPath, err)
	}
	return b
}

func oauthConfig(ctx context.Context, kc *kms.KeyManagementClient) []byte {
	if oauthConfigEnc != "" {
		return getKMSSecret(ctx, kc, oauthConfigEnc)
	}
	b, err := getSecret(*oauthConfigPath)
	if err != nil {
		glog.Fatalf("error getting secret %q: %q", *oauthConfigPath, err)
	}
	return b
}

func getKMSSecret(ctx context.Context, kc *kms.KeyManagementClient, cipherText string) []byte {
	if kmsKey == "" {
		glog.Infof("GUFF_KMS_KEY not set, skipping")
		return []byte{}
	}
	b, err := base64.StdEncoding.DecodeString(cipherText)
	if err != nil {
		glog.Errorf("error decoding secret: %v", err)
		return []byte{}
	}
	req := &kmspb.DecryptRequest{Name: kmsKey, Ciphertext: b}
	resp, err := kc.Decrypt(ctx, req)
	if err != nil {
		glog.Errorf("error decrypting secret: %v", err)
		return []byte{}
	}
	return resp.GetPlaintext()
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

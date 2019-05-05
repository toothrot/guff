package test

import (
	"crypto/rand"
	"crypto/rsa"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/coreos/go-oidc"
	"golang.org/x/oauth2"
	"gopkg.in/square/go-jose.v2"
)

type ProviderJSON struct {
	Issuer      string `json:"issuer"`
	AuthURL     string `json:"authorization_endpoint"`
	TokenURL    string `json:"token_endpoint"`
	JWKSURL     string `json:"jwks_uri"`
	UserInfoURL string `json:"userinfo_endpoint"`
}

type IdToken struct {
	Issuer   string   `json:"iss"`
	Subject  string   `json:"sub"`
	Audience []string `json:"aud"`
	Expiry   int64    `json:"exp"`
	IssuedAt int64    `json:"iat"`
	Nonce    string   `json:"nonce"`
	AtHash   string   `json:"at_hash"`
	oidc.UserInfo
}

type KeyServer struct {
	t        *testing.T
	URL      string
	provider ProviderJSON
}

func (k *KeyServer) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	k.t.Helper()
	var resp interface{}
	switch r.URL.Path {
	case "/.well-known/openid-configuration":
		resp = k.provider
	default:
		resp = getKeySet()
	}
	w.Header().Set("Cache-Control", "max-age=600")
	if err := json.NewEncoder(w).Encode(resp); err != nil {
		k.t.Fatalf("json.NewEncoder(%v).Encode(%#v) = %q, want no error", w, keySet, err)
	}
}

var signKey *rsa.PrivateKey
var keySet jose.JSONWebKeySet

func NewToken(c *oauth2.Config, userInfo oidc.UserInfo) (string, error) {
	var err error
	priv := jose.JSONWebKey{Algorithm: "RS256", Key: getSignKey(), KeyID: "123", Use: "sig"}
	signer, err := jose.NewSigner(jose.SigningKey{Algorithm: jose.RS256, Key: priv}, nil)
	if err != nil {
		return "", err
	}

	id := &IdToken{
		Issuer:   c.Endpoint.AuthURL,
		Audience: []string{c.ClientID},
		Expiry:   time.Now().Add(5 * time.Minute).UTC().Unix(),
		IssuedAt: time.Now().UTC().Unix(),
		UserInfo: userInfo,
	}
	payload, err := json.Marshal(&id)
	if err != nil {
		return "", err
	}
	signed, err := signer.Sign(payload)
	if err != nil {
		return "", err
	}
	return signed.CompactSerialize()
}

func NewKeyServer(t *testing.T) (server *httptest.Server, cleanup func()) {
	// Arrange: Set up JWK keyserver.
	ks := &KeyServer{t: t}
	s := httptest.NewServer(ks)
	ks.provider.AuthURL = s.URL
	ks.provider.Issuer = s.URL
	ks.provider.JWKSURL = s.URL
	ks.URL = s.URL
	return s, s.Close
}

func getSignKey() *rsa.PrivateKey {
	var err error
	if signKey == nil {
		if signKey, err = rsa.GenerateKey(rand.Reader, 2048); err != nil {
			return nil
		}
	}
	return signKey
}

func getKeySet() jose.JSONWebKeySet {
	if len(keySet.Keys) == 0 {
		keySet = jose.JSONWebKeySet{
			Keys: []jose.JSONWebKey{
				{Algorithm: "RS256", Key: getSignKey().Public(), KeyID: "123", Use: "sig"},
			},
		}
	}
	return keySet
}

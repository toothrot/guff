package auth

import (
	"context"
	"crypto/rand"
	"crypto/rsa"
	"encoding/json"
	"fmt"
	"net"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
	"time"

	"github.com/coreos/go-oidc"
	"github.com/google/go-cmp/cmp"
	"golang.org/x/oauth2"
	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
	"gopkg.in/square/go-jose.v2"

	"github.com/toothrot/guff/backend/generated"
)

type providerJSON struct {
	Issuer      string `json:"issuer"`
	AuthURL     string `json:"authorization_endpoint"`
	TokenURL    string `json:"token_endpoint"`
	JWKSURL     string `json:"jwks_uri"`
	UserInfoURL string `json:"userinfo_endpoint"`
}

type idToken struct {
	Issuer   string   `json:"iss"`
	Subject  string   `json:"sub"`
	Audience []string `json:"aud"`
	Expiry   int64    `json:"exp"`
	IssuedAt int64    `json:"iat"`
	Nonce    string   `json:"nonce"`
	AtHash   string   `json:"at_hash"`
	oidc.UserInfo
}

type keyServer struct {
	t        *testing.T
	url      string
	provider providerJSON
}

func (k *keyServer) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	k.t.Helper()
	var resp interface{}
	switch r.URL.Path {
	case "/.well-known/openid-configuration":
		resp = k.provider
	default:
		resp = keySet
	}
	w.Header().Set("Cache-Control", "max-age=600")
	if err := json.NewEncoder(w).Encode(resp); err != nil {
		k.t.Fatalf("json.NewEncoder(%v).Encode(%#v) = %q, want no error", w, keySet, err)
	}
}

var signKey *rsa.PrivateKey
var keySet jose.JSONWebKeySet

func newToken(c *oauth2.Config, userInfo oidc.UserInfo) (string, error) {
	var err error
	if signKey == nil {
		if signKey, err = rsa.GenerateKey(rand.Reader, 2048); err != nil {
			return "", err
		}
	}
	priv := jose.JSONWebKey{Algorithm: "RS256", Key: signKey, KeyID: "123", Use: "sig"}
	pub := jose.JSONWebKey{Algorithm: "RS256", Key: signKey.Public(), KeyID: "123", Use: "sig"}
	keySet = jose.JSONWebKeySet{Keys: []jose.JSONWebKey{pub}}
	signer, err := jose.NewSigner(jose.SigningKey{Algorithm: jose.RS256, Key: priv}, nil)
	if err != nil {
		return "", err
	}

	id := &idToken{
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

type fakeService struct {
	guff_proto.UnimplementedTestServiceServer
}

func (f *fakeService) TestEcho(ctx context.Context, req *guff_proto.TestEchoRequest) (*guff_proto.TestEchoResponse, error) {
	return &guff_proto.TestEchoResponse{
		Message: strings.TrimSpace(fmt.Sprintf("hiya %s", EmailFromContext(ctx))),
	}, nil
}

func newTestServer(t *testing.T, opt ...grpc.ServerOption) (addr string, cleanup func()) {
	// Arrange: Setup fake GRPC handler.
	server := grpc.NewServer(opt...)
	guff_proto.RegisterTestServiceServer(server, &fakeService{})

	// Arrange: Start GRPC server.
	lis, err := net.Listen("tcp", ":0")
	if err != nil {
		t.Fatalf("net.Listen(%q, %q) = _, %q, want no error", "tcp", ":0", err)
	}
	go server.Serve(lis)
	return lis.Addr().String(), server.Stop
}

func newTestClient(ctx context.Context, t *testing.T, addr string) guff_proto.TestServiceClient {
	// Arrange: Setup client
	conn, err := grpc.DialContext(ctx, addr, grpc.WithInsecure())
	if err != nil {
		t.Fatalf("grpc.Dial(%q) = _, %q, want no error", addr, err)
	}
	return guff_proto.NewTestServiceClient(conn)
}

func newKeyServer(t *testing.T) (server *httptest.Server, cleanup func()) {
	// Arrange: Set up JWK keyserver.
	ks := &keyServer{t: t}
	s := httptest.NewServer(ks)
	ks.provider.AuthURL = s.URL
	ks.provider.Issuer = s.URL
	ks.provider.JWKSURL = s.URL
	return s, s.Close
}

func TestAuthMiddleware(t *testing.T) {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	ks, ksCleanup := newKeyServer(t)
	defer ksCleanup()

	// Arrange: Configure middleware
	config := &oauth2.Config{
		ClientID: "abc123",
		Endpoint: oauth2.Endpoint{AuthURL: ks.URL},
	}
	m, err := NewAuthMiddleware(ctx, config)
	if err != nil {
		t.Fatalf("NewAuthMiddleware(%v, %#v) = _, %q, want no error", ctx, config, err)
	}
	addr, stop := newTestServer(t, grpc.UnaryInterceptor(m.ServerInterceptor))
	defer stop()

	tok, _ := newToken(config, oidc.UserInfo{Email: "mario@example.com"})

	cases := []struct{
		desc string
		md metadata.MD
		want *guff_proto.TestEchoResponse
	}{
		{
			desc: "empty metadata",
			md: metadata.MD{},
			want: &guff_proto.TestEchoResponse{Message: "hiya"},
		},
		{
			desc: "empty authorization value",
			md: metadata.Pairs("authorization", ""),
			want: &guff_proto.TestEchoResponse{Message: "hiya"},
		},
		{
			desc: "malformated authorization value",
			md: metadata.Pairs("authorization", "barf"),
			want: &guff_proto.TestEchoResponse{Message: "hiya"},
		},
		{
			desc: "invalid token",
			md: metadata.Pairs("authorization", "Bearer garbage"),
			want: &guff_proto.TestEchoResponse{Message: "hiya"},
		},
		{
			desc: "valid token",
			md: metadata.Pairs("authorization", fmt.Sprintf("Bearer %s", tok)),
			want: &guff_proto.TestEchoResponse{Message: "hiya mario@example.com"},
		},
	}
	for _, c := range cases {
		// Act: Do request
		client := newTestClient(ctx, t, addr)
		ctx = metadata.NewOutgoingContext(ctx, c.md)
		req := &guff_proto.TestEchoRequest{Message: "hiya"}
		resp, err := client.TestEcho(ctx, req)
		if err != nil {
			t.Errorf("%q: client.TestEcho(%v, %#v) = _, %q, want no error", c.desc, ctx, req, err)
		}

		// Assert: response
		if diff := cmp.Diff(c.want, resp); diff != "" {
			t.Errorf("client.TestEcho(%v, %#v) mismatch (-want +got):\n%s", ctx, req, diff)
		}
	}

}

package auth

import (
	"context"
	"net/url"
	"strings"

	"github.com/coreos/go-oidc"
	"github.com/golang/glog"
	"golang.org/x/oauth2"
	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
)

type authContextKey int

const emailKey authContextKey = iota

func NewAuthMiddleware(ctx context.Context, c *oauth2.Config) (*AuthMiddleware, error) {
	issuer, err := url.Parse(c.Endpoint.AuthURL)
	if err != nil {
		return nil, err
	}
	issuer.Path = ""
	p, err := oidc.NewProvider(ctx, issuer.String())
	if err != nil {
		return nil, err
	}
	v := p.Verifier(&oidc.Config{ClientID: c.ClientID})
	return &AuthMiddleware{
		verifier: v,
	}, nil
}

type Provider interface {
	Verifier(config *oidc.Config) *oidc.IDTokenVerifier
}

type IDTokenVerifier interface {
	Verify(ctx context.Context, rawIDToken string) (*oidc.IDToken, error)
}

type AuthMiddleware struct {
	verifier IDTokenVerifier
}

func (a *AuthMiddleware) ServerInterceptor(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (resp interface{}, err error) {
	md, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		return handler(ctx, req)
	}

	auth := md.Get("authorization")
	if len(auth) == 0 {
		return handler(ctx, req)
	}

	tok := strings.TrimPrefix(auth[0], "Bearer ")
	token, err := a.verifier.Verify(ctx, tok)
	if err != nil {
		glog.Errorf("failed to verify: %q", err)
		return handler(ctx, req)
	}

	var userInfo oidc.UserInfo
	if err := token.Claims(&userInfo); err != nil {
		return handler(ctx, req)
	}
	ctx = context.WithValue(ctx, emailKey, userInfo.Email)

	return handler(ctx, req)
}

func EmailFromContext(ctx context.Context) string {
	email, ok := ctx.Value(emailKey).(string)
	if !ok {
		return ""
	}
	return email
}

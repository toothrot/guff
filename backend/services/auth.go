package services

import (
	"context"
	"net/url"
	"strings"

	"github.com/coreos/go-oidc"
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
		provider: p,
		verifier: v,
	}, nil
}

type AuthMiddleware struct {
	provider *oidc.Provider
	verifier *oidc.IDTokenVerifier
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
		return handler(ctx, req)
	}

	var claims struct {
		Email string `json:"email"`
	}
	if err := token.Claims(&claims); err != nil {
		return handler(ctx, req)
	}
	ctx = context.WithValue(ctx, emailKey, claims.Email)

	return handler(ctx, req)
}

func EmailFromContext(ctx context.Context) string {
	email, ok := ctx.Value(emailKey).(string)
	if !ok {
		return ""
	}
	return email
}

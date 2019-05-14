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

	"github.com/toothrot/guff/backend/models"
)

type authContextKey int

const (
	emailKey authContextKey = iota
	userKey  authContextKey = iota
)

func NewMiddleware(ctx context.Context, c *oauth2.Config, p models.Persist) (*Middleware, error) {
	issuer, err := url.Parse(c.Endpoint.AuthURL)
	if err != nil {
		return nil, err
	}
	issuer.Path = ""
	provider, err := oidc.NewProvider(ctx, issuer.String())
	if err != nil {
		return nil, err
	}
	v := provider.Verifier(&oidc.Config{ClientID: c.ClientID})
	return &Middleware{
		verifier: v,
		Persist:  p,
	}, nil
}

type IDTokenVerifier interface {
	Verify(ctx context.Context, rawIDToken string) (*oidc.IDToken, error)
}

// Middleware implements middleware for validating tokens.
type Middleware struct {
	verifier IDTokenVerifier
	Persist  models.Persist
}

// ServerInterceptor is a GRPC Unary interceptor for validating Authorization headers.
//
// It relies on the go-oidc library for validation. The underlying library ensures that the Client ID matches the audience
// of the token, and verifies that the signature is correct for the provider. It also handles caching of keys.
func (a *Middleware) ServerInterceptor(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (resp interface{}, err error) {
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
	u, err := a.Persist.FindOrCreateUser(ctx, strings.ToLower(userInfo.Email))
	if err != nil {
		glog.Errorf("FindOrCreateUser() = %v", err)
		return handler(ctx, req)
	}
	ctx = context.WithValue(ctx, emailKey, u.Email)
	ctx = context.WithValue(ctx, userKey, &u)

	return handler(ctx, req)
}

// FakeMiddleware is a stub middleware to inject in tests to set auth values on the context.
type FakeMiddleware struct {
	Email string
}

// ServerInterceptor is a GRPC Unary interceptor for faking out auth contexts.
func (f *FakeMiddleware) ServerInterceptor(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (resp interface{}, err error) {
	return handler(context.WithValue(ctx, emailKey, f.Email), req)
}

func AuthContext(ctx context.Context, email string, user models.User) context.Context {
	ctx = context.WithValue(ctx, emailKey, email)
	ctx = context.WithValue(ctx, userKey, &user)
	return ctx
}

func EmailFromContext(ctx context.Context) string {
	email, ok := ctx.Value(emailKey).(string)
	if !ok {
		return ""
	}
	return email
}

func UserFromContext(ctx context.Context) *models.User {
	u, ok := ctx.Value(userKey).(*models.User)
	if !ok {
		return nil
	}
	return u
}

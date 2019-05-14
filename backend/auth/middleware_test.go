package auth

import (
	"context"
	"fmt"
	"net"
	"strings"
	"testing"

	"github.com/coreos/go-oidc"
	"github.com/google/go-cmp/cmp"
	"golang.org/x/oauth2"
	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"

	"github.com/toothrot/guff/backend/auth/test"
	"github.com/toothrot/guff/backend/generated"
	"github.com/toothrot/guff/backend/models"
)

type fakeService struct {
	guff_proto.UnimplementedTestServiceServer
}

func (f *fakeService) TestEcho(ctx context.Context, req *guff_proto.TestEchoRequest) (*guff_proto.TestEchoResponse, error) {
	return &guff_proto.TestEchoResponse{
		Message: strings.TrimSpace(fmt.Sprintf("hiya %s %t", EmailFromContext(ctx), UserFromContext(ctx).GetIsAdmin())),
	}, nil
}

func NewTestServer(t *testing.T, opt ...grpc.ServerOption) (addr string, cleanup func()) {
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

func NewTestClient(ctx context.Context, t *testing.T, addr string) guff_proto.TestServiceClient {
	// Arrange: Setup client
	conn, err := grpc.DialContext(ctx, addr, grpc.WithInsecure())
	if err != nil {
		t.Fatalf("grpc.Dial(%q) = _, %q, want no error", addr, err)
	}
	return guff_proto.NewTestServiceClient(conn)
}

func TestAuthMiddleware(t *testing.T) {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	ks, ksCleanup := test.NewKeyServer(t)
	defer ksCleanup()
	if err := models.DefaultMemoryPersist.TruncateUsers(ctx); err != nil {
		t.Fatalf("models.DefaultMemoryPersist.TruncateUsers(%v) = %v, wanted no error", ctx, err)
	}

	// Arrange: Configure middleware
	config := &oauth2.Config{
		ClientID: "abc123",
		Endpoint: oauth2.Endpoint{AuthURL: ks.URL},
	}
	m, err := NewMiddleware(ctx, config, models.DefaultMemoryPersist)
	if err != nil {
		t.Fatalf("NewMiddleware(%v, %#v) = _, %q, want no error", ctx, config, err)
	}
	addr, stop := NewTestServer(t, grpc.UnaryInterceptor(m.ServerInterceptor))
	defer stop()

	tok, _ := test.NewToken(config, oidc.UserInfo{Email: "mario@example.com"})
	adminTok, _ := test.NewToken(config, oidc.UserInfo{Email: "admin@example.com"})
	if _, err := models.DefaultMemoryPersist.FindOrCreateUser(ctx, "admin@example.com"); err != nil {
		t.Fatalf("models.DefaultMemoryPersist.FindOrCreateUser(%v, %v) = _, %v, wanted no error", ctx, "admin@example.com", err)
	}

	cases := []struct {
		desc string
		md   metadata.MD
		want *guff_proto.TestEchoResponse
	}{
		{
			desc: "empty metadata",
			md:   metadata.MD{},
			want: &guff_proto.TestEchoResponse{Message: "hiya  false"},
		},
		{
			desc: "empty authorization value",
			md:   metadata.Pairs("authorization", ""),
			want: &guff_proto.TestEchoResponse{Message: "hiya  false"},
		},
		{
			desc: "malformated authorization value",
			md:   metadata.Pairs("authorization", "barf"),
			want: &guff_proto.TestEchoResponse{Message: "hiya  false"},
		},
		{
			desc: "invalid token",
			md:   metadata.Pairs("authorization", "Bearer garbage"),
			want: &guff_proto.TestEchoResponse{Message: "hiya  false"},
		},
		{
			desc: "valid token",
			md:   metadata.Pairs("authorization", fmt.Sprintf("Bearer %s", tok)),
			want: &guff_proto.TestEchoResponse{Message: "hiya mario@example.com false"},
		},
		{
			desc: "valid admin token",
			md:   metadata.Pairs("authorization", fmt.Sprintf("Bearer %s", adminTok)),
			want: &guff_proto.TestEchoResponse{Message: "hiya admin@example.com true"},
		},
	}
	for _, c := range cases {
		t.Run(c.desc, func(t *testing.T) {
			// Act: Do request
			client := NewTestClient(ctx, t, addr)
			ctx = metadata.NewOutgoingContext(ctx, c.md)
			req := &guff_proto.TestEchoRequest{Message: "hiya"}
			resp, err := client.TestEcho(ctx, req)
			if err != nil {
				t.Errorf("client.TestEcho(%v, %#v) = _, %q, want no error", ctx, req, err)
			}

			// Assert: response
			if diff := cmp.Diff(c.want, resp); diff != "" {
				t.Errorf("client.TestEcho(%v, %#v) mismatch (-want +got):\n%s", ctx, req, diff)
			}
		})
	}

}

package services

import (
	"context"
	"net/http"

	"github.com/golang/glog"
	"google.golang.org/grpc/metadata"

	"github.com/toothrot/guff/go/core"

	guff_proto "github.com/toothrot/guff/go/generated"
)

type Users struct {
	guff_proto.UnimplementedUsersServiceServer
	Config *core.Config
}

func (u *Users) GetCurrentUser(ctx context.Context, req *guff_proto.GetCurrentUserRequest) (*guff_proto.GetCurrentUserResponse, error) {
	md, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		glog.Errorf("fromincomingcontext ok: %v", ok)
	}
	glog.Infof("metadata: %#v", md)
	r := &http.Request{Header: http.Header{}}
	for k, v := range md {
		for _, val := range v {
			r.Header.Add(k, val)
		}
	}
	session, err := u.Config.CookieStore.Get(r, "guff")
	if err != nil {
		glog.Errorf("Cookiestore get failure: %q", err)
		return nil, nil
	}
	resp := &guff_proto.GetCurrentUserResponse{
		Email: session.Values["email"].(string),
	}
	return resp, nil
}

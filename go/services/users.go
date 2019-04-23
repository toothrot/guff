package services

import (
	"context"

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

	resp := &guff_proto.GetCurrentUserResponse{
		GoogleOauthConfig: &guff_proto.GoogleOAuthConfig{
			ClientId: u.Config.OAuthConfig.ClientID,
			LoginURL: u.Config.OAuthConfig.AuthCodeURL(""),
		},
	}
	return resp, nil
}

package services

import (
	"context"

	"github.com/toothrot/guff/backend/core"

	guff_proto "github.com/toothrot/guff/backend/generated"
)

type Users struct {
	guff_proto.UnimplementedUsersServiceServer
	Config *core.Config
}

func (u *Users) GetCurrentUser(ctx context.Context, req *guff_proto.GetCurrentUserRequest) (*guff_proto.GetCurrentUserResponse, error) {
	resp := &guff_proto.GetCurrentUserResponse{
		Email: EmailFromContext(ctx),
		GoogleOauthConfig: &guff_proto.GoogleOAuthConfig{
			ClientId: u.Config.OAuthConfig.ClientID,
			LoginURL: u.Config.OAuthConfig.AuthCodeURL(""),
		},
	}
	return resp, nil
}

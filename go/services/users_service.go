package services

import (
	"context"

	guff_proto "github.com/toothrot/guff/go/generated"
)

type UsersService struct {
}

func (u *UsersService) GetCurrentUser(ctx context.Context, req *guff_proto.GetCurrentUserRequest) (*guff_proto.GetCurrentUserResponse, error) {
	resp := &guff_proto.GetCurrentUserResponse{
		AppStatus: &guff_proto.AppStatus{Code: guff_proto.AppStatusCode_APP_STATUS_OK},
	}
	return resp, nil
}

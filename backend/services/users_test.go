package services

import (
	"context"
	"testing"

	"github.com/google/go-cmp/cmp"
	"golang.org/x/oauth2"

	"github.com/toothrot/guff/backend/auth"
	"github.com/toothrot/guff/backend/core"
	guff_proto "github.com/toothrot/guff/backend/generated"
	"github.com/toothrot/guff/backend/models"
)

func TestUsers_GetCurrentUser(t *testing.T) {
	u := &Users{
		Config: &core.Config{
			OAuthConfig: &oauth2.Config{ClientID: "123"},
		},
	}
	tests := []struct {
		name  string
		email string
		user  models.User
		want  *guff_proto.GetCurrentUserResponse
	}{
		{
			name: "unauthenticated",
			want: &guff_proto.GetCurrentUserResponse{
				GoogleOauthConfig: &guff_proto.GoogleOAuthConfig{ClientId: "123", LoginURL: "?client_id=123&response_type=code"},
			},
		},
		{
			name:  "authenticated",
			email: "testuser@example.com",
			user:  models.User{Email: "testuser@example.com", IsAdmin: true},
			want: &guff_proto.GetCurrentUserResponse{
				Email:             "testuser@example.com",
				IsAdmin:           true,
				GoogleOauthConfig: &guff_proto.GoogleOAuthConfig{ClientId: "123", LoginURL: "?client_id=123&response_type=code"},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctx, cancel := context.WithCancel(auth.AuthContext(context.Background(), tt.email, tt.user))
			defer cancel()
			req := &guff_proto.GetCurrentUserRequest{}

			got, err := u.GetCurrentUser(ctx, req)
			if err != nil {
				t.Errorf("Users.GetCurrentUser(%v) = _, %v, wanted no error", req, err)
				return
			}
			if diff := cmp.Diff(tt.want, got); diff != "" {
				t.Errorf("a.Scrape(%v, %v) mismatch (-want +got):\n%s", ctx, req, diff)
			}
		})
	}
}

package core

import (
	"github.com/gorilla/sessions"
	"golang.org/x/oauth2"
)

type Config struct {
	OAuthConfig *oauth2.Config
	CookieStore *sessions.CookieStore
}

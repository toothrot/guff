package core

import (
	"github.com/gorilla/sessions"
	"golang.org/x/oauth2"
)

type Config struct {
	OAuthConfig *oauth2.Config
	CookieStore *sessions.CookieStore

	// URL to scrape for list of Programs (divisions in a league)
	ProgramsURL string
}

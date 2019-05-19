package core

import (
	"golang.org/x/oauth2"
)

type Config struct {
	OAuthConfig *oauth2.Config

	// URL to scrape for list of Programs (divisions in a league)
	ProgramsURL  string
	DBName       string
	DBPassword   string
	DBURL        string
	RequireHTTPS string
}

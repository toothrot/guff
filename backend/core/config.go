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
	ScheduleURL  string
}

const scheduleURL = "https://royalpalmsshuffle.leagueapps.com/ajax/loadSchedule"

func (c *Config) GetScheduleURL() string {
	if c.ScheduleURL == "" {
		return scheduleURL
	}
	return c.ScheduleURL
}

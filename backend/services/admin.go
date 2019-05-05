package services

import (
	"context"
	"io/ioutil"
	"net/http"
	"regexp"

	"github.com/toothrot/guff/backend/core"
	"github.com/toothrot/guff/backend/generated"
)

var (
	divisions   []Division
	divisionRex = regexp.MustCompile("leagues/([0-9]+)/schedule")
)

type Admin struct {
	Config *core.Config

	guff_proto.UnimplementedAdminServiceServer
}

func (a *Admin) Scrape(ctx context.Context, req *guff_proto.ScrapeRequest) (*guff_proto.ScrapeResponse, error) {
	resp, err := http.Get(a.Config.ProgramsURL)
	if err != nil {
		return nil, err
	}
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	matches := divisionRex.FindAllSubmatch(body, -1)

	for _, match := range matches {
		divisions = append(divisions, Division{ID: string(match[1])})
	}
	return &guff_proto.ScrapeResponse{}, nil
}

type Division struct {
	ID   string
	Name string
}

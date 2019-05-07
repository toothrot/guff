package services

import (
	"context"
	"io/ioutil"
	"net/http"

	"github.com/toothrot/guff/backend/core"
	"github.com/toothrot/guff/backend/generated"
	"github.com/toothrot/guff/backend/models"
)

var divisions []models.Division

type Admin struct {
	Config *core.Config

	guff_proto.UnimplementedAdminServiceServer
}

func (a *Admin) Scrape(ctx context.Context, req *guff_proto.ScrapeRequest) (*guff_proto.ScrapeResponse, error) {
	resp, err := http.Get(a.Config.ProgramsURL)
	if err != nil {
		return nil, err
	}
	b, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	divisions = models.ParseDivisions(b)
	return &guff_proto.ScrapeResponse{}, nil
}

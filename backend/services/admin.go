package services

import (
	"context"
	"io/ioutil"
	"net/http"
	"net/url"

	"github.com/golang/glog"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"

	"github.com/toothrot/guff/backend/auth"
	"github.com/toothrot/guff/backend/core"
	"github.com/toothrot/guff/backend/generated"
	"github.com/toothrot/guff/backend/models"
)

const (
	scheduleQuery = "origin=site&scope=program&publishedOnly=0&itemType=games_events&programId=1039207"
)

type DivisionParserFunc func(b []byte) []models.Division
type TeamParserFunc func(b []byte) []models.Team

type Admin struct {
	Config         *core.Config
	Persist        models.Persist
	DivisionParser DivisionParserFunc
	TeamParser     TeamParserFunc

	guff_proto.UnimplementedAdminServiceServer
}

func (a *Admin) Scrape(ctx context.Context, req *guff_proto.ScrapeRequest) (*guff_proto.ScrapeResponse, error) {
	user := auth.UserFromContext(ctx)
	if !user.GetIsAdmin() {
		return nil, grpc.Errorf(codes.PermissionDenied, codes.PermissionDenied.String())
	}
	if err := a.scrapeDivisions(ctx); err != nil {
		return nil, err
	}
	divisions, err := a.Persist.GetDivisions(ctx)
	if err != nil {
		return nil, err
	}
	for _, d := range divisions {
		if err := a.scrapeTeams(ctx, d.ID); err != nil {
			return nil, err
		}
	}
	return &guff_proto.ScrapeResponse{}, nil
}

func (a *Admin) scrapeDivisions(ctx context.Context) error {
	resp, err := http.Get(a.Config.ProgramsURL)
	if err != nil {
		return err
	}
	b, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}
	ds := a.DivisionParser(b)
	if err := a.Persist.UpsertDivisions(ctx, ds); err != nil {
		glog.Errorf("a.Persist.UpsertDivisions(%v, %v) = %q", ctx, ds, err)
		return grpc.Errorf(codes.Internal, codes.Internal.String())
	}
	return nil
}

func (a *Admin) scrapeTeams(ctx context.Context, divisionID string) error {
	u, err := url.Parse(a.Config.GetScheduleURL())
	if err != nil {
		return err
	}
	q, err := url.ParseQuery(scheduleQuery)
	if err != nil {
		return err
	}
	q.Set("programId", divisionID)
	u.RawQuery = q.Encode()
	resp, err := http.Get(u.String())
	if err != nil {
		return err
	}
	b, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}
	ts := a.TeamParser(b)
	if err := a.Persist.UpsertTeams(ctx, ts); err != nil {
		return err
	}
	return nil
}

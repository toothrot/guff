package services

import (
	"context"
	"io/ioutil"
	"net/http"

	"github.com/golang/glog"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"

	"github.com/toothrot/guff/backend/auth"
	"github.com/toothrot/guff/backend/core"
	"github.com/toothrot/guff/backend/generated"
	"github.com/toothrot/guff/backend/models"
)

type Admin struct {
	Config  *core.Config
	Persist models.Persist

	guff_proto.UnimplementedAdminServiceServer
}

func (a *Admin) Scrape(ctx context.Context, req *guff_proto.ScrapeRequest) (*guff_proto.ScrapeResponse, error) {
	user := auth.UserFromContext(ctx)
	if !user.GetIsAdmin() {
		return nil, grpc.Errorf(codes.PermissionDenied, codes.PermissionDenied.String())
	}
	resp, err := http.Get(a.Config.ProgramsURL)
	if err != nil {
		return nil, err
	}
	b, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	ds := models.ParseDivisions(b)
	if err := a.Persist.UpsertDivisions(ctx, ds); err != nil {
		glog.Errorf("a.Persist.UpsertDivisions(%v, %v) = %q", ctx, ds, err)
		return nil, grpc.Errorf(codes.Internal, codes.Internal.String())
	}
	return &guff_proto.ScrapeResponse{}, nil
}

package services

import (
	"context"

	"github.com/golang/glog"

	"github.com/toothrot/guff/backend/generated"
	"github.com/toothrot/guff/backend/models"
)

type Teams struct {
	Persist models.Persist

	guff_proto.UnimplementedTeamsServiceServer
}

func (t *Teams) GetTeams(ctx context.Context, req *guff_proto.GetTeamsRequest) (*guff_proto.GetTeamsResponse, error) {
	resp := &guff_proto.GetTeamsResponse{}
	ts, err := t.Persist.GetTeams(ctx, req.GetDivisionId())
	if err != nil {
		glog.Errorf("t.Persist.GetTeams(_, %q) = _, %v, wanted no error", req.GetDivisionId(), err)
		return resp, err
	}
	for _, team := range ts {
		resp.Teams = append(resp.Teams, team.ToProto())
	}
	return resp, nil
}

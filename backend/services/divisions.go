package services

import (
	"context"

	"github.com/golang/glog"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"

	"github.com/toothrot/guff/backend/generated"
	"github.com/toothrot/guff/backend/models"
)

type Divisions struct {
	Persist models.Persist

	guff_proto.UnimplementedDivisionsServiceServer
}

func (d *Divisions) GetDivisions(ctx context.Context, req *guff_proto.GetDivisionsRequest) (*guff_proto.GetDivisionsResponse, error) {
	resp := &guff_proto.GetDivisionsResponse{}
	divisions, err := d.Persist.GetDivisions(ctx)
	if err != nil {
		glog.Errorf("d.Persist.GetDivisions(%v) = %v", ctx, err)
		return nil, grpc.Errorf(codes.Internal, codes.Internal.String())
	}
	for _, d := range divisions {
		resp.Divisions = append(resp.Divisions, d.ToProto())
	}
	return resp, nil
}

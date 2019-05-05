package services

import (
	"context"

	"github.com/toothrot/guff/backend/generated"
)

type DivisionService struct {
	guff_proto.UnimplementedDivisionsServiceServer
}

func (d *DivisionService) GetDivisions(ctx context.Context, req *guff_proto.GetDivisionsRequest) (*guff_proto.GetDivisionsResponse, error) {
	resp := &guff_proto.GetDivisionsResponse{}
	for _, d := range divisions {
		resp.Divisions = append(resp.Divisions, &guff_proto.Division{Id: d.ID})
	}
	return resp, nil
}

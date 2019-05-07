package services

import (
	"context"

	"github.com/toothrot/guff/backend/generated"
)

type Divisions struct {
	guff_proto.UnimplementedDivisionsServiceServer
}

func (d *Divisions) GetDivisions(ctx context.Context, req *guff_proto.GetDivisionsRequest) (*guff_proto.GetDivisionsResponse, error) {
	resp := &guff_proto.GetDivisionsResponse{}
	for _, d := range divisions {
		resp.Divisions = append(resp.Divisions, &guff_proto.Division{Id: d.ID})
	}
	return resp, nil
}

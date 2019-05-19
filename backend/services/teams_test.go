package services

import (
	"context"
	"testing"

	"github.com/google/go-cmp/cmp"

	guff_proto "github.com/toothrot/guff/backend/generated"
	"github.com/toothrot/guff/backend/models"
)

func TestTeams_GetTeams(t *testing.T) {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	defer models.DefaultMemoryPersist.TruncateTeams()
	models.DefaultMemoryPersist.TruncateTeams()

	want := []*guff_proto.Team{{Id: "123"}}
	ts := []models.Team{{ID: "123"}}
	if err := models.DefaultMemoryPersist.UpsertTeams(ctx, ts); err != nil {
		t.Fatalf("models.DefaultMemoryPersist.UpsertTeams(%v, %v) = %v, wanted no error", ctx, ts, err)
	}
	d := &Teams{Persist: models.DefaultMemoryPersist}

	req := &guff_proto.GetTeamsRequest{}
	got, err := d.GetTeams(context.Background(), req)
	if err != nil {
		t.Fatalf("a.GetTeams(%v, %v) = _, %q", context.Background(), req, err)
	}
	if diff := cmp.Diff(want, got.GetTeams()); diff != "" {
		t.Errorf("d.GetTeams(%v, %v) mismatch (-want +got):\n%s", ctx, req, diff)
	}
}

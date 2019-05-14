package services

import (
	"context"
	"testing"

	"github.com/google/go-cmp/cmp"

	guff_proto "github.com/toothrot/guff/backend/generated"
	"github.com/toothrot/guff/backend/models"
)

func TestDivisions_GetDivisions(t *testing.T) {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	defer models.DefaultMemoryPersist.TruncateDivisions(ctx)
	if err := models.DefaultMemoryPersist.TruncateDivisions(ctx); err != nil {
		t.Fatalf("a.Persist.TruncateDivisions(%v) = %v, wanted no error", ctx, err)
	}
	want := []*guff_proto.Division{{Id: "1039207"}, {Id: "1039210"}}
	ds := []models.Division{{ID: "1039207"}, {ID: "1039210"}}
	if err := models.DefaultMemoryPersist.UpsertDivisions(ctx, ds); err != nil {
		t.Fatalf("models.DefaultMemoryPersist.UpsertDivisions(%v, %v) = %v, wanted no error", ctx, want, err)
	}
	d := &Divisions{Persist: models.DefaultMemoryPersist}

	req := &guff_proto.GetDivisionsRequest{}
	got, err := d.GetDivisions(context.Background(), req)
	if err != nil {
		t.Fatalf("a.GetDivisions(%v, %v) = _, %q", context.Background(), req, err)
	}
	if diff := cmp.Diff(want, got.GetDivisions()); diff != "" {
		t.Errorf("d.GetDivisions(%v, %v) mismatch (-want +got):\n%s", ctx, req, diff)
	}
}

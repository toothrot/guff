package services

import (
	"context"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/google/go-cmp/cmp"

	"github.com/toothrot/guff/backend/core"
	"github.com/toothrot/guff/backend/generated"
	"github.com/toothrot/guff/backend/models"
)

const programsList = `
sonnets
leagues/1039207/schedule
leagues/1039210/schedule
html
leagues/1039208/bar/schedule/
prose
leagues/1039209/standings/
`

func TestAdmin_Scrape(t *testing.T) {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	defer models.DefaultMemoryPersist.TruncateDivisions(ctx)
	s := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if _, err := w.Write([]byte(programsList)); err != nil {
			t.Fatalf("w.Write() = %q, wanted no error", err)
		}
	}))
	if err := models.DefaultMemoryPersist.TruncateDivisions(ctx); err != nil {
		t.Fatalf("a.Persist.TruncateDivisions(%v) = %v, wanted no error", ctx, err)
	}
	a := &Admin{
		Config:  &core.Config{ProgramsURL: s.URL},
		Persist: models.DefaultMemoryPersist,
	}

	req := &guff_proto.ScrapeRequest{}
	if _, err := a.Scrape(context.Background(), req); err != nil {
		t.Fatalf("a.Scrape(%v, %v) = _, %q", context.Background(), req, err)
	}

	want := []models.Division{{ID: "1039207"}, {ID: "1039210"}}
	got, err := models.DefaultMemoryPersist.GetDivisions(ctx)
	if err != nil {
		t.Fatalf("models.DefaultMemoryPersist.GetDivisions(%v) = _, %v, wanted no error", ctx, err)
	}
	if diff := cmp.Diff(want, got); diff != "" {
		t.Errorf("a.Scrape(%v, %v) mismatch (-want +got):\n%s", ctx, req, diff)
	}
}

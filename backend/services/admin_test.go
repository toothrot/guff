package services

import (
	"context"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/toothrot/guff/backend/core"
	"github.com/toothrot/guff/backend/generated"
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
	s := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if _, err := w.Write([]byte(programsList)); err != nil {
			t.Fatalf("w.Write() = %q, wanted no error", err)
		}
	}))
	a := &Admin{
		Config: &core.Config{ProgramsURL: s.URL},
	}

	if _, err := a.Scrape(context.Background(), &guff_proto.ScrapeRequest{}); err != nil {
		t.Fatalf("a.Scrape(%v, %v) = _, %q", context.Background(), &guff_proto.ScrapeRequest{}, err)
	}

	if len(divisions) != 2 {
		t.Fatalf("len(%v) = %d, wanted %d", divisions, len(divisions), 2)
	}

	if divisions[0].ID != "1039207" {
		t.Errorf("divisions[%d].ID = %q, wanted %q", 0, divisions[0].ID, "1039207")
	}

	if divisions[1].ID != "1039210" {
		t.Errorf("divisions[%d].ID = %q, wanted %q", 1, divisions[1].ID, "1039210")
	}
}

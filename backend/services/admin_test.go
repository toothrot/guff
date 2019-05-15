package services

import (
	"context"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/google/go-cmp/cmp"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/toothrot/guff/backend/auth"
	"github.com/toothrot/guff/backend/core"
	"github.com/toothrot/guff/backend/generated"
	"github.com/toothrot/guff/backend/models"
)

type fakeParser struct {
	wantArgs  []byte
	divisions []models.Division
}

func (f *fakeParser) parse(b []byte) []models.Division {
	if string(b) != string(f.wantArgs) {
		return []models.Division{}
	}
	return f.divisions
}

func TestAdmin_Scrape(t *testing.T) {
	u := models.User{Email: "testuser@example.com", IsAdmin: true}
	ctx := auth.AuthContext(context.Background(), u.Email, u)
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()
	defer models.DefaultMemoryPersist.TruncateDivisions(ctx)
	if err := models.DefaultMemoryPersist.TruncateDivisions(ctx); err != nil {
		t.Fatalf("a.Persist.TruncateDivisions(%v) = %v, wanted no error", ctx, err)
	}

	s := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if _, err := w.Write([]byte("some divisions")); err != nil {
			t.Fatalf("w.Write() = %q, wanted no error", err)
		}
	}))
	want := []models.Division{{ID: "1039207"}, {ID: "1039210"}}
	f := fakeParser{divisions: want, wantArgs: []byte("some divisions")}
	a := &Admin{
		Config:         &core.Config{ProgramsURL: s.URL},
		Persist:        models.DefaultMemoryPersist,
		DivisionParser: f.parse,
	}

	req := &guff_proto.ScrapeRequest{}
	if _, err := a.Scrape(ctx, req); err != nil {
		t.Fatalf("a.Scrape(%v, %v) = _, %q", context.Background(), req, err)
	}

	got, err := models.DefaultMemoryPersist.GetDivisions(ctx)
	if err != nil {
		t.Fatalf("models.DefaultMemoryPersist.GetDivisions(%v) = _, %v, wanted no error", ctx, err)
	}
	if diff := cmp.Diff(want, got); diff != "" {
		t.Errorf("a.Scrape(%v, %v) mismatch (-want +got):\n%s", ctx, req, diff)
	}
}

func TestAdmin_ScrapeErrors(t *testing.T) {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	defer models.DefaultMemoryPersist.TruncateDivisions(ctx)
	if err := models.DefaultMemoryPersist.TruncateDivisions(ctx); err != nil {
		t.Fatalf("a.Persist.TruncateDivisions(%v) = %v, wanted no error", ctx, err)
	}

	tests := []struct {
		desc     string
		user     models.User
		wantCode codes.Code
	}{
		{
			desc:     "unauthenticated",
			wantCode: codes.PermissionDenied,
		},
		{
			desc:     "authenticated as non-admin",
			user:     models.User{Email: "rando@example.com", IsAdmin: false},
			wantCode: codes.PermissionDenied,
		},
	}
	for _, tt := range tests {
		t.Run(tt.desc, func(t *testing.T) {
			ctx = auth.AuthContext(ctx, tt.user.Email, tt.user)
			a := &Admin{}
			req := &guff_proto.ScrapeRequest{}
			_, err := a.Scrape(ctx, req)
			if err == nil {
				t.Fatalf("a.Scrape(%v, %v) = _, %q, wanted error", context.Background(), req, err)
			}
			if s, ok := status.FromError(err); !ok || s.Code() != tt.wantCode {
				t.Errorf("status.FromError(%v) = %v, %v, wanted %v, %v", err, s, ok, tt.wantCode.String(), true)
			}
		})

	}
}

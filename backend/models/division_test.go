package models

import (
	"io/ioutil"
	"os"
	"testing"

	"github.com/google/go-cmp/cmp"
)

const leaguesFile = "./testdata/leagues.html"

func TestParseDivisions(t *testing.T) {
	f, err := os.Open(leaguesFile)
	if err != nil {
		t.Fatalf("os.Open(%q) = _, %v, wanted no error", leaguesFile, err)
	}
	html, err := ioutil.ReadAll(f)
	if err != nil {
		t.Fatalf("ioutil.ReadAll(%v) = _, %v, wanted no error", f, err)
	}

	tests := []struct {
		name    string
		leagues []byte
		want    []Division
	}{
		{
			name:    "static leagues",
			leagues: html,
			want: []Division{
				{ID: "1039207", Name: "Spring 2019 Monday 7pm Pilot Division"},
				{ID: "1039210", Name: "Spring 2019 Monday 8pm Cherry Division"},
				{ID: "1039213", Name: "Spring 2019 Monday 9pm Hammer Division (CHALLENGE)"},
				{ID: "1039216", Name: "Spring 2019 Tuesday 7pm Pilot Division"},
				{ID: "1039219", Name: "Spring 2019 Tuesday 8pm Cherry Division"},
				{ID: "1039222", Name: "Spring 2019 Tuesday 9pm Hammer Division"},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := ParseDivisions(tt.leagues)
			if diff := cmp.Diff(tt.want, got); diff != "" {
				t.Errorf("a.ParseDivisions(tt.leagues) mismatch (-want +got):\n%s", diff)
			}
		})
	}
}

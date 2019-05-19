package models

import (
	"io/ioutil"
	"os"
	"sort"
	"testing"

	"github.com/google/go-cmp/cmp"
)

const scheduleFile = "./testdata/loadSchedule.html"

func TestParseTeams(t *testing.T) {
	f, err := os.Open(scheduleFile)
	if err != nil {
		t.Fatalf("os.Open(%q) = _, %v, wanted no error", leaguesFile, err)
	}
	html, err := ioutil.ReadAll(f)
	if err != nil {
		t.Fatalf("ioutil.ReadAll(%v) = _, %v, wanted no error", f, err)
	}

	tests := []struct {
		name  string
		input []byte
		want  []Team
	}{
		{
			name:  "schedule fixture",
			input: html,
			want: []Team{
				{ID: "2218972", DivisionID: "1039207", Name: "Palm Pilots"},
				{ID: "2225212", DivisionID: "1039207", Name: "Moral Qualms"},
				{ID: "2232661", DivisionID: "1039207", Name: "Original Tangsters"},
				{ID: "2244430", DivisionID: "1039207", Name: "Tangs of New York"},
				{ID: "2245132", DivisionID: "1039207", Name: "Truffle Shuffle"},
				{ID: "2246191", DivisionID: "1039207", Name: "Hammer Time"},
				{ID: "2246242", DivisionID: "1039207", Name: "Disky Business"},
				{ID: "2247763", DivisionID: "1039207", Name: "Tang U, Next"},
				{ID: "2248279", DivisionID: "1039207", Name: "Golden Tangs and the Kitchen of Doom"},
				{ID: "2254816", DivisionID: "1039207", Name: "VIPPS"},
				{ID: "2260024", DivisionID: "1039207", Name: "Shuffle Puck Time"},
				{ID: "2267368", DivisionID: "1039207", Name: "Park Slope Parent Teacher Association"},
				{ID: "2268253", DivisionID: "1039207", Name: "Welcome To Flavortown"},
				{ID: "2273749", DivisionID: "1039207", Name: "The Bloodhound Tang (formerly the Jager Palmers)"},
				{ID: "2274175", DivisionID: "1039207", Name: "Biscuit Babes"},
				{ID: "2275039", DivisionID: "1039207", Name: "Old Guys Rule"},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := ParseTeams(tt.input)

			trans := cmp.Transformer("Sort", func(in []Team) []Team {
				out := append([]Team(nil), in...)
				sort.Sort(TeamByID(out))
				return out
			})
			if diff := cmp.Diff(tt.want, got, trans); diff != "" {
				t.Errorf("ParseTeams(tt.input) mismatch (-want +got):\n%s", diff)
			}
		})
	}
}

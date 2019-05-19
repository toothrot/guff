package models

import (
	"bytes"
	"regexp"

	"github.com/PuerkitoBio/goquery"

	guff_proto "github.com/toothrot/guff/backend/generated"
)

var teamRex = regexp.MustCompile("leagues/([0-9]+)/teams/([0-9]+)")

type Team struct {
	ID         string
	Name       string
	DivisionID string
}

func (t *Team) ToProto() *guff_proto.Team {
	return &guff_proto.Team{
		Id: t.ID,
		Name: t.Name,
		DivisionId: t.DivisionID,
	}
}

func ParseTeams(b []byte) []Team {
	var teams []Team
	idTeams := map[string]Team{}
	doc, err := goquery.NewDocumentFromReader(bytes.NewReader(b))
	if err != nil {
		return teams
	}
	doc.Find(".game .team-score").Each(func(_ int, selection *goquery.Selection) {
		link := selection.Find("a").First()
		href, ok := link.Attr("href")
		if !ok {
			return
		}
		matches := teamRex.FindStringSubmatch(href)
		if len(matches) < 3 {
			return
		}
		if _, ok := idTeams[matches[2]]; ok {
			return
		}
		idTeams[matches[2]] = Team{
			DivisionID: matches[1],
			ID: matches[2],
			Name: link.Text(),
		}
	})
	for _, team := range idTeams {
		teams = append(teams, team)
	}
	return teams
}

// TeamByID implements sort.Interface for []Team based on ID
type TeamByID []Team

func (a TeamByID) Len() int           { return len(a) }
func (a TeamByID) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a TeamByID) Less(i, j int) bool { return a[i].ID < a[j].ID }

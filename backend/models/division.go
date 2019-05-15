package models

import (
	"bytes"
	"regexp"

	"github.com/PuerkitoBio/goquery"

	"github.com/toothrot/guff/backend/generated"
)

var divisionRex = regexp.MustCompile("leagues/([0-9]+)/schedule")
var titleRex = regexp.MustCompile("leagues/shuffleboard/([0-9]+)-")

type Division struct {
	ID   string
	Name string
}

func (d *Division) ToProto() *guff_proto.Division {
	return &guff_proto.Division{
		Id:   d.ID,
		Name: d.Name,
	}
}

func ParseDivisions(b []byte) []Division {
	var divisions []Division
	doc, err := goquery.NewDocumentFromReader(bytes.NewReader(b))
	if err != nil {
		return divisions
	}
	doc.Find(".meta-info").Each(func(i int, selection *goquery.Selection) {
		var d Division
		dLink := selection.Find("h2 a").First()
		href, ok := dLink.Attr("href")
		if !ok {
			return
		}
		matches := titleRex.FindStringSubmatch(href)
		if len(matches) < 2 {
			return
		}
		d.ID = matches[1]
		d.Name = dLink.Text()
		divisions = append(divisions, d)
	})
	return divisions
}

package models

import "regexp"

var divisionRex = regexp.MustCompile("leagues/([0-9]+)/schedule")

type Division struct {
	ID   string
	Name string
}

func ParseDivisions(b []byte) []Division {
	var divisions []Division
	for _, match := range divisionRex.FindAllSubmatch(b, -1) {
		divisions = append(divisions, Division{ID: string(match[1])})
	}
	return divisions
}

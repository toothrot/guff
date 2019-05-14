package models

import (
	"regexp"

	"github.com/toothrot/guff/backend/generated"
)

var divisionRex = regexp.MustCompile("leagues/([0-9]+)/schedule")

type Division struct {
	ID   string
	Name string
}

func (d *Division) ToProto() *guff_proto.Division {
	return &guff_proto.Division{
		Id: d.ID,
	}
}

func ParseDivisions(b []byte) []Division {
	var divisions []Division
	for _, match := range divisionRex.FindAllSubmatch(b, -1) {
		divisions = append(divisions, Division{ID: string(match[1])})
	}
	return divisions
}

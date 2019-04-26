package services

import "testing"

func TestTruth(t *testing.T) {
	if true != true {
		t.Errorf("true is true.")
	}
}

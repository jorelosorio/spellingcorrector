package spellingcorrector

import (
	"testing"
)

var spelling = NewSpelling("./dictionaries/es.dic")

func TestCorrection(t *testing.T) {
	if spelling.Correction("espanol") != "español" {
		t.Error("Expected returned string to be 'español'")
	}
}

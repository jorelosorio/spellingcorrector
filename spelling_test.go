package spellingcorrector

import (
	"testing"
)

var spelling = NewSpelling("./dictionaries/es.dic")

func testCorrection(t *testing.T) {
	if spelling.Correction("espanol") != "español" {
		t.Error("Expected returned string to be 'español'")
	}
}

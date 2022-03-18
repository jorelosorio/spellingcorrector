package spellingcorrector

import (
	"testing"
)

var spelling = NewSpelling("./dictionaries/es.dic")

func TestCorrection_FoundCorrection(t *testing.T) {
	if spelling.Correction("espanol") != "español" {
		t.Error("Expected returned string to be 'español'")
	}
}

func TestCorrection_CorrectWord(t *testing.T) {
	if spelling.Correction("español") != "español" {
		t.Error("Expected returned string to be 'español'")
	}
}

func TestCorrection_NotFound(t *testing.T) {
	if spelling.Correction("jorge") != "jorge" {
		t.Error("Expected returned string to be 'jorge'")
	}
}

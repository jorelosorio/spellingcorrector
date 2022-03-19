package spellingcorrector

import (
	"testing"
)

var spelling, _ = NewSpelling("./dictionaries/es.dic")

func TestCorrectedWord(t *testing.T) {
	if spelling.Correction("espanol") != "español" {
		t.Error("Expected returned string to be 'español'")
	}

	if spelling.Correction("aritocraticamente") != "aristocráticamente" {
		t.Error("Expected returned string to be 'socializar'")
	}
}

func TestExistingWord(t *testing.T) {
	if spelling.Correction("español") != "español" {
		t.Error("Expected returned string to be 'español'")
	}
}

func TestWordNotFound(t *testing.T) {
	if spelling.Correction("jorge") != "jorge" {
		t.Error("Expected returned string to be 'jorge'")
	}
}

func TestLongWordNotFound(t *testing.T) {
	if spelling.Correction("calculadorase") != "calculadorase" {
		t.Error("Expected returned string to be 'calculadorase'")
	}
}

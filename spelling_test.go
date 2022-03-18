package spellingcorrector

import (
	"testing"
)

var spelling, _ = NewSpelling("./dictionaries/es.dic")

func TestCorrection_CorrectedWords(t *testing.T) {
	if spelling.Correction("espanol") != "español" {
		t.Error("Expected returned string to be 'español'")
	}

	if spelling.Correction("aritocraticamente") != "aristocráticamente" {
		t.Error("Expected returned string to be 'socializar'")
	}
}

func TestCorrection_ExistingWord(t *testing.T) {
	if spelling.Correction("español") != "español" {
		t.Error("Expected returned string to be 'español'")
	}
}

func TestCorrection_WordNotFound(t *testing.T) {
	if spelling.Correction("jorge") != "jorge" {
		t.Error("Expected returned string to be 'jorge'")
	}
}

func TestCorrection_LongWordNotFound(t *testing.T) {
	if spelling.Correction("calculadorase") != "calculadorase" {
		t.Error("Expected returned string to be 'calculadorase'")
	}
}

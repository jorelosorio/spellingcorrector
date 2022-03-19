package spellingcorrector

import (
	"testing"
)

const (
	dictionaryFilePath = "./dictionaries/es.dic"
)

func newSpelling(t *testing.T) *Spelling {
	spelling, err := NewSpelling(dictionaryFilePath)
	if err != nil {
		t.Error(err)
	}

	return spelling
}

func TestCorrectedWord(t *testing.T) {
	spelling := newSpelling(t)

	if spelling.Correction("espanol") != "español" {
		t.Error("Expected returned string to be 'español'")
	}

	if spelling.Correction("aritocraticamente") != "aristocráticamente" {
		t.Error("Expected returned string to be 'socializar'")
	}
}

func TestExistingWord(t *testing.T) {
	spelling := newSpelling(t)

	if spelling.Correction("español") != "español" {
		t.Error("Expected returned string to be 'español'")
	}
}

func TestWordNotFound(t *testing.T) {
	spelling := newSpelling(t)

	if spelling.Correction("jorge") != "jorge" {
		t.Error("Expected returned string to be 'jorge'")
	}
}

func TestLongWordNotFound(t *testing.T) {
	spelling := newSpelling(t)

	if spelling.Correction("calculadorase") != "calculadorase" {
		t.Error("Expected returned string to be 'calculadorase'")
	}
}

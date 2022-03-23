package spellingcorrector

import (
	"os"
	"path"
	"testing"
)

func newSpelling(t *testing.T) *Spelling {
	currentPath, err := os.Getwd()
	if err != nil {
		t.Error(err)
	}

	dictionaryFilePath := path.Join(currentPath, "dictionaries", "es.dic")
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
		t.Error("Expected returned string to be 'aristocráticamente'")
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

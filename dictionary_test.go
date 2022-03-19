package spellingcorrector

import (
	"os"
	"path"
	"testing"
)

func newDictionary(t *testing.T) *Dictionary {
	currentPath, err := os.Getwd()
	if err != nil {
		t.Error(err)
	}

	dictionaryFilePath := path.Join(currentPath, "tmp", "es.dic")
	dic, err := NewDictionary(dictionaryFilePath, ESAlphabet)
	if err != nil {
		t.Error(err)
	}

	return dic
}

func TestTrain(t *testing.T) {
	dic := newDictionary(t)
	dic.TrainFromTextFile("./examples/test_sample1.txt")

	if len(dic.Words) != 2 {
		t.Error("Expected returned length should be 2")
	}
}

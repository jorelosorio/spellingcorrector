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

func TestLoadDictionary(t *testing.T) {
	currentPath, err := os.Getwd()
	if err != nil {
		t.Error(err)
	}

	dictionaryFilePath := path.Join(currentPath, "dictionaries", "wrong.dic")
	_, err = LoadDictionary(dictionaryFilePath)

	if err == nil {
		t.Error("If the dictionary file does not exist, then it should fail")
	}
}

func TestTrainDictionary(t *testing.T) {
	currentPath, err := os.Getwd()
	if err != nil {
		t.Error(err)
	}

	textFilePath := path.Join(currentPath, "examples", "test_sample1.txt")

	dic := newDictionary(t)
	dic.TrainFromTextFile(textFilePath)

	if len(dic.Words) != 2 {
		t.Error("Expected returned length should be 2")
	}
}

func TestTrainDictionaryWithWrongText(t *testing.T) {
	currentPath, err := os.Getwd()
	if err != nil {
		t.Error(err)
	}

	textFilePath := path.Join(currentPath, "examples", "wrong.txt")

	dic := newDictionary(t)
	err = dic.TrainFromTextFile(textFilePath)

	if err == nil {
		t.Error("If the text file does not exist, then it should fail")
	}
}

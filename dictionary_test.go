package spellingcorrector

import (
	"log"
	"os"
	"path"
	"testing"
)

func TestTrain(t *testing.T) {
	currentPath, err := os.Getwd()
	if err != nil {
		log.Println(err)
	}

	dictionaryFilePath := path.Join(currentPath, "tmp", "es.dic")
	dic, err := NewDictionary(dictionaryFilePath, ESAlphabet)
	if err != nil {
		t.Error(err)
	}

	dic.TrainFromTextFile("./examples/test_sample1.txt")

	if len(dic.Words) != 2 {
		t.Error("Expected returned length should be 2")
	}
}

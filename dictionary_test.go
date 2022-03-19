package spellingcorrector

import (
	"testing"
)

func TestTrain(t *testing.T) {
	dic, err := NewDictionary("./tmp/es.dic", ESAlphabet)
	if err != nil {
		t.Error(err)
	}

	dic.TrainFromTextFile("./examples/test_sample1.txt")

	if len(dic.Words) != 2 {
		t.Error("Expected returned length should be 2")
	}
}

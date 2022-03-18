package spellingcorrector

import (
	"testing"
)

func TestTrain(t *testing.T) {
	dic, _ := NewDictionary("./tmp/es.dic", ESAlphabet)
	dic.TrainFromTextFile("./examples/test_sample1.txt")

	if len(dic.Words) != 2 {
		t.Error("Expected returned length should be 2")
	}
}

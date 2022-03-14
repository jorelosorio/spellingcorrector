package helpers

import (
	"io/ioutil"
	"regexp"
	"strings"
)

func TrainFromFile(filePath string) {
	text, err := ioutil.ReadFile(filePath)
	if err != nil {
		panic(err)
	}

	dic := *LoadDictionary(filePath)

	pattern := regexp.MustCompile(Alphabet)
	for _, word := range pattern.FindAllString(strings.ToLower(string(text)), -1) {
		dic[word]++
	}

	StoreDictionary(&dic, filePath)
}

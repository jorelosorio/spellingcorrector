package spellingcorrector

import (
	"encoding/gob"
	"fmt"
	"io/ioutil"
	"os"
	"path"
	"regexp"
	"strings"
)

const (
	ESAlphabet = "abcdefghijklmnopqrstuvwxyzñáéíóúü"
	ENAlphabet = "abcdefghijklmnopqrstuvwxyz"
)

type Dictionary struct {
	Alphabet string
	Words    map[string]int
	filePath string
}

func NewDictionary(filePath string, alphabet string) (*Dictionary, error) {
	dic := Dictionary{Alphabet: alphabet, filePath: filePath, Words: make(map[string]int)}

	err := dic.save()
	if err != nil {
		return nil, err
	}

	return &dic, nil
}

func LoadDictionary(filePath string) (*Dictionary, error) {
	file, err := os.Open(filePath)
	if err != nil {
		return nil, err
	}

	dic := Dictionary{filePath: filePath}
	decoder := gob.NewDecoder(file)

	if err = decoder.Decode(&dic); err != nil {
		return nil, err
	}

	file.Close()

	return &dic, nil
}

func (d *Dictionary) TrainFromTextFile(textFilePath string) error {
	text, err := ioutil.ReadFile(textFilePath)
	if err != nil {
		return err
	}

	pattern := regexp.MustCompile(`[\p{L}]+`)
	words := pattern.FindAllString(strings.ToLower(string(text)), -1)

	fmt.Fprintln(os.Stdout, words, d)

	for _, word := range words {
		d.Words[word]++
	}

	d.save()

	return nil
}

func (d *Dictionary) save() error {
	// Create dictionary's directory if not exists
	dir := path.Dir(d.filePath)
	_ = os.Mkdir(dir, os.ModeDir)

	file, err := os.Create(d.filePath)
	if err != nil {
		return err
	}

	encoder := gob.NewEncoder(file)

	if err = encoder.Encode(*d); err != nil {
		return err
	}

	file.Close()

	return nil
}

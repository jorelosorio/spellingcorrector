package spellingcorrector

import (
	"encoding/gob"
	"io/ioutil"
	"os"
	"path"
	"regexp"
	"strings"
)

const (
	// Spanish alphabet
	ESAlphabet = "abcdefghijklmnopqrstuvwxyzñáéíóúü"
	// English alphabet
	ENAlphabet = "abcdefghijklmnopqrstuvwxyz"
)

type Dictionary struct {
	Alphabet string
	Words    map[string]int
	filePath string
}

// Creates a new dictionary file in the specified path and the alphabet that correspond to it.
// It returns a new Dictionary structure and any write error encountered.
func NewDictionary(filePath string, alphabet string) (*Dictionary, error) {
	dic := &Dictionary{Alphabet: alphabet, filePath: filePath, Words: make(map[string]int)}

	err := dic.save()
	if err != nil {
		return nil, err
	}

	return dic, nil
}

// Loads a dictionary in the specified file path
// It returns a new Dictionary structure and any read error encountered.
func LoadDictionary(filePath string) (*Dictionary, error) {
	file, err := os.Open(filePath)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	dic := &Dictionary{filePath: filePath, Words: make(map[string]int)}
	decoder := gob.NewDecoder(file)

	if err = decoder.Decode(dic); err != nil {
		return nil, err
	}

	return dic, nil
}

// It reads all the words that can be found in the text file specified path,
// those will be used to train the dictionary.
// It returns any read or write errors encountered.
func (d *Dictionary) TrainFromTextFile(textFilePath string) error {
	text, err := ioutil.ReadFile(textFilePath)
	if err != nil {
		return err
	}

	pattern := regexp.MustCompile(`[\p{L}]+`)
	words := pattern.FindAllString(strings.ToLower(string(text)), -1)

	for _, word := range words {
		d.Words[word]++
	}

	d.save()

	return nil
}

func (d *Dictionary) save() error {
	// Create dictionary's directory if not exists
	dir := path.Dir(d.filePath)

	os.Mkdir(dir, os.ModeDir)

	file, err := os.Create(d.filePath)
	if err != nil {
		return err
	}
	defer file.Close()

	encoder := gob.NewEncoder(file)

	if err = encoder.Encode(d); err != nil {
		return err
	}

	return nil
}

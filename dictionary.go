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
	// ESAlphabet represents the alphabet of the Spanish vocabulary.
	ESAlphabet = "abcdefghijklmnopqrstuvwxyzñáéíóúü"
	// ENAlphabet represents the alphabet of the English vocabulary.
	ENAlphabet = "abcdefghijklmnopqrstuvwxyz"
)

// Dictionary object is the main structure of the algorithm and it contains
// the alphabet in and the words.
type Dictionary struct {
	// Alphabet of the dictionary.
	Alphabet string
	// Words is a map that contains the word and the frequency number on texts,
	// it will help to calculate the most probable correction.
	Words    map[string]int
	filePath string
}

// NewDictionary creates a new dictionary file at the specified path and the alphabet that correspond to it.
// It returns a new Dictionary structure and any write error encountered.
func NewDictionary(filePath string, alphabet string) (*Dictionary, error) {
	dic := &Dictionary{Alphabet: alphabet, filePath: filePath, Words: make(map[string]int)}

	err := dic.save()
	if err != nil {
		return nil, err
	}

	return dic, nil
}

// LoadDictionary loads a dictionary in the specified file path
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

// TrainFromTextFile reads all the words that can be found in the text file specified path,
// those will be used to train the dictionary.
// It returns any read/write errors encountered.
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

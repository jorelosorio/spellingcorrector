package helpers

import (
	"encoding/gob"
	"os"
)

type Dictionary map[string]int

const Alphabet = "abcdefghijklmnopqrstuvwxyzáéíóúüñ"

func LoadDictionary(filePath string) *Dictionary {
	file, err := os.Open(filePath)

	dic := make(Dictionary)

	if err != nil && os.IsNotExist(err) {
		return &dic
	} else if err != nil {
		panic(err)
	}

	fileInfo, err := file.Stat()
	if err != nil {
		panic(err)
	}

	if fileInfo.Size() > 0 {
		decoder := gob.NewDecoder(file)

		if err = decoder.Decode(&dic); err != nil {
			panic(err)
		}

		file.Close()
	}

	return &dic
}

func StoreDictionary(dic *Dictionary, filePath string) {
	encodeFile, err := os.Create(filePath)
	if err != nil {
		panic(err)
	}

	encoder := gob.NewEncoder(encodeFile)

	if err = encoder.Encode(dic); err != nil {
		panic(err)
	}

	encodeFile.Close()
}

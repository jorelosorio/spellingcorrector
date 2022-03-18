package spellingcorrector

import (
	"github.com/jorelosorio/spellingcorrector/internals"
)

type Spelling struct {
	dic *internals.Dictionary
}

func (s *Spelling) edits1(word string) []string {
	var splits [][]string
	for i := 0; i < len(word)+1; i++ {
		splits = append(splits, []string{word[:i], word[i:]})
	}

	var words []string
	for _, v := range splits {
		l, r := v[0], v[1]
		lr := len(r)
		if lr > 0 {
			// Deletes
			words = append(words, l+r[1:])
		}
		if lr > 1 {
			// Transposes
			words = append(words, l+string(r[1])+string(r[0])+r[2:])
		}
		for _, c := range internals.Alphabet {
			if lr > 0 {
				// Replaces
				words = append(words, l+string(c)+r[1:])
			}

			// Inserts
			words = append(words, l+string(c)+r)
		}
	}

	return words
}

func (s *Spelling) edits2(word string) []string {
	var e2 []string
	for _, e1 := range s.edits1(word) {
		e2 = append(e2, s.edits1(e1)...)
	}

	return e2
}

func (s *Spelling) selectBestFor(word string, words []string) string {
	maxFreq := 0
	correction := ""
	for _, word := range words {
		if word == "" {
			break
		}

		if freq, present := (*s.dic)[word]; present && freq > maxFreq {
			maxFreq, correction = freq, word
		}
	}

	return correction
}

func NewSpelling(dicFilePath string) *Spelling {
	dic := internals.LoadDictionary(dicFilePath)
	return &Spelling{dic}
}

func (s *Spelling) Correction(word string) string {
	if _, present := (*s.dic)[word]; present {
		return word
	}

	if correction := s.selectBestFor(word, s.edits1(word)); correction != "" {
		return correction
	}

	if correction := s.selectBestFor(word, s.edits2(word)); correction != "" {
		return correction
	}

	return word
}

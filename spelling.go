package spellingcorrector

// Spelling object contains a dictionary object
// The main purpose of this struct is to provide actions/features that require process
// the dictionary data.
type Spelling struct {
	dic *Dictionary
}

// NewSpelling creates a new structure that contains a dictionary inside
// and it gets as a parameter the file path that points to the required dictionary.
// It returns a new Spelling structure and any read error encountered.
func NewSpelling(dicFilePath string) (*Spelling, error) {
	dic, err := LoadDictionary(dicFilePath)
	if err != nil {
		return nil, err
	}

	return &Spelling{dic}, nil
}

// Correction select the best possible correction for the specified word.
// Returns the correction if there was one.
func (s *Spelling) Correction(word string) string {
	if _, present := s.dic.Words[word]; present {
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
		for _, c := range s.dic.Alphabet {
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
		if freq, present := s.dic.Words[word]; present && freq > maxFreq {
			maxFreq, correction = freq, word
		}
	}

	return correction
}

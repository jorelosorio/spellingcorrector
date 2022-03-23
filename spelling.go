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

	words := make(chan string)
	go s.genAltsOf(word, words)

	if correction := s.findMaxFreqWord(words); correction != "" {
		return correction
	}

	expanedWords := make(chan string)
	go s.genAltsOf(word, words)

	for expandedWord := range words {
		if expandedWord == "" {
			break
		}

		go s.genAltsOf(expandedWord, expanedWords)
	}

	if correction := s.findMaxFreqWord(expanedWords); correction != "" {
		return correction
	}

	return word
}

func (s *Spelling) findMaxFreqWord(words chan string) string {
	maxFreq := 0
	correction := ""

	for word := range words {
		if word == "" {
			break
		}

		if freq, present := s.dic.Words[word]; present && freq > maxFreq {
			maxFreq, correction = freq, word
		}
	}

	return correction
}

func (s *Spelling) genAltsOf(word string, words chan string) {
	splits := [][]string{}
	for i := 0; i < len(word)+1; i++ {
		splits = append(splits, []string{word[:i], word[i:]})
	}

	for _, wordPair := range splits {
		l, r := wordPair[0], wordPair[1]
		lr := len(r)

		if lr > 0 {
			words <- l + r[1:]
		}

		if lr > 1 {
			words <- l + string(r[1]) + string(r[0]) + r[2:]
		}

		for _, c := range s.dic.Alphabet {
			if lr > 0 {
				words <- l + string(c) + r[1:]
			}

			words <- l + string(c) + r
		}
	}

	words <- ""
}

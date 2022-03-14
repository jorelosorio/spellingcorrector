package helpers

func edits1(word string) []string {
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
		for _, c := range Alphabet {
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

func edits2(word string) []string {
	var e2 []string
	for _, e1 := range edits1(word) {
		e2 = append(e2, edits1(e1)...)
	}

	return e2
}

func selectBestFor(word string, words []string, dic *Dictionary) string {
	maxFreq := 0
	correction := ""
	for _, word := range words {
		if word == "" {
			break
		}

		if freq, present := (*dic)[word]; present && freq > maxFreq {
			maxFreq, correction = freq, word
		}
	}

	return correction
}

func Correction(word string, dic *Dictionary) string {
	if _, present := (*dic)[word]; present {
		return word
	}

	if correction := selectBestFor(word, edits1(word), dic); correction != "" {
		return correction
	}

	if correction := selectBestFor(word, edits2(word), dic); correction != "" {
		return correction
	}

	return word
}

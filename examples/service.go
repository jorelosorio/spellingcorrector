package main

import (
	"net/http"
	"os"

	sc "github.com/jorelosorio/spellingcorrector"
)

func main() {
	// Dictionary file path
	dictionaryFilePath := os.Args[1]

	spelling, err := sc.NewSpelling(dictionaryFilePath)
	if err != nil {
		panic(err)
	}

	http.HandleFunc("/spelling", func(w http.ResponseWriter, r *http.Request) {
		word := r.URL.Query().Get("word")
		correction := spelling.Correction(word)
		_, _ = w.Write([]byte(correction))
	})

	err = http.ListenAndServe(":80", nil)
	if err != nil {
		panic(err)
	}
}

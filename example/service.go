package main

import (
	"net/http"
	"os"

	sc "github.com/jorelosorio/spellingcorrector"
)

func main() {
	dictionaryFilePath := os.Args[1]
	spelling := sc.NewSpelling(dictionaryFilePath)

	http.HandleFunc("/spelling", func(w http.ResponseWriter, r *http.Request) {
		word := r.URL.Query().Get("word")
		correction := spelling.Correction(word)
		_, _ = w.Write([]byte(correction))
	})

	err := http.ListenAndServe(":80", nil)
	if err != nil {
		panic(err)
	}
}

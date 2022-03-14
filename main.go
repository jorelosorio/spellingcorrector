package main

import (
	"net/http"
	"spelling-corrector/helpers"
)

func main() {
	dic := helpers.LoadDictionary("./dic")

	http.HandleFunc("/spelling", func(w http.ResponseWriter, r *http.Request) {
		word := r.URL.Query().Get("word")
		correction := helpers.Correction(word, dic)
		_, _ = w.Write([]byte(correction))
	})

	err := http.ListenAndServe(":80", nil)
	if err != nil {
		panic(err)
	}
}

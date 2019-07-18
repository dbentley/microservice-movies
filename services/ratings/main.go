package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello ratings")
	})
	http.HandleFunc("/rating", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "GET,PUT,POST,DELETE")
		titles := r.URL.Query()["title"]
		title := ""
		if len(titles) > 0 {
			title = titles[0]
		}
		fmt.Fprintf(w, "%d", rateMovie(title))
	})

	http.ListenAndServe(":80", nil)
}

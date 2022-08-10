package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

func writeJsonSearchResults(w http.ResponseWriter, r *http.Request, searchResults any) {
	bytes, err := json.Marshal(searchResults)
	if err != nil {
		log.Fatal(err)
	}
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	_, err = fmt.Fprint(w, string(bytes))
	if err != nil {
		log.Fatal(err)
	}
}

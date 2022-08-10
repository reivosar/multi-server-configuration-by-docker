package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	mappingEndpointsToHandlers()
	log.Fatal(http.ListenAndServe(":80", nil))
}

func mappingEndpointsToHandlers() {
	http.HandleFunc("/", indexHandler)
	http.HandleFunc("/users", userReadAPIHandler)
}

func indexHandler(w http.ResponseWriter, r *http.Request) {
	_, err := fmt.Fprint(w, "read-api")
	if err != nil {
		return
	}
}

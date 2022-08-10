package main

import (
	"net/http"

	_ "github.com/go-sql-driver/mysql"
)

func userReadAPIHandler(w http.ResponseWriter, r *http.Request) {
	writeToResponseWithJsonSearchResults(w, r, getAllUsers())
}

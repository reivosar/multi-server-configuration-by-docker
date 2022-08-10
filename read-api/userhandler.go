package main

import (
	"net/http"

	_ "github.com/go-sql-driver/mysql"
)

func userReadAPIHandler(w http.ResponseWriter, r *http.Request) {
	writeJsonSearchResults(w, r, getAllUsers())
}

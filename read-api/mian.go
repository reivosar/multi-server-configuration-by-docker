package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"

	_ "github.com/go-sql-driver/mysql"
)

func indexHandler(w http.ResponseWriter, r *http.Request) {
	_, err := fmt.Fprint(w, "read-api")
	if err != nil {
		return
	}
}

func main() {
	http.HandleFunc("/", indexHandler)
	http.HandleFunc("/users", userSearchAllIndexHandler)
	log.Fatal(http.ListenAndServe(":80", nil))
}

type User struct {
	id       int    `json:"id"`
	username string `json:"username"`
	email    string `json:"email"`
}

func userSearchAllIndexHandler(w http.ResponseWriter, r *http.Request) {
	users := selectAll()
	bytes, err := json.Marshal(users)
	if err != nil {
		log.Fatal(err)
	}
	_, err = w.Write([]byte(string(bytes)))
	if err != nil {
		log.Fatal(err)
	}
}

func selectAll() []User {
	db := Connect()
	defer db.Close()

	rows, err := db.Query("SELECT id, username, email FROM users")
	if err != nil {
		panic(err.Error())
	}

	results := make([]User, 0)
	for rows.Next() {
		var user User
		err = rows.Scan(&user)
		if err != nil {
			panic(err.Error())
		}
		results = append(results, user)
	}
	return results
}

func Connect() *sql.DB {
	user := os.Getenv("MYSQL_USER")
	pw := os.Getenv("MYSQL_PASSWORD")
	db_name := os.Getenv("MYSQL_DATABASE")
	var path string = fmt.Sprintf("%s:%s@tcp(db-container:3306)/%s?charset=utf8", user, pw, db_name)
	log.Print(path)
	var db, err = sql.Open("mysql", path)
	if err != nil {
		panic(err.Error())
	}
	return db
}

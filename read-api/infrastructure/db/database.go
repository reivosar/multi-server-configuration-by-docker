package db

import (
	"fmt"
	"os"
	_ "github.com/go-sql-driver/mysql"
)

var Db *sql.DB


func Connect() *sql.DB {
	user := os.Getenv("MYSQL_USER")
	pw := os.Getenv("MYSQL_PASSWORD")
	db_name := os.Getenv("MYSQL_DATABASE")
	var path string = fmt.Sprintf("%s:%s@tcp(db-container:3306)/%s?charset=utf8&parseTime=true", user, pw, db_name)
	if db, err = sql.Open("mysql", path); 
	err != nil {
		panic(err.Error())
	}
    return db
}
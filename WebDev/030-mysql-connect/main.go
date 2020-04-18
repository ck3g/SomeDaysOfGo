package main

import (
	"database/sql"
	"fmt"
	"io"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
)

var db *sql.DB
var err error

func main() {
	// $ mysql
	// mysql> CREATE DATABASE goplayground;
	db, err := sql.Open("mysql", "root@/goplayground")
	handleError(err)
	defer db.Close()

	err = db.Ping()
	handleError(err)

	http.HandleFunc("/", index)
	http.Handle("/favicon.ico", http.NotFoundHandler())
	err = http.ListenAndServe(":8080", nil)
	handleError(err)
}

func index(w http.ResponseWriter, req *http.Request) {
	_, err = io.WriteString(w, "Successfully completed.")
	handleError(err)
}

func handleError(err error) {
	if err != nil {
		fmt.Println(err)
	}
}

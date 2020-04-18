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
	db, err = sql.Open("mysql", "root@/goplayground")
	handleError(err)
	defer db.Close()

	err = db.Ping()
	handleError(err)

	http.HandleFunc("/", index)
	http.HandleFunc("/create", create)
	http.HandleFunc("/insert", insert)
	http.HandleFunc("/read", read)
	http.HandleFunc("/update", update)
	http.HandleFunc("/delete", del)
	http.HandleFunc("/drop", drop)
	http.Handle("/favicon.ico", http.NotFoundHandler())
	err = http.ListenAndServe(":8080", nil)
	handleError(err)
}

func index(w http.ResponseWriter, req *http.Request) {
	_, err = io.WriteString(w, "Index page.")
	handleError(err)
}

func handleError(err error) {
	if err != nil {
		fmt.Println(err)
	}
}

func create(w http.ResponseWriter, req *http.Request) {
	statement, err := db.Prepare(`CREATE TABLE people (name VARCHAR(20));`)
	handleError(err)
	defer statement.Close()

	r, err := statement.Exec()
	handleError(err)

	n, err := r.RowsAffected()
	handleError(err)

	fmt.Fprintln(w, "CREATED TABLE people", n)
}

func insert(w http.ResponseWriter, req *http.Request) {
	statement, err := db.Prepare(`INSERT INTO people (name) VALUES ("Alice"), ("Bob"), ("John"), ("Walter");`)
	handleError(err)
	defer statement.Close()

	r, err := statement.Exec()
	handleError(err)

	n, err := r.RowsAffected()
	handleError(err)

	fmt.Fprintln(w, "INSERTED RECORDS", n)
}

func read(w http.ResponseWriter, req *http.Request) {
	rows, err := db.Query(`SELECT name FROM people;`)
	handleError(err)
	defer rows.Close()

	// data to be used in query
	var s, name string
	s = "RETRIEVED RECORDS:\n"

	// query
	for rows.Next() {
		err = rows.Scan(&name)
		handleError(err)
		s += name + "\n"
	}

	fmt.Fprintln(w, s)
}

func update(w http.ResponseWriter, req *http.Request) {
	statement, err := db.Prepare(`UPDATE people SET name = "Jimmy" WHERE name = "Bob";`)
	handleError(err)
	defer statement.Close()

	r, err := statement.Exec()
	handleError(err)

	n, err := r.RowsAffected()
	handleError(err)

	fmt.Fprintln(w, "UPDATED RECORD", n)
}

func del(w http.ResponseWriter, req *http.Request) {
	statement, err := db.Prepare(`DELETE FROM people WHERE name = "Jimmy";`)
	handleError(err)
	defer statement.Close()

	r, err := statement.Exec()
	handleError(err)

	n, err := r.RowsAffected()
	handleError(err)

	fmt.Fprintln(w, "DELETED RECORD", n)
}

func drop(w http.ResponseWriter, req *http.Request) {
	statement, err := db.Prepare(`DROP TABLE people;`)
	handleError(err)
	defer statement.Close()

	_, err = statement.Exec()
	handleError(err)

	fmt.Fprintln(w, "DROPPED TABLE people")
}

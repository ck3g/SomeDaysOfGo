package config

import (
	"database/sql"
	"fmt"

	// import Postgres adapter
	_ "github.com/lib/pq"
)

// DB provides a connection to the database
var DB *sql.DB

func init() {
	var err error
	DB, err = sql.Open("postgres", "postgres://localhost/go_bookstore?sslmode=disable")
	if err != nil {
		panic(err)
	}

	if err = DB.Ping(); err != nil {
		panic(err)
	}
	fmt.Println("You connected to your database.")
}

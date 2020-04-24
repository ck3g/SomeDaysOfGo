package main

import (
	"database/sql"
	"net/http"
	"strconv"

	"github.com/ck3g/SomeDaysOfGo/WebDev/038-postgres/config"
)

// Book model
type Book struct {
	Isbn   string
	Title  string
	Author string
	Price  float32
}

func main() {
	http.HandleFunc("/", index)
	http.HandleFunc("/books", booksIndex)
	http.HandleFunc("/books/show", bookShow)
	http.HandleFunc("/books/create", bookCreateForm)
	http.HandleFunc("/books/create/process", bookCreateProcess)
	http.HandleFunc("/books/update", bookUpdateForm)
	http.HandleFunc("/books/update/process", bookUpdateProcess)
	http.HandleFunc("/books/delete/process", bookDeleteProcess)
	http.ListenAndServe(":8080", nil)
}

func index(w http.ResponseWriter, r *http.Request) {
	http.Redirect(w, r, "/books", http.StatusSeeOther)
}

func booksIndex(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		http.Error(w, http.StatusText(405), http.StatusMethodNotAllowed)
		return
	}

	rows, err := config.DB.Query("SELECT * FROM books;")
	if err != nil {
		http.Error(w, http.StatusText(500), 500)
		return
	}
	defer rows.Close()

	books := make([]Book, 0)
	for rows.Next() {
		book := Book{}
		err := rows.Scan(&book.Isbn, &book.Title, &book.Author, &book.Price) // order matters
		if err != nil {
			http.Error(w, http.StatusText(500), 500)
			return
		}
		books = append(books, book)
	}

	if err = rows.Err(); err != nil {
		http.Error(w, http.StatusText(500), 500)
		return
	}

	config.TPL.ExecuteTemplate(w, "books.gohtml", books)
}

func bookShow(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		http.Error(w, http.StatusText(400), http.StatusBadRequest)
		return
	}

	isbn := r.FormValue("isbn")
	if isbn == "" {
		http.Error(w, http.StatusText(400), http.StatusBadRequest)
		return
	}

	row := config.DB.QueryRow("SELECT * FROM books WHERE isbn = $1", isbn)

	book := Book{}
	err := row.Scan(&book.Isbn, &book.Title, &book.Author, &book.Price)
	switch {
	case err == sql.ErrNoRows:
		http.NotFound(w, r)
		return
	case err != nil:
		http.Error(w, http.StatusText(500), http.StatusInternalServerError)
		return
	}

	config.TPL.ExecuteTemplate(w, "show.gohtml", book)
}

func bookCreateForm(w http.ResponseWriter, r *http.Request) {
	config.TPL.ExecuteTemplate(w, "create.gohtml", nil)
}

func bookCreateProcess(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		http.Error(w, http.StatusText(405), http.StatusMethodNotAllowed)
		return
	}

	// get form values
	book := Book{}
	book.Isbn = r.FormValue("isbn")
	book.Title = r.FormValue("title")
	book.Author = r.FormValue("author")
	price := r.FormValue("price")

	// validate form values
	if book.Isbn == "" || book.Title == "" || book.Author == "" || price == "" {
		http.Error(w, http.StatusText(400), http.StatusBadRequest)
		return
	}

	// convert form values
	f64, err := strconv.ParseFloat(price, 32)
	if err != nil {
		http.Error(w, http.StatusText(406)+"Please hit back and enter a number for the price", http.StatusNotAcceptable)
		return
	}
	book.Price = float32(f64)

	// insert values
	_, err = config.DB.Exec(
		"INSERT INTO books (isbn, title, author, price) VALUES ($1, $2, $3, $4)",
		book.Isbn, book.Title, book.Author, book.Price,
	)
	if err != nil {
		http.Error(w, http.StatusText(500), http.StatusInternalServerError)
		return
	}

	// confirm insertion
	config.TPL.ExecuteTemplate(w, "created.gohtml", book)
}

func bookUpdateForm(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		http.Error(w, http.StatusText(405), http.StatusMethodNotAllowed)
		return
	}

	isbn := r.FormValue("isbn")
	if isbn == "" {
		http.Error(w, http.StatusText(400), http.StatusBadRequest)
		return
	}

	row := config.DB.QueryRow("SELECT * FROM books WHERE isbn = $1", isbn)

	book := Book{}
	err := row.Scan(&book.Isbn, &book.Title, &book.Author, &book.Price)
	switch {
	case err == sql.ErrNoRows:
		http.NotFound(w, r)
		return
	case err != nil:
		http.Error(w, http.StatusText(500), http.StatusInternalServerError)
		return
	}

	config.TPL.ExecuteTemplate(w, "update.gohtml", book)
}

func bookUpdateProcess(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		http.Error(w, http.StatusText(405), http.StatusMethodNotAllowed)
		return
	}

	// get form values
	// get form values
	book := Book{}
	book.Isbn = r.FormValue("isbn")
	book.Title = r.FormValue("title")
	book.Author = r.FormValue("author")
	p := r.FormValue("price")

	// validate form values
	if book.Isbn == "" || book.Title == "" || book.Author == "" || p == "" {
		http.Error(w, http.StatusText(400), http.StatusBadRequest)
		return
	}

	// convert form values
	f64, err := strconv.ParseFloat(p, 32)
	if err != nil {
		http.Error(w, http.StatusText(406)+"Please hit back and enter a number for the price", http.StatusNotAcceptable)
		return
	}
	book.Price = float32(f64)

	// insert values
	_, err = config.DB.Exec(
		"UPDATE books SET isbn=$1, title=$2, author=$3, price=$4 WHERE isbn=$1;",
		book.Isbn, book.Title, book.Author, book.Price,
	)
	if err != nil {
		http.Error(w, http.StatusText(500), http.StatusInternalServerError)
		return
	}

	// confirm update
	config.TPL.ExecuteTemplate(w, "updated.gohtml", book)
}

func bookDeleteProcess(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		http.Error(w, http.StatusText(405), http.StatusMethodNotAllowed)
		return
	}

	isbn := r.FormValue("isbn")
	if isbn == "" {
		http.Error(w, http.StatusText(400), http.StatusBadRequest)
		return
	}

	// delete book
	_, err := config.DB.Exec("DELETE FROM books WHERE isbn=$1;", isbn)
	if err != nil {
		http.Error(w, http.StatusText(500), http.StatusInternalServerError)
		return
	}

	http.Redirect(w, r, "/books", http.StatusSeeOther)
}

package books

import (
	"errors"
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

// All fetches all books
func All() ([]Book, error) {
	rows, err := config.DB.Query("SELECT * FROM books;")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	books := make([]Book, 0)
	for rows.Next() {
		book := Book{}
		err := rows.Scan(&book.Isbn, &book.Title, &book.Author, &book.Price) // order matters
		if err != nil {
			return nil, err
		}
		books = append(books, book)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return books, nil
}

// One fetches a single book
func One(r *http.Request) (Book, error) {
	book := Book{}
	isbn := r.FormValue("isbn")
	if isbn == "" {
		return book, errors.New("400. Bad Request")
	}

	row := config.DB.QueryRow("SELECT * FROM books WHERE isbn = $1", isbn)

	err := row.Scan(&book.Isbn, &book.Title, &book.Author, &book.Price)
	if err != nil {
		return book, err
	}

	return book, err
}

// Create inserts a new book into DB
func Create(r *http.Request) (Book, error) {
	// get form values
	book := Book{}
	book.Isbn = r.FormValue("isbn")
	book.Title = r.FormValue("title")
	book.Author = r.FormValue("author")
	price := r.FormValue("price")

	// validate form values
	if book.Isbn == "" || book.Title == "" || book.Author == "" || price == "" {
		return book, errors.New("400. Bad Request")
	}

	// convert form values
	f64, err := strconv.ParseFloat(price, 32)
	if err != nil {
		return book, errors.New("406. Not Acceptable. Enter number for price")
	}
	book.Price = float32(f64)

	// insert values
	_, err = config.DB.Exec(
		"INSERT INTO books (isbn, title, author, price) VALUES ($1, $2, $3, $4)",
		book.Isbn, book.Title, book.Author, book.Price,
	)
	if err != nil {
		return book, errors.New("500. Internal Server Error")
	}

	return book, nil
}

// Update amends a book
func Update(r *http.Request) (Book, error) {
	// get form values
	book := Book{}
	book.Isbn = r.FormValue("isbn")
	book.Title = r.FormValue("title")
	book.Author = r.FormValue("author")
	p := r.FormValue("price")

	// validate form values
	if book.Isbn == "" || book.Title == "" || book.Author == "" || p == "" {
		return book, errors.New("400. Bad Request. Fields can't be empty")
	}

	// convert form values
	f64, err := strconv.ParseFloat(p, 32)
	if err != nil {
		return book, errors.New("406. Not Acceptable. Enter number for price")
	}
	book.Price = float32(f64)

	// insert values
	_, err = config.DB.Exec(
		"UPDATE books SET isbn=$1, title=$2, author=$3, price=$4 WHERE isbn=$1;",
		book.Isbn, book.Title, book.Author, book.Price,
	)
	if err != nil {
		return book, err
	}

	return book, nil
}

// Delete - deletes a book
func Delete(r *http.Request) error {
	isbn := r.FormValue("isbn")
	if isbn == "" {
		return errors.New("400. Bad Request")
	}

	// delete book
	_, err := config.DB.Exec("DELETE FROM books WHERE isbn=$1;", isbn)
	if err != nil {
		return errors.New("500. Internal Server Error")
	}

	return nil
}

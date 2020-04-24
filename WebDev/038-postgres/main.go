package main

import (
	"database/sql"
	"net/http"

	"github.com/ck3g/SomeDaysOfGo/WebDev/038-postgres/books"
	"github.com/ck3g/SomeDaysOfGo/WebDev/038-postgres/config"
)

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

	bks, err := books.All()
	if err != nil {
		http.Error(w, http.StatusText(500), http.StatusInternalServerError)
	}

	config.TPL.ExecuteTemplate(w, "books.gohtml", bks)
}

func bookShow(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		http.Error(w, http.StatusText(400), http.StatusBadRequest)
		return
	}

	book, err := books.One(r)
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

	book, err := books.Create(r)
	if err != nil {
		http.Error(w, http.StatusText(406), http.StatusNotAcceptable)
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

	book, err := books.One(r)
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

	book, err := books.Update(r)
	if err != nil {
		http.Error(w, http.StatusText(406), http.StatusNotAcceptable)
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

	err := books.Delete(r)
	if err != nil {
		http.Error(w, http.StatusText(400), http.StatusBadRequest)
	}

	http.Redirect(w, r, "/books", http.StatusSeeOther)
}

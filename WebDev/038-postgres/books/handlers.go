package books

import (
	"database/sql"
	"net/http"

	"github.com/ck3g/SomeDaysOfGo/WebDev/038-postgres/config"
)

// Index handler for list of books
func Index(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		http.Error(w, http.StatusText(405), http.StatusMethodNotAllowed)
		return
	}

	bks, err := All()
	if err != nil {
		http.Error(w, http.StatusText(500), http.StatusInternalServerError)
	}

	config.TPL.ExecuteTemplate(w, "books.gohtml", bks)
}

// Show handler for show a book
func Show(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		http.Error(w, http.StatusText(400), http.StatusBadRequest)
		return
	}

	book, err := One(r)
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

// CreateForm route to render edit form
func CreateForm(w http.ResponseWriter, r *http.Request) {
	config.TPL.ExecuteTemplate(w, "create.gohtml", nil)
}

// CreateProcess - create book action
func CreateProcess(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		http.Error(w, http.StatusText(405), http.StatusMethodNotAllowed)
		return
	}

	book, err := Create(r)
	if err != nil {
		http.Error(w, http.StatusText(406), http.StatusNotAcceptable)
		return
	}

	// confirm insertion
	config.TPL.ExecuteTemplate(w, "created.gohtml", book)
}

// UpdateForm - action to edit a book
func UpdateForm(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		http.Error(w, http.StatusText(405), http.StatusMethodNotAllowed)
		return
	}

	book, err := One(r)
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

// UpdateProcess - handler to process book update
func UpdateProcess(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		http.Error(w, http.StatusText(405), http.StatusMethodNotAllowed)
		return
	}

	book, err := Update(r)
	if err != nil {
		http.Error(w, http.StatusText(406), http.StatusNotAcceptable)
		return
	}

	// confirm update
	config.TPL.ExecuteTemplate(w, "updated.gohtml", book)
}

// DeleteProcess - handler to delete a book
func DeleteProcess(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		http.Error(w, http.StatusText(405), http.StatusMethodNotAllowed)
		return
	}

	err := Delete(r)
	if err != nil {
		http.Error(w, http.StatusText(400), http.StatusBadRequest)
	}

	http.Redirect(w, r, "/books", http.StatusSeeOther)
}

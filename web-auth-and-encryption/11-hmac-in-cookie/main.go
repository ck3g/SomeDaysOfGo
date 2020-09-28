package main

import (
	"io"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", home)
	http.HandleFunc("/submit", submit)
	http.ListenAndServe(":8080", nil)
}

func getCode(msg string) string {
	return ""
}

func submit(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Redirect(w, r, "/", http.StatusSeeOther)
		return
	}

	email := r.FormValue("email")
	if email == "" {
		http.Redirect(w, r, "/", http.StatusSeeOther)
		return
	}

	c := http.Cookie{
		Name:  "session",
		Value: "",
	}

	log.Println(c)
}

func home(w http.ResponseWriter, r *http.Request) {
	html := `
	<!DOCTYPE html>
	<html lang="en">
		<head>
			<title>HMAC - example</title>
		</head>
		<body>
			<form action="/submit" method="post">
				<input type="text" name="email" />
				<input type="submit" />
			</form>
		</body>
	</html>
	`
	io.WriteString(w, html)
}

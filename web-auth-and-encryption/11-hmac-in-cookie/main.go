package main

import (
	"crypto/hmac"
	"crypto/sha256"
	"fmt"
	"io"
	"net/http"
)

func main() {
	http.HandleFunc("/", home)
	http.HandleFunc("/submit", submit)
	http.ListenAndServe(":8080", nil)
}

func getCode(msg string) string {
	h := hmac.New(sha256.New, []byte("a very secret key"))
	h.Write([]byte(msg))

	return fmt.Sprintf("%x", h.Sum(nil))
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

	code := getCode(email)

	// "hash / message digest / digest / hash value" | "what we stored"
	c := http.Cookie{
		Name:  "session",
		Value: code + "|" + email,
	}
	http.SetCookie(w, &c)
	http.Redirect(w, r, "/", http.StatusSeeOther)
}

func home(w http.ResponseWriter, r *http.Request) {
	c, err := r.Cookie("session")
	if err != nil {
		c = &http.Cookie{}
	}

	html := `
	<!DOCTYPE html>
	<html lang="en">
		<head>
			<title>HMAC - example</title>
		</head>
		<body>
			<p> Cookie value: ` + c.Value + `</p>
			<form action="/submit" method="post">
				<input type="text" name="email" />
				<input type="submit" />
			</form>
		</body>
	</html>
	`
	io.WriteString(w, html)
}

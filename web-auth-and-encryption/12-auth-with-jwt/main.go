package main

import (
	"fmt"
	"io"
	"net/http"
	"time"

	"github.com/dgrijalva/jwt-go"
)

func main() {
	http.HandleFunc("/", home)
	http.HandleFunc("/submit", submit)
	http.ListenAndServe(":8080", nil)
}

func getJWT(msg string) (string, error) {
	myKey := "a very secret key"

	type myClaims struct {
		jwt.StandardClaims
		Email string
	}

	claims := myClaims{
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(5 * time.Minute).Unix(),
		},
		Email: msg,
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, &claims)
	ss, err := token.SignedString([]byte(myKey))
	if err != nil {
		return "", fmt.Errorf("coulnd't SignedString %w", err)
	}

	return ss, nil
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

	ss, err := getJWT(email)
	if err != nil {
		http.Error(w, "couldn't getJWT", http.StatusInternalServerError)
	}

	// "hash / message digest / digest / hash value" | "what we stored"
	c := http.Cookie{
		Name:  "session",
		Value: ss,
	}
	http.SetCookie(w, &c)
	http.Redirect(w, r, "/", http.StatusSeeOther)
}

func home(w http.ResponseWriter, r *http.Request) {
	c, err := r.Cookie("session")
	if err != nil {
		c = &http.Cookie{}
	}

	isEqual := false

	message := "Not logged in"
	if isEqual {
		message = "Logged in"
	}

	html := `
	<!DOCTYPE html>
	<html lang="en">
		<head>
			<title>HMAC - example</title>
		</head>
		<body>
			<p> Cookie value: ` + c.Value + `</p>
			<p>` + message + `</p>
			<form action="/submit" method="post">
				<input type="text" name="email" />
				<input type="submit" />
			</form>
		</body>
	</html>
	`
	io.WriteString(w, html)
}

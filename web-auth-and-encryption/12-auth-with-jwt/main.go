package main

import (
	"fmt"
	"io"
	"net/http"
	"time"

	"github.com/dgrijalva/jwt-go"
)

type myClaims struct {
	jwt.StandardClaims
	Email string
}

const myKey = "a very secret key"

func main() {
	http.HandleFunc("/", home)
	http.HandleFunc("/submit", submit)
	http.ListenAndServe(":8080", nil)
}

func getJWT(msg string) (string, error) {

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

	ss := c.Value
	afterVerificationToken, err := jwt.ParseWithClaims(ss, &myClaims{}, func(beforeVerificationToken *jwt.Token) (interface{}, error) {
		if beforeVerificationToken.Method.Alg() != jwt.SigningMethodHS256.Alg() {
			return nil, fmt.Errorf("SOMEONE TRIED TO HACK changed signing method")
		}
		return []byte(myKey), nil
	})

	// StandardClaims has the `Valid()` error method which means it implements `Claims` interface
	// ...when you ParseWithClaims the `Valid()` gets run and the `token.Valid` field will be `true`

	isEqual := err == nil && afterVerificationToken.Valid
	message := "Not logged in"
	if isEqual {
		message = "Logged in"
		claims := afterVerificationToken.Claims.(*myClaims)
		fmt.Println(claims.Email)
		fmt.Println(claims.ExpiresAt)
	}

	html := `
	<!DOCTYPE html>
	<html lang="en">
		<head>
			<title>JWT - example</title>
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

package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/ck3g/SomeDaysOfGo/WebDev/036-http-service-with-mongodb/models"
	"github.com/julienschmidt/httprouter"
)

func main() {
	r := httprouter.New()
	r.GET("/", index)
	r.GET("/user/:id", getUser)
	http.ListenAndServe("localhost:8080", r)
}

func index(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	s := `<!DOCTYPE html>
<html lang="en">
<head>
<meta charset="UTF-8">
<title>Index</title>
</head>
<body>
<a href="/user/9872309847">GO TO: /user/9872309847</a>
</body>
</html>`

	w.Header().Set("Content-Type", "text/html; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(s))
}

func getUser(w http.ResponseWriter, req *http.Request, p httprouter.Params) {
	user := models.User{
		Name:   "John Doe",
		Gender: "male",
		Age:    32,
		ID:     p.ByName("id"),
	}

	userJSON, err := json.Marshal(user)
	if err != nil {
		fmt.Println(err)
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "%s\n", userJSON)
}

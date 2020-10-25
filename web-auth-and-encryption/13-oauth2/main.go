package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/", index)
	http.HandleFunc("/oauth/github", startGitHubOAuth)
	http.ListenAndServe(":8080", nil)
}

func index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, `<!DOCTYPE html>
<html lang="en">
<head>
  <meta charset="UTF-8">
  <title>Document</title>
<head>
<body>
  <form action="/oauth/github" method="post">
    <input type="submit" value="Login with GitHub">
  </form>
</body>
</html>`)
}

func startGitHubOAuth(w http.ResponseWriter, r *http.Request) {
}

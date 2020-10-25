package main

import (
	"flag"
	"fmt"
	"net/http"
	"os"

	"golang.org/x/oauth2"
	"golang.org/x/oauth2/github"
)

var githubOauthConfig = &oauth2.Config{
	Endpoint: github.Endpoint,
}

func main() {
	clientID := flag.String("clientid", "", "GitHub OAuth2 ClientID (Required)")
	secret := flag.String("secret", "", "GitHub OAuth2 Client Secret (Required)")
	flag.Parse()

	if *clientID == "" || *secret == "" {
		fmt.Println("`clientid` and `secret` must be set")
		flag.PrintDefaults()
		os.Exit(1)
	}

	githubOauthConfig.ClientID = *clientID
	githubOauthConfig.ClientSecret = *secret

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
	// 0000 - is a fake ID for loging attempts
	redirectURL := githubOauthConfig.AuthCodeURL("0000")
	http.Redirect(w, r, redirectURL, http.StatusSeeOther)
}

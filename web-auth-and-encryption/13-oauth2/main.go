package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strings"

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
	http.HandleFunc("/oauth2/github", startGitHubOAuth)
	http.HandleFunc("/oauth2/callback", completeGitHubOAuth)
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
  <form action="/oauth2/github" method="post">
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

func completeGitHubOAuth(w http.ResponseWriter, r *http.Request) {
	code := r.FormValue("code")
	state := r.FormValue("state")

	// Checking agains the fake state set in `startGitHubOAuth`
	if state != "0000" {
		http.Error(w, "State is incorrect", http.StatusBadRequest)
		return
	}

	token, err := githubOauthConfig.Exchange(r.Context(), code)
	if err != nil {
		http.Error(w, "Couldn't login", http.StatusInternalServerError)
		return
	}

	tokenSource := githubOauthConfig.TokenSource(r.Context(), token)
	client := oauth2.NewClient(r.Context(), tokenSource)

	requestBody := strings.NewReader(`{"query": "query {viewer {id login name}}"`)
	resp, err := client.Post("https://api.github.com/graphql", "application/json", requestBody)
	if err != nil {
		http.Error(w, "Couldn't get user", http.StatusInternalServerError)
		return
	}
	defer resp.Body.Close()

	bs, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		http.Error(w, "Couldn't read GitHub information", http.StatusInternalServerError)
		return
	}

	log.Println(string(bs))
}

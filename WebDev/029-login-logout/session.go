package main

import "net/http"

func getUser(req *http.Request) user {
	var usr user

	// get cookie
	cookie, err := req.Cookie("session")
	if err != nil {
		return usr
	}

	// if the user exists already, get user
	if username, ok := dbSessions[cookie.Value]; ok {
		usr = dbUsers[username]
	}

	return usr
}

func alreadyLoggedIn(req *http.Request) bool {
	cookie, err := req.Cookie("session")
	if err != nil {
		return false
	}

	username := dbSessions[cookie.Value]
	_, ok := dbUsers[username]
	return ok
}

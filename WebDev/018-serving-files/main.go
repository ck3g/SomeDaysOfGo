package main

import (
	"io"
	"net/http"
	"os"
)

func main() {
	http.HandleFunc("/", home)
	http.HandleFunc("/iocopy", ioCopy)
	http.ListenAndServe(":8080", nil)
}

func home(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	io.WriteString(w, `
		<h3>Serving files</h3>

		<div>Try different urls</div?
		<ul>
			<li>/iocopy</li>
		</ul>
	`)
}

func ioCopy(w http.ResponseWriter, req *http.Request) {
	f, err := os.Open("WebDev/assets/firefox.jpg")
	if err != nil {
		http.Error(w, "file not found", 404)
		return
	}
	defer f.Close()

	io.Copy(w, f)
}

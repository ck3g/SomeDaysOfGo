package main

import (
	"io"
	"net/http"
	"os"
)

func main() {
	http.HandleFunc("/", home)
	http.HandleFunc("/iocopy", ioCopy)
	http.HandleFunc("/servecontent", serveContent)
	http.HandleFunc("/servefile", serveFile)
	http.Handle("/resources/", http.StripPrefix("/resources", http.FileServer(http.Dir("./WebDev/assets/"))))
	http.ListenAndServe(":8080", nil)
}

func home(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	io.WriteString(w, `
		<h3>Serving files</h3>

		<div>Try different urls</div?
		<ul>
			<li>/iocopy</li>
			<li>/servecontent</li>
			<li>/servefile</li>
			<li>/resources</li>
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

func serveContent(w http.ResponseWriter, req *http.Request) {
	f, err := os.Open("WebDev/assets/firefox.jpg")
	if err != nil {
		http.Error(w, "file not found", 404)
		return
	}
	defer f.Close()

	fi, err := f.Stat()
	if err != nil {
		http.Error(w, "file not found", 404)
		return
	}

	http.ServeContent(w, req, f.Name(), fi.ModTime(), f)
}

func serveFile(w http.ResponseWriter, req *http.Request) {
	http.ServeFile(w, req, "WebDev/assets/firefox.jpg")
}

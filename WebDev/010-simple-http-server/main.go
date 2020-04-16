package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"strings"
)

type requestInfo struct {
	method string
	uri    string
}

var routes map[string]func() string

func main() {
	routes = map[string]func() string{
		"GET /":      homePage,
		"GET /about": aboutPage,
	}

	li, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Fatalln(err.Error())
	}
	defer li.Close()

	for {
		conn, err := li.Accept()
		if err != nil {
			log.Println(err.Error())
			continue
		}
		go handle(conn)
	}
}

func handle(conn net.Conn) {
	defer conn.Close()

	// read request
	request := handleRequest(conn)

	// write response
	handleResponse(conn, request)
}

func handleRequest(conn net.Conn) requestInfo {
	var ri requestInfo
	i := 0
	scanner := bufio.NewScanner(conn)
	for scanner.Scan() {
		ln := scanner.Text()
		fmt.Println(ln)
		if i == 0 {
			// request line
			method := strings.Fields(ln)[0]
			uri := strings.Fields(ln)[1]
			fmt.Println("***METHOD", method)
			fmt.Println("***URI", uri)
			ri.method = method
			ri.uri = uri
		}
		if ln == "" {
			// headers are done
			break
		}
		i++
	}

	return ri
}

func handleResponse(conn net.Conn, request requestInfo) {
	htmlBody := buildHTMLBody(request)
	responseBody := `<!DOCTYPE html><html lang="en"><head><meta charset="UTF-8"><title></title></head><body>` + htmlBody + `</body></html>`

	fmt.Fprint(conn, "HTTP/1.1 200 OK\r\n")
	fmt.Fprintf(conn, "Content-Length: %d\r\n", len(responseBody))
	fmt.Fprint(conn, "Content-Type: text/html\r\n")
	fmt.Fprint(conn, "\r\n")
	fmt.Fprint(conn, responseBody)
}

func buildHTMLBody(request requestInfo) string {
	var pageContent string

	route := request.method + " " + request.uri
	page, ok := routes[route]
	if ok {
		pageContent = page()
	} else {
		pageContent = notFoundPage()
	}

	body := `
<div><strong>Request info</strong></div>
<div>METHOD: ` + request.method + `</div>
<div>URI: ` + request.uri + `</div>
` + pageContent

	return body
}

func homePage() string {
	return `<div>This is home page</div>`
}

func aboutPage() string {
	return `<div>This is about page</div>`
}

func notFoundPage() string {
	return `<div>404: Page Not Found</div>`
}

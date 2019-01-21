package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"strings"
)

func main() {
	li, err := net.Listen("tcp", ":8081")
	if err != nil {
		log.Fatalln(err.Error())
	}
	defer li.Close()

	for {
		conn, err := li.Accept()
		if err != nil {
			log.Fatalln(err.Error())
		}

		go handle(conn)
	}
}

func handle(conn net.Conn) {
	defer conn.Close()

	// read request
	request(conn)

}

func request(conn net.Conn) {
	i := 0
	scanner := bufio.NewScanner(conn)

	for scanner.Scan() {
		ln := scanner.Text()
		fmt.Println(ln)

		if i == 0 {
			// Route the request
			mux(conn, ln)
		}
		// Based on request data we can route
		if ln == "" {
			// As per RC17 spec
			// Blank line means that we are done with request headers.
			// So we ... ->
			break
		}
		i++
	}
}

func mux(conn net.Conn, ln string) {

	method := strings.Fields(ln)[0]
	url := strings.Fields(ln)[1]

	if method == "GET" && url == "/" {
		index(conn)
	}

	if method == "GET" && url == "/about" {
		about(conn)
	}
}

func index(conn net.Conn) {

	body := `<!DOCTYPE html><html lang="en"><head></head><body><strong>Go Server Index Page</strong></body></html>`

	fmt.Fprint(conn, "HTTP/1.1 200 OK\r\n")
	fmt.Fprintf(conn, "Content-Length: %d\r\n", len(body))
	fmt.Fprint(conn, "Content-Type: text/html\r\n")
	fmt.Fprint(conn, "\r\n")
	fmt.Fprint(conn, body)

}

func about(conn net.Conn) {

	body := `<!DOCTYPE html><html lang="en"><head></head><body><strong>Go Server About Page</strong></body></html>`

	fmt.Fprint(conn, "HTTP/1.1 200 OK\r\n")
	fmt.Fprintf(conn, "Content-Length: %d\r\n", len(body))
	fmt.Fprint(conn, "Content-Type: text/html\r\n")
	fmt.Fprint(conn, "\r\n")
	fmt.Fprint(conn, body)

}

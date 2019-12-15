package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"strings"
)

func main() {
	sev, err := net.Listen("tcp", ":9090")
	if err != nil {
		log.Fatalln(err.Error())
	}
	defer sev.Close()

	for {
		conn, err := sev.Accept()
		if err != nil {
			log.Fatalln(err.Error())
			continue
		}
		go handle(conn)
	}
}

func handle(conn net.Conn) {
	defer conn.Close()
	// read request
	request(conn)
	// respond to the request
	respond(conn)
}

func request(conn net.Conn) {
	i := 0
	scanner := bufio.NewScanner(conn)
	for scanner.Scan() {
		ln := scanner.Text()
		fmt.Println(ln)
		if i == 0 {
			// process fields
			m := strings.Fields(ln)
			fmt.Printf("======Method: %s\n", m)
		}
		if ln == "" {
			// done with the headers
			break
		}
		i++
	}
}

func respond(conn net.Conn) {
	body := `<!DOCTYPE html><html lang="en"><head><meta charset="utf-8"><title></title></head> <body><strong> Hello World</strong> </body></html>`

	fmt.Fprintf(conn, "HTTP/1.1 200 OK \r\n")
	fmt.Fprintf(conn, "Content-Lenght: %d", len(body))
	fmt.Fprintf(conn, "Content-Type: text/html\r\n")
	fmt.Fprintf(conn, "\r\n")
	fmt.Fprintf(conn, body)
}

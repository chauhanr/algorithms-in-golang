package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"net"
)

func main() {
	server, err := net.Listen("tcp", ":9090")
	if err != nil {
		log.Fatalln("Error creating TCP server")
	}
	defer server.Close()
	for {
		conn, err := server.Accept()
		if err != nil {
			log.Println("Error obtaining connection to tcp server: ", err)
			continue
		}
		scanner := bufio.NewScanner(conn)

		for scanner.Scan() {
			ln := scanner.Text()
			fmt.Println(ln)
			if ln == "" {
				fmt.Println("=========Break the Scan==================")
				break
			}
		}

		io.WriteString(conn, "Writing the the request connection")
		conn.Close()
	}
}

package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"time"
)

func main() {
	tcp, err := net.Listen("tcp", ":9090")
	if err != nil {
		log.Panic(err)
	}
	defer tcp.Close()

	for {
		conn, err := tcp.Accept()
		if err != nil {
			log.Println(err)
		}
		go handle(conn)
	}
}

func handle(conn net.Conn) {
	// Inorder to set a timeout on the connection we add a timer to the conn.
	err := conn.SetDeadline(time.Now().Add(10 * time.Second))
	if err != nil {
		log.Println("Connection Timed Out")
	}
	scanner := bufio.NewScanner(conn)
	for scanner.Scan() {
		ln := scanner.Text()
		fmt.Println(ln)
		fmt.Fprintf(conn, "I hear you say %s\n", ln)
	}
	defer conn.Close()
}

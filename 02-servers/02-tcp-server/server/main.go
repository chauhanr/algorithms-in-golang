package main

import (
	"fmt"
	"io"
	"log"
	"net"
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

		io.WriteString(conn, "\n Hello from the tcp server\n")
		fmt.Fprintf(conn, "How is your day\n")
		fmt.Fprintf(conn, "%v", "Well, I hope.")
		conn.Close()
	}

}

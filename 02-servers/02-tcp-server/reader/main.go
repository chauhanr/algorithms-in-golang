package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"net"
	"strings"
)

func BufioScanner() {
	s := "I feel so good like anything is possible\n I hit cruise control and rubbed my eyes \n The last 3 days \n The rain was unstoppable."

	scanner := bufio.NewScanner(strings.NewReader(s))

	for scanner.Scan() {
		line := scanner.Text()
		fmt.Println(line)
	}
}

func main() {
	sev, err := net.Listen("tcp", ":9090")
	if err != nil {
		panic(err)
	}

	for {
		conn, err := sev.Accept()
		if err != nil {
			log.Println(err)
		}

		io.WriteString(conn, "\nHello from the TCP server\n")
		fmt.Fprintln(conn, "How is your day")
		conn.Close()
	}

}

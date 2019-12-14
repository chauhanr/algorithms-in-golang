package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"net"
	"os"
)

var rw string
var msg string

func init() {
	flag.StringVar(&rw, "rw", "r", "Read or write flag to a connection")
	flag.StringVar(&msg, "msg", "hello world", "Write a message to the connection")
}

func main() {
	flag.Parse()

	if len(os.Args) == 0 {
		fmt.Printf("Usage: -rw <read>/<write>")
	}

	if rw == "r" {
		read()
	} else {
		write(msg)
	}
}

func write(message string) {
	conn, err := net.Dial("tcp", "localhost:9090")
	if err != nil {
		panic(err)
	}
	defer conn.Close()
	fmt.Fprintln(conn, msg)
}

func read() {
	conn, err := net.Dial("tcp", "localhost:9090")
	if err != nil {
		panic(err)
	}
	defer conn.Close()

	b, err := ioutil.ReadAll(conn)
	if err != nil {
		log.Println(err)
	}
	fmt.Println(string(b))

}

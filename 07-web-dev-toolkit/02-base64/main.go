package main

import (
	"encoding/base64"
	"fmt"
	"log"
)

func main() {
	info := "This is a string the we would want to encode +-<>,."
	s64 := base64.StdEncoding.EncodeToString([]byte(info))

	fmt.Printf("Encoded information: %s\n", s64)

	bs, err := base64.StdEncoding.DecodeString(s64)
	if err != nil {
		log.Fatalln("Error: ", err)
	}
	fmt.Printf("Decoded: %s\n", string(bs))
}

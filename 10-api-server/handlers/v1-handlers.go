package handlers

import (
	"log"
	"net/http"
)

func StatusV1(w http.ResponseWriter, r *http.Request) {
	_, err := w.Write([]byte("/api/v1 is Up"))
	if err != nil {
		log.Printf("Error writing status message %s\n", err)
	}
}

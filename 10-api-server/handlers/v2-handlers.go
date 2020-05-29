package handlers

import (
	"log"
	"net/http"
)

func StatusV2(w http.ResponseWriter, r *http.Request) {
	_, err := w.Write([]byte("/api/v2 status is: up"))
	if err != nil {
		log.Printf("Error writing status for /api/v2 %s\n", err)
	}

}

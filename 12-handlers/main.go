package main

import (
	"log"
	"net/http"

	"github.com/gorilla/context"
	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()

	r.Handle("/api/a", http.HandlerFunc(RoutePatternAHandler))
	r.Handle("/api/b", http.HandlerFunc(RoutePatternBHandler))

	log.Printf("Starting Server Handlers Port: 8081\n")
	err := http.ListenAndServe(":8081", context.ClearHandler(r))
	if err != nil {
		log.Fatal(err)
	}

}

package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

func main() {
	http.HandleFunc("/", dog)
	http.HandleFunc("/dew", dogPicServe)
	http.ListenAndServe(":9090", nil)
}

func dog(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	io.WriteString(w, `<img src="dew.jpg">`)
}

func dogPicServe(w http.ResponseWriter, r *http.Request) {
	f, err := os.Open("./dew.jpg")
	if err != nil {
		fmt.Printf("Error reading file: %s\n", err)
		http.Error(w, "file not found", 404)
		return
	}

	defer f.Close()
	io.Copy(w, f)

}

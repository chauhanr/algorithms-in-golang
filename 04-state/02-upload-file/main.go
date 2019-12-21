package main

import (
	"io"
	"io/ioutil"
	"log"
	"net/http"
)

func main() {

	http.Handle("/", http.FileServer(http.Dir(".")))
	http.Handle("/favicon.ico", http.NotFoundHandler())
	http.HandleFunc("/dew", dew)

	log.Fatalln(http.ListenAndServe(":9090", nil))
}

func dew(w http.ResponseWriter, r *http.Request) {
	var s string
	if r.Method == http.MethodPost {
		f, _, err := r.FormFile("q")
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		defer f.Close()
		//fmt.Printf("\nfile: %s \nheader: %s, \n err: %s", f, h, err)
		br, err := ioutil.ReadAll(f)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		s = string(br)
	}
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	io.WriteString(w, `
         <form method="POST" enctype="multipart/form-data">
	    <input type="file" name="q">
	        <input type="submit"/> 
	    </input> 
	 </form> 
         <br>`+s)
}

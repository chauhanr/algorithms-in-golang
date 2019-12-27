package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"time"
)

func main() {

	http.HandleFunc("/bar", bar)
	http.HandleFunc("/foo", foo)
	http.Handle("/favicon.ico", http.NotFoundHandler())

	http.ListenAndServe(":9090", nil)

}

func foo(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	ctx = context.WithValue(ctx, "userID", 777)
	results, err := dbAccess(ctx)

	if err != nil {
		http.Error(w, err.Error(), http.StatusRequestTimeout)
		return
	}
	fmt.Fprintln(w, results)

}

func dbAccess(ctx context.Context) (int, error) {
	ctx, cancel := context.WithTimeout(ctx, 1*time.Second)
	defer cancel()

	ch := make(chan int)

	go func() {
		uid := ctx.Value("userID").(int)
		time.Sleep(10 * time.Second)

		if ctx.Err() != nil {
			return
		}
		ch <- uid
	}()
	select {
	case <-ctx.Done():
		return 0, ctx.Err()
	case i := <-ch:
		return i, nil
	}

}

func bar(w http.ResponseWriter, r *http.Request) {

	c := r.Context()
	log.Println(c)
	fmt.Fprintln(w, c)

}

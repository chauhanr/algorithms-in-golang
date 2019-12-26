package main

import (
	"database/sql"
	"fmt"
	"io"
	"log"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
)

var db *sql.DB
var err error

func main() {
	db, err = sql.Open("mysql", "chauhr:chauhr@tcp(localhost:3306)/local?charset=utf8")

	check(err)

	defer db.Close()
	err = db.Ping()
	check(err)

	http.HandleFunc("/", index)
	http.HandleFunc("/create", create)
	http.HandleFunc("/insert", insert)
	http.HandleFunc("/read", read)
	http.HandleFunc("/drop", drop)
	http.Handle("/favicon.ico", http.NotFoundHandler())
	err = http.ListenAndServe(":9090", nil)
	check(err)
}

func read(w http.ResponseWriter, r *http.Request) {
	q, err := db.Query(`SELECT * FROM customer;`)
	check(err)
	defer q.Close()

	var name string
	for q.Next() {
		err := q.Scan(&name)
		check(err)
		fmt.Fprintf(w, "RECEIVED Record: %s \n", name)
	}

}

func insert(w http.ResponseWriter, r *http.Request) {
	stmt, err := db.Prepare(`INSERT INTO customer VALUES("James");`)
	check(err)
	defer stmt.Close()

	row, err := stmt.Exec()
	check(err)
	n, err := row.RowsAffected()
	fmt.Fprintln(w, "inserted record in table", n)
}

func create(w http.ResponseWriter, r *http.Request) {
	stmt, err := db.Prepare(`CREATE TABLE customer (name VARCHAR(20));`)
	check(err)
	defer stmt.Close()

	row, err := stmt.Exec()
	check(err)
	n, err := row.RowsAffected()
	check(err)
	fmt.Fprintln(w, "CREATED TABLE cusotomer", n)

}

func drop(w http.ResponseWriter, r *http.Request) {
	d, err := db.Prepare(`DROP TABLE customer;`)
	check(err)
	defer d.Close()

	_, err = d.Exec()
	check(err)
	fmt.Fprintf(w, "Dropped the table customer\n")
}

func index(w http.ResponseWriter, r *http.Request) {
	_, err := io.WriteString(w, "Hello Db App")
	check(err)
}

func check(err error) {
	if err != nil {
		log.Printf("Error: %s", err.Error())
	}
}

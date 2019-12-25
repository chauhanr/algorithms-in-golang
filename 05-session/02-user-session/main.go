package main

import (
	"fmt"
	"io"
	"net/http"

	uuid "github.com/satori/go.uuid"
)

type User struct {
	First    string
	Last     string
	Username string
	Password string
}

var dbUsers = map[string]User{}
var dbSessions = map[string]string{}

func main() {
	// localhost:9090/foo?first=ritesh&last=chauhan
	http.HandleFunc("/foo", foo)
	http.HandleFunc("/bar", bar)
	//localhost:9090/signup?first=ritesh&last=chauhan&uname=chauhr&pass=pass
	http.HandleFunc("/signup", signup)
	http.ListenAndServe(":9090", nil)
}

func foo(w http.ResponseWriter, r *http.Request) {

	c, err := r.Cookie("session")
	if err != nil {
		id, err := uuid.NewV4()
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		c = &http.Cookie{
			Name:  "session",
			Value: id.String(),
		}
		http.SetCookie(w, c)
	}

	var u User

	if un, ok := dbSessions[c.Value]; ok {
		u = dbUsers[un]
	}

	uname := r.FormValue("uname")
	f := r.FormValue("first")
	l := r.FormValue("last")
	u = User{
		First:    f,
		Last:     l,
		Username: uname,
	}
	dbSessions[c.Value] = uname
	dbUsers[uname] = u
}

func signup(w http.ResponseWriter, r *http.Request) {
	if alreadyLoggedIn(r) {
		io.WriteString(w, "You are already logged in.\n")
		return
	}
	p := r.FormValue("pass")
	un := r.FormValue("username")
	f := r.FormValue("first")
	l := r.FormValue("last")

	if _, ok := dbUsers[un]; ok {
		http.Error(w, "username already taken", http.StatusForbidden)
		return
	}
	id, err := uuid.NewV4()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	c := &http.Cookie{
		Name:  "session",
		Value: id.String(),
	}
	http.SetCookie(w, c)
	dbSessions[c.Value] = un

	u := User{
		First:    f,
		Last:     l,
		Username: un,
		Password: p,
	}

	dbUsers[un] = u
	http.Redirect(w, r, "/bar", http.StatusSeeOther)

}

func bar(w http.ResponseWriter, r *http.Request) {
	c, err := r.Cookie("session")
	if err != nil {
		http.Redirect(w, r, "/foo", http.StatusSeeOther)
		return
	}
	un, ok := dbSessions[c.Value]
	if !ok {
		http.Redirect(w, r, "/foo", http.StatusSeeOther)
		return
	}
	u := dbUsers[un]
	fmt.Printf("User Details: %v\n", u)
	user := fmt.Sprintf("{UserName: %s, FirstName: %s, LastName: %s}\n", u.Username, u.First, u.Last)
	io.WriteString(w, user)
}

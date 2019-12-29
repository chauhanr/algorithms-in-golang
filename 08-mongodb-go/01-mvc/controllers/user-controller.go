package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/chauhanr/golang-web/08-mongodb-go/01-mvc/models"
	"github.com/julienschmidt/httprouter"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

/*
   The User controller has been set up as a struct so that later on we can add connection sessions for mongo or mysql databse.
   This allows for testing the api by mocking database sessions.
*/

type UserController struct {
	session *mgo.Session
}

func NewUserController(s *mgo.Session) *UserController {
	return &UserController{session: s}
}

func (u UserController) CreateUser(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	user := models.User{}
	json.NewDecoder(r.Body).Decode(&user)
	user.Id = bson.NewObjectId()
	u.session.DB("go-webdev-db").C("users").Insert(user)

	uj, _ := json.Marshal(user)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated) // 202 is being set
	fmt.Fprintf(w, "%s", uj)
}

func (u UserController) GetUser(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	id := p.ByName("id")
	// verify if the id is indeed at hex
	if !bson.IsObjectIdHex(id) {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	oid := bson.ObjectIdHex(id)
	us := models.User{}
	if err := u.session.DB("go-webdev-db").C("users").FindId(oid).One(&us); err != nil {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	uj, _ := json.Marshal(us)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "%s\n", uj)

}

func (u UserController) DeleteUser(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	id := p.ByName("id")
	if !bson.IsObjectIdHex(id) {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	oid := bson.ObjectIdHex(id)
	if err := u.session.DB("go-webdev-db").C("users").RemoveId(oid); err != nil {
		w.WriteHeader(404)
		return
	}
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "Deleted user: %s\n", oid)
	return

}

func (u UserController) DeleteAllUsers(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {

}

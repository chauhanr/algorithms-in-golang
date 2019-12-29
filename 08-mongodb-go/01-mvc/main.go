package main

import (
	"log"
	"net/http"

	"gopkg.in/mgo.v2"

	"github.com/chauhanr/golang-web/08-mongodb-go/01-mvc/controllers"
	"github.com/julienschmidt/httprouter"
)

func main() {

	r := httprouter.New()
	s := getMongoSession()
	uc := controllers.NewUserController(s)

	r.GET("/user/:id", uc.GetUser)
	r.POST("/user", uc.CreateUser)
	r.DELETE("/user/:id", uc.DeleteUser)

	http.ListenAndServe(":9090", r)

}

func getMongoSession() *mgo.Session {
	s, err := mgo.Dial("mongodb://localhost")
	if err != nil {
		log.Fatalln(err)
		panic(err)
	}
	return s
}

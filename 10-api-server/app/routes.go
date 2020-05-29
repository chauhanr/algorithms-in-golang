package app

import (
	"github.com/chauhanr/golang-web/10-api-server/handlers"
	"github.com/gorilla/mux"
	"net/http"
)

func (s *Server) ApiRoutes() {
	var api = s.Router.PathPrefix("/api").Subrouter()
	api.NotFoundHandler = http.HandlerFunc(handlers.NotFoundHandler)
	s.apiV1Routes(api)
	s.apiV2Routes(api)
}

func (s *Server) apiV1Routes(subRouter *mux.Router) {
	var api = subRouter.PathPrefix("/v1").Subrouter()
	api.NotFoundHandler = http.HandlerFunc(handlers.NotFoundHandler)
	api.HandleFunc("/status", handlers.StatusV1)
}

func (s *Server) apiV2Routes(subRouter *mux.Router) {
	var api = subRouter.PathPrefix("/v2").Subrouter()
	api.NotFoundHandler = http.HandlerFunc(handlers.NotFoundHandler)
	api.HandleFunc("/status", handlers.StatusV2)
}

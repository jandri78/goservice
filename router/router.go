package router

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/jandri78/goservice/controller"
)

func NewRouter() http.Handler {
	r := mux.NewRouter()
	r.HandleFunc("/users", controller.GetUsersHandler).Methods("GET")
	return r
}

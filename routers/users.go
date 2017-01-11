package routers

import (
	"github.com/colombia9503/RESTful-Mysql/controllers"
	"github.com/gorilla/mux"
)

func SetUsersRouters(router *mux.Router) *mux.Router {
	router.HandleFunc("/users", controllers.Users.Create).Methods("POST")
	router.HandleFunc("/users", controllers.Users.Get).Methods("GET")
	router.HandleFunc("/users/{id}", controllers.Users.GetOne).Methods("GET")
	router.HandleFunc("/users/{id}", controllers.Users.Update).Methods("PUT")
	router.HandleFunc("/users/{id}", controllers.Users.Delete).Methods("DELETE")
	return router
}

package routers

import (
	"github.com/colombia9503/RESTful-Mysql/controllers"
	"github.com/gorilla/mux"
)

func SetUsuariosRouters(router *mux.Router) *mux.Router {
	router.HandleFunc("/usuarios", controllers.Usuarios.Create).Methods("POST")
	router.HandleFunc("/usuarios", controllers.Usuarios.Get).Methods("GET")
	router.HandleFunc("/usuarios/{id}", controllers.Usuarios.GetOne).Methods("GET")
	router.HandleFunc("/usuarios/{id}", controllers.Usuarios.Update).Methods("PUT")
	router.HandleFunc("/usuarios/{id}", controllers.Usuarios.Delete).Methods("DELETE")
	return router
}

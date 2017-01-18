package routers

import (
	"github.com/colombia9503/RESTful-Mysql/controllers"
	"github.com/gorilla/mux"
)

func SetTransformadoresRouters(router *mux.Router) *mux.Router {
	router.HandleFunc("/api/transformadores", controllers.Transformadores.Create).Methods("POST")
	router.HandleFunc("/api/transformadores", controllers.Transformadores.Get).Methods("GET")
	router.HandleFunc("/api/transformadores/{id}", controllers.Transformadores.GetOne).Methods("GET")
	router.HandleFunc("/api/transformadores/{id}", controllers.Transformadores.Update).Methods("PUT")
	router.HandleFunc("/api/transformadores/{id}", controllers.Transformadores.Delete).Methods("DELETE")
	return router
}

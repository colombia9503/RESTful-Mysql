package routers

import (
	"github.com/colombia9503/RESTful-Mysql/controllers"
	"github.com/gorilla/mux"
)

func SetConcentracionesRouters(router *mux.Router) *mux.Router {
	router.HandleFunc("/api/concentracion", controllers.Concentraciones.Create).Methods("POST")
	router.HandleFunc("/api/concentracion", controllers.Concentraciones.Get).Methods("GET")
	router.HandleFunc("/api/concentracion/{id}", controllers.Concentraciones.GetOne).Methods("GET")
	router.HandleFunc("/api/concentracion/{id}", controllers.Concentraciones.Update).Methods("PUT")
	router.HandleFunc("/api/concentracion/{id}", controllers.Concentraciones.Delete).Methods("DELETE")
	return router
}

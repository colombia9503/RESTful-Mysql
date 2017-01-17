package routers

import (
	"github.com/colombia9503/RESTful-Mysql/controllers"
	"github.com/gorilla/mux"
)

func SetClientesRouters(router *mux.Router) *mux.Router {
	router.HandleFunc("/api/clientes", controllers.Clientes.Get).Methods("GET");
	router.HandleFunc("/api/clientes/{id}", controllers.Clientes.GetOne).Methods("GET");	
	return router
}

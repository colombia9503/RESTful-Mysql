package routers

import (
    "github.com/colombia9503/RESTful-Mysql/controllers"
	"github.com/gorilla/mux"
)

func SetMarcasRouters(router *mux.Router) *mux.Router {
    router.HandleFunc("/api/marcas",controllers.Marcas.Create).Methods("POST");
    router.HandleFunc("/api/marcas",controllers.Marcas.Get).Methods("GET");
    router.HandleFunc("/api/marcas/{id}",controllers.Marcas.GetOne).Methods("GET");
    router.HandleFunc("/api/marcas/{id}",controllers.Marcas.Update).Methods("PUT");
    router.HandleFunc("/api/marcas/{id}",controllers.Marcas.Delete).Methods("DELETE");
    return router
}
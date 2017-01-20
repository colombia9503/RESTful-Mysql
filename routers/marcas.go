package routers

import (
	"github.com/colombia9503/RESTful-Mysql/common"
	"github.com/colombia9503/RESTful-Mysql/controllers"
	"github.com/gorilla/mux"
	"github.com/urfave/negroni"
)

func SetMarcasRouters(router *mux.Router) *mux.Router {
	marcasRouter := mux.NewRouter()
	marcasRouter.HandleFunc("/api/marcas", controllers.Marcas.Create).Methods("POST")
	marcasRouter.HandleFunc("/api/marcas", controllers.Marcas.Get).Methods("GET")
	marcasRouter.HandleFunc("/api/marcas/{id}", controllers.Marcas.GetOne).Methods("GET")
	marcasRouter.HandleFunc("/api/marcas/{id}", controllers.Marcas.Update).Methods("PUT")
	marcasRouter.HandleFunc("/api/marcas/{id}", controllers.Marcas.Delete).Methods("DELETE")
	router.PathPrefix("/api/marcas").Handler(
		negroni.New(
			negroni.HandlerFunc(common.Authorize),
			negroni.Wrap(marcasRouter),
		))
	return router
}

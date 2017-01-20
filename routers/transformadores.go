package routers

import (
	"github.com/colombia9503/RESTful-Mysql/common"
	"github.com/colombia9503/RESTful-Mysql/controllers"
	"github.com/gorilla/mux"
	"github.com/urfave/negroni"
)

func SetTransformadoresRouters(router *mux.Router) *mux.Router {
	transformadoresRouter := mux.NewRouter()
	transformadoresRouter.HandleFunc("/api/transformadores", controllers.Transformadores.Create).Methods("POST")
	transformadoresRouter.HandleFunc("/api/transformadores", controllers.Transformadores.Get).Methods("GET")
	transformadoresRouter.HandleFunc("/api/transformadores/{id}", controllers.Transformadores.GetOne).Methods("GET")
	transformadoresRouter.HandleFunc("/api/transformadores/{id}", controllers.Transformadores.Update).Methods("PUT")
	transformadoresRouter.HandleFunc("/api/transformadores/{id}", controllers.Transformadores.Delete).Methods("DELETE")
	router.PathPrefix("/api/transformadores").Handler(
		negroni.New(
			negroni.HandlerFunc(common.Authorize),
			negroni.Wrap(transformadoresRouter),
		))
	return router
}

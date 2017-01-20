package routers

import (
	"github.com/colombia9503/RESTful-Mysql/common"
	"github.com/colombia9503/RESTful-Mysql/controllers"
	"github.com/gorilla/mux"
	"github.com/urfave/negroni"
)

func SetConcentracionesRouters(router *mux.Router) *mux.Router {
	concentracionesRouter := mux.NewRouter()
	concentracionesRouter.HandleFunc("/api/concentracion", controllers.Concentraciones.Create).Methods("POST")
	concentracionesRouter.HandleFunc("/api/concentracion", controllers.Concentraciones.Get).Methods("GET")
	concentracionesRouter.HandleFunc("/api/concentracion/{id}", controllers.Concentraciones.GetOne).Methods("GET")
	concentracionesRouter.HandleFunc("/api/concentracion/{id}", controllers.Concentraciones.Update).Methods("PUT")
	concentracionesRouter.HandleFunc("/api/concentracion/{id}", controllers.Concentraciones.Delete).Methods("DELETE")
	router.PathPrefix("/api/concentracion").Handler(
		negroni.New(
			negroni.HandlerFunc(common.Authorize),
			negroni.Wrap(concentracionesRouter),
		))
	return router
}

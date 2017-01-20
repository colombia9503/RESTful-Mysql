package routers

import (
	"github.com/colombia9503/RESTful-Mysql/common"
	"github.com/colombia9503/RESTful-Mysql/controllers"
	"github.com/gorilla/mux"
	"github.com/urfave/negroni"
)

func SetClientesRouters(router *mux.Router) *mux.Router {
	clientesRouter := mux.NewRouter()
	clientesRouter.HandleFunc("/api/clientes", controllers.Clientes.Get).Methods("GET")
	clientesRouter.HandleFunc("/api/clientes/{id}", controllers.Clientes.GetOne).Methods("GET")
	router.PathPrefix("/api/clientes").Handler(
		negroni.New(
			negroni.HandlerFunc(common.Authorize),
			negroni.Wrap(clientesRouter),
		))
	return router
}

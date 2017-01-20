package routers

import (
	"github.com/colombia9503/RESTful-Mysql/common"
	"github.com/colombia9503/RESTful-Mysql/controllers"
	"github.com/gorilla/mux"
	"github.com/urfave/negroni"
)

func SetUsuariosRouters(router *mux.Router) *mux.Router {
	usuariosRouter := mux.NewRouter()
	usuariosRouter.HandleFunc("/api/usuarios", controllers.Usuarios.Create).Methods("POST")
	usuariosRouter.HandleFunc("/api/usuarios", controllers.Usuarios.Get).Methods("GET")
	usuariosRouter.HandleFunc("/api/usuarios/{id}", controllers.Usuarios.GetOne).Methods("GET")
	usuariosRouter.HandleFunc("/api/usuarios/{id}", controllers.Usuarios.Update).Methods("PUT")
	usuariosRouter.HandleFunc("/api/usuarios/{id}", controllers.Usuarios.Delete).Methods("DELETE")
	router.PathPrefix("/api/usuarios").Handler(
		negroni.New(
			negroni.HandlerFunc(common.Authorize),
			negroni.Wrap(usuariosRouter),
		))
	return router
}

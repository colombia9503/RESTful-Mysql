package routers

import "github.com/gorilla/mux"

func InitRouters() *mux.Router {
	router := mux.NewRouter().StrictSlash(false)
	router = SetUsuariosRouters(router)
	router = SetAuthRouter(router)
	router = SetClientesRouters(router)
	router = SetConcentracionesRouters(router)
	router = SetMarcasRouters(router)
	router = SetTransformadoresRouters(router)
	return router
}

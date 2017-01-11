package routers

import "github.com/gorilla/mux"

func InitRouters() *mux.Router {
	router := mux.NewRouter().StrictSlash(false)
	router = SetUsuariosRouters(router)
	router = SetAuthRouter(router)
	return router
}

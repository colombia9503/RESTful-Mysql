package routers

import (
	"github.com/colombia9503/RESTful-Mysql/controllers"
	"github.com/gorilla/mux"
)

func SetAuthRouter(router *mux.Router) *mux.Router {
	//router.HandleFunc("/auth", controllers.Auth.).Methods("GET")
	router.HandleFunc("/api/auth", controllers.Auth.Login).Methods("POST")
	return router
}

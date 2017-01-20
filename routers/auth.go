package routers

import (
	"github.com/colombia9503/RESTful-Mysql/controllers"
	"github.com/gorilla/mux"
)

//requeres HEADER: Autorization: Beared + TOKEN
func SetAuthRouter(router *mux.Router) *mux.Router {
	router.HandleFunc("/auth/login", controllers.Auth.Login).Methods("POST")
	//router.HandleFunc("/auth/token-auth", controllers.Auth.Login).Methods("POST")
	/**router.Handle("/auth/refresh-token-auth",
		negroni.New(
			negroni.HandlerFunc(authentication.RequireTokenAuthentication),
			negroni.HandlerFunc(controllers.Auth.RefreshToken),
		)).Methods("GET")

	router.Handle("/auth/logout",
		negroni.New(
			negroni.HandlerFunc(authentication.RequireTokenAuthentication),
			negroni.HandlerFunc(controllers.Auth.Logout),
		)).Methods("GET")*/
	return router
}

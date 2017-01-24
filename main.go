package main

import (
	"log"
	"net/http"

	"github.com/colombia9503/RESTful-Mysql/common"
	"github.com/colombia9503/RESTful-Mysql/routers"
	"github.com/urfave/negroni"
)

func main() {
	common.StartUp()
	r := routers.InitRouters()

	n := negroni.Classic()
	n.UseHandler(r)

	log.Println("Listening.. 8082")
	http.ListenAndServe(":8082", n)
}

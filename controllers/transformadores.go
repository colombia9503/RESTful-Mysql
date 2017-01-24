package controllers

import (
	"net/http"

	"github.com/colombia9503/RESTful-Mysql/common"
	"github.com/colombia9503/RESTful-Mysql/models"
	"github.com/gorilla/mux"
)

type transformadoresController struct{}

var Transformadores = new(transformadoresController)

func (mc *transformadoresController) Create(w http.ResponseWriter, r *http.Request) {

}

func (mc *transformadoresController) Get(w http.ResponseWriter, r *http.Request) {

}

func (mc *transformadoresController) GetOne(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]
	println("GetOne: ", id)
}

func (mc *transformadoresController) Update(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]
	println("Put: ", id)
}

func (mc *transformadoresController) Delete(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]
	println("Delete: ", id)
}

func (mc *transformadoresController) Upload(w http.ResponseWriter, r *http.Request) {
	err := models.Transformadores.UploadImage(r)
	if err != nil {
		common.JsonError(w, err, http.StatusBadRequest)
		return
	}
	common.JsonStatus(w, http.StatusOK)
}

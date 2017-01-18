package controllers

import (
	"encoding/json"
	"net/http"

	"github.com/colombia9503/RESTful-MDB/common"
	"github.com/colombia9503/RESTful-Mysql/models"
	"github.com/gorilla/mux"
)

type marcasController struct{}

var Marcas = new(marcasController)

func (mc *marcasController) Create(w http.ResponseWriter, r *http.Request) {
	var m models.Marca
	if err := json.NewDecoder(r.Body).Decode(&m); err != nil {
		common.JsonError(w, err, http.StatusBadRequest)
		return
	}

	err := models.Marcas.Insert(m.Marca)
	if err != nil {
		common.JsonError(w, err, http.StatusBadRequest)
		return
	}

	common.JsonStatus(w, http.StatusOK)
}

func (mc *marcasController) Get(w http.ResponseWriter, r *http.Request) {

}

func (mc *marcasController) GetOne(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]
	println("GetOne: ", id)
}

func (mc *marcasController) Update(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]
	println("Put: ", id)
}

func (mc *marcasController) Delete(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]
	println("Delete: ", id)
}

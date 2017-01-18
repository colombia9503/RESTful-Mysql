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
	results, err := models.Marcas.SelectAll()
	if err != nil {
		common.JsonError(w, err, http.StatusBadRequest)
		return
	}

	res, err := json.Marshal(results)
	if err != nil {
		common.JsonError(w, err, http.StatusBadRequest)
		return
	}

	common.JsonOk(w, res, http.StatusOK)
}

func (mc *marcasController) GetOne(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]
	println("GetOne: ", id)
	result, err := models.Marcas.SelectOne(id)
	if err != nil {
		common.JsonError(w, err, http.StatusBadRequest)
		return
	}

	res, err := json.Marshal(result)
	if err != nil {
		common.JsonError(w, err, http.StatusBadRequest)
		return
	}

	common.JsonOk(w, res, http.StatusOK)
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

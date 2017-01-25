package controllers

import (
	"encoding/json"
	"net/http"

	"github.com/colombia9503/RESTful-Mysql/common"
	"github.com/colombia9503/RESTful-Mysql/models"
	"github.com/gorilla/mux"
)

type transformadoresController struct{}

var Transformadores = new(transformadoresController)

func (mc *transformadoresController) Create(w http.ResponseWriter, r *http.Request) {
	var t models.Transformador
	if err := json.NewDecoder(r.Body).Decode(&t); err != nil {
		common.JsonError(w, err, http.StatusBadRequest)
		return
	}

	err := models.Transformadores.Insert(t)
	if err != nil {
		common.JsonError(w, err, http.StatusBadRequest)
		return
	}

	common.JsonStatus(w, http.StatusOK)
}

func (mc *transformadoresController) Get(w http.ResponseWriter, r *http.Request) {
	results, err := models.Transformadores.SelectAll()
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

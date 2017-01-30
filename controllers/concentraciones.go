package controllers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/colombia9503/RESTful-Mysql/common"
	"github.com/colombia9503/RESTful-Mysql/models"
	"github.com/gorilla/mux"
)

type concentracionesController struct{}

var Concentraciones = new(concentracionesController)

func (coc *concentracionesController) Create(w http.ResponseWriter, r *http.Request) {
	var m models.Concentracion
	if err := json.NewDecoder(r.Body).Decode(&m); err != nil {
		common.JsonError(w, err, http.StatusBadRequest)
		return
	}
	err := models.Concentraciones.Insert(m)
	if err != nil {
		common.JsonError(w, err, http.StatusBadRequest)
		return
	}

	common.JsonStatus(w, http.StatusOK)
}

func (coc *concentracionesController) Get(w http.ResponseWriter, r *http.Request) {
	results, err := models.Concentraciones.SelectAll()
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

func (coc *concentracionesController) GetOne(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]
	result, err := models.Concentraciones.SelectOne(id)
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

func (coc *concentracionesController) Update(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		common.JsonError(w, err, http.StatusBadRequest)
		return
	}

	var c models.Concentracion
	if err := json.NewDecoder(r.Body).Decode(&c); err != nil {
		common.JsonError(w, err, http.StatusBadRequest)
		return
	}

	result, err := models.Concentraciones.Update(id, c)
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

func (coc *concentracionesController) Delete(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		common.JsonError(w, err, http.StatusBadRequest)
		return
	}

	err = models.Concentraciones.Delete(id)
	if err != nil {
		common.JsonError(w, err, http.StatusBadRequest)
		return
	}

	common.JsonStatus(w, http.StatusOK)
}

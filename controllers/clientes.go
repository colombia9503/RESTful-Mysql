package controllers

import (
	"encoding/json"
	"net/http"

	"github.com/colombia9503/RESTful-MDB/common"
	"github.com/colombia9503/RESTful-Mysql/models"
	"github.com/gorilla/mux"
)

type clientesController struct{}

var Clientes = new(clientesController)

func (cc *clientesController) Get(w http.ResponseWriter, r *http.Request) {
	query := r.URL.Query()["ctransf"]
	clts, err := models.Clientes.SelectAll(query[0])
	if err != nil {
		common.JsonError(w, err, http.StatusBadRequest)
	}

	res, err := json.Marshal(clts)
	if err != nil {
		common.JsonError(w, err, http.StatusBadRequest)
		return
	}

	common.JsonOk(w, res, http.StatusOK)
}

func (cc *clientesController) GetOne(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]
	result, err := models.Clientes.SelectOne(id)
	if err != nil {
		common.JsonError(w, err, http.StatusNotFound)
		return
	}

	res, err := json.Marshal(result)
	if err != nil {
		common.JsonError(w, err, http.StatusNotFound)
		return
	}

	common.JsonOk(w, res, http.StatusOK)
}

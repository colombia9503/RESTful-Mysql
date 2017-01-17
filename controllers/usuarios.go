package controllers

import (
	"encoding/json"
	"net/http"

	"github.com/colombia9503/RESTful-Mysql/common"
	"github.com/colombia9503/RESTful-Mysql/models"
	"github.com/gorilla/mux"
)

type usuariosController struct{}

var Usuarios = new(usuariosController)

func (uc *usuariosController) Create(w http.ResponseWriter, r *http.Request) {

}

func (uc *usuariosController) Get(w http.ResponseWriter, r *http.Request) {
	usrs, err := models.Usuarios.SelectAll()
	if err != nil {
		common.JsonError(w, err, http.StatusBadRequest)
		return
	}
	res, err := json.Marshal(usrs)
	if err != nil {
		common.JsonError(w, err, http.StatusBadRequest)
		return
	}
	common.JsonOk(w, res, http.StatusOK)
}

func (uc *usuariosController) GetOne(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]
	println("GetOne: ", id)
}

func (uc *usuariosController) Update(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]
	println("Put: ", id)
}

func (uc *usuariosController) Delete(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]
	println("Delete: ", id)
}

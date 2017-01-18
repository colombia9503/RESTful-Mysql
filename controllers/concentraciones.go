package controllers

import (
	"net/http"

	"github.com/gorilla/mux"
)

type concentracionesController struct{}

var Concentraciones = new(concentracionesController)

func (coc *concentracionesController) Create(w http.ResponseWriter, r *http.Request) {

}

func (coc *concentracionesController) Get(w http.ResponseWriter, r *http.Request) {

}

func (coc *concentracionesController) GetOne(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]
	println("GetOne: ", id)
}

func (coc *concentracionesController) Update(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]
	println("Put: ", id)
}

func (coc *concentracionesController) Delete(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]
	println("Delete: ", id)
}

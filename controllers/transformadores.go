package controllers

import (
	"net/http"

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

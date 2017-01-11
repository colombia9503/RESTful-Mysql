package controllers

import (
	"net/http"

	"github.com/gorilla/mux"
)

type usersController struct{}

var Users = new(usersController)

func (uc *usersController) Create(w http.ResponseWriter, r *http.Request) {

}

func (uc *usersController) Get(w http.ResponseWriter, r *http.Request) {

}

func (uc *usersController) GetOne(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]
	println("GetOne: ", id)
}

func (uc *usersController) Update(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]
	println("Put: ", id)
}

func (uc *usersController) Delete(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]
	println("Delete: ", id)
}

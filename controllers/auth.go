package controllers

import (
	"encoding/json"
	"net/http"

	"github.com/colombia9503/RESTful-MDB/common"
	"github.com/colombia9503/RESTful-Mysql/models"
)

type authController struct{}

var Auth = new(authController)

func (ac *authController) Login(w http.ResponseWriter, r *http.Request) {
	var u models.Usuario
	if err := json.NewDecoder(r.Body).Decode(&u); err != nil {
		common.JsonError(w, err, http.StatusBadRequest)
		return
	}

	result, err := models.Auth.Login(u.Usuario, u.Clave)
	if err != nil {
		common.JsonError(w, err, http.StatusUnauthorized)
		return
	}

	res, err := json.Marshal(result)
	if err != nil {
		common.JsonError(w, err, http.StatusBadRequest)
		return
	}

	common.JsonOk(w, res, http.StatusOK)
}

func (ac *authController) RefreshToken(w http.ResponseWriter, r *http.Request, next http.HandlerFunc) {

}

func (ac *authController) Logout(w http.ResponseWriter, r *http.Request, next http.HandlerFunc) {

}

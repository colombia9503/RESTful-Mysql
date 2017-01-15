package models

import (
	"crypto/sha512"
	"encoding/hex"

	"github.com/colombia9503/RESTful-Mysql/common"
)

var Auth = new(auth)

type auth struct{}

func (auth) LogIn(User, Pwd string) string {
	var u *Usuario
	row := common.DB.QueryRow("select usuario, clave, salt, rol from usuario where usuario ='" + User + "' and activo = 1 and borrado = 0")

	if err := row.Scan(&u.Usuario, &u.Clave, &u.Salt, &u.Rol); err != nil {
		return "no hay columnas"
	}

	if ComparePasswords(u.Clave, Pwd, u.Salt) {
		return "correcto"
	}

	return "credenciales incorrectas"
}

func ComparePasswords(bdpwd, pwd, salt string) bool {
	h512 := sha512.New()
	h512.Write([]byte(pwd))
	pw := hex.EncodeToString(h512.Sum(nil)[:]) + salt
	h512pw := sha512.New()
	h512pw.Write([]byte(pw))
	if hex.EncodeToString(h512pw.Sum(nil)[:]) == bdpwd {
		return true
	}
	return false
}

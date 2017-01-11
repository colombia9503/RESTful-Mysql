package models

import (
	"crypto/sha512"
	"encoding/hex"

	"github.com/colombia9503/RESTful-Mysql/common"
)

var Auth = new(auth)

type auth struct{}

func (auth) LogIn(User, Pwd string) string {
	var u Usuario
	row := common.DB.QueryRow("select usuario, clave, salt, rol from usuario where usuario ='" + User + "' and activo = 1 and borrado = 0")
	if row == nil {
		return "Usuario no existe"
	}

	if err := row.Scan(&u.Usuario, &u.Clave, &u.Salt, &u.Rol); err != nil {
		return "err"
	}

	if ComparePasswords(u.Clave, Pwd, u.Salt) {
		return "logeado"
	}

	return "Contraseña y usuario incorrectos"
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

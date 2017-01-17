package models

import (
	"crypto/sha512"
	"encoding/hex"

	"github.com/colombia9503/RESTful-Mysql/common"
)

type Credencial struct {
	Nombre  string `db:"nombre"`
	Usuario string `db:"usuario"`
	Rol     string `db:"rol"`
}

var Auth = new(auth)

type auth struct{}

func (auth) LogIn(User, Pwd string) (*Credencial, error) {
	var u Usuario
	var c *Credencial
	row := common.DB.QueryRow("select nombre, usuario, clave, salt, rol from usuario where usuario ='" + User + "' and activo = 1 and borrado = 0")

	if err := row.Scan(&u.Nombre, &u.Usuario, &u.Clave, &u.Salt, &u.Rol); err != nil {
		//err = no sql rows
		return nil, common.NewLogErr("Invalid Password/Account")
	}

	if ComparePasswords(u.Clave, Pwd, u.Salt) {
		c = &Credencial{
			Nombre:  u.Nombre,
			Usuario: u.Usuario,
			Rol:     u.Rol,
		}
		return c, nil
	}

	return nil, common.NewLogErr("Invalid Password/Account")
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

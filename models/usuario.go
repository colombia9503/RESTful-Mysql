package models

import (
	"github.com/colombia9503/RESTful-Mysql/common"
)

type Usuario struct {
	Codigo  int    `db:"codigo"`
	Cedula  string `db:"cedula"`
	Nombre  string `db:"nombre"`
	Usuario string `db:"usuario"`
	Clave   string `db:"clave"`
	Salt    string `db:"salt"`
	Rol     string `db:"rol"`
	Activo  int    `db:"activo"`
	Borrado int    `db:"borrado"`
}

var Usuarios = new(usuarios)

type usuarios struct{}

//d1u2v3a4n5
func (usuarios) SelectAll() ([]Usuario, error) {
	us := []Usuario{}
	rows, err := common.DB.Query("select codigo, cedula, nombre, usuario, clave, salt, rol, activo, borrado from usuario;")
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	for rows.Next() {
		var u Usuario
		if err := rows.Scan(&u.Codigo, &u.Cedula, &u.Nombre, &u.Usuario, &u.Clave, &u.Salt, &u.Rol, &u.Activo, &u.Borrado); err != nil {
			return nil, err
		}
		us = append(us, u)
	}
	return us, nil
}

func (usuarios) SelectOne(id string) (*Usuario, error) {
	return nil, nil
}

func (usuarios) Insert(name, last_name string) (*Usuario, error) {
	return nil, nil
}

func (usuarios) Update(id, name, last_name string) error {
	return nil
}

func (usuarios) Delete(id string) error {
	return nil
}

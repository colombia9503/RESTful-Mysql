package models

import (
	"github.com/colombia9503/RESTful-Mysql/common"
)

type Cliente struct {
	Codigo      int    `db:"codigo"`
	Sede        int    `db:"sede"`
	ID          string `db:"id"`
	RazonSocial string `db:"razonsocial"`
	Telefono    string `db:"telefono"`
	Direccion   string `db:"Direccion"`
	Borrado     int    `db:"borrado"`
}

var Clientes = new(clientes)

type clientes struct{}

func (clientes) SelectAll() ([]Cliente, error) {
	cl := []Cliente{}
	rows, err := common.DB.Query("select * from cliente where borrado = 0;")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var c Cliente
		if err := rows.Scan(&c.Codigo, &c.Sede, &c.ID, &c.RazonSocial, &c.Telefono, &c.Direccion, &c.Borrado); err != nil {
			return nil, err
		}
		cl = append(cl, c)
	}
	return cl, nil

}

func (clientes) SelectOne(ID string) (Cliente, error) {
	var c Cliente
	row := common.DB.QueryRow("select * from cliente where id = '" + ID + "' and borrado = 0;")
	return c, row.Scan(&c.Codigo, &c.Sede, &c.ID, &c.RazonSocial, &c.Telefono, &c.Direccion, &c.Borrado)
}

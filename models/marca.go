package models

import (
	"github.com/colombia9503/RESTful-Mysql/common"
)

type Marca struct {
	ID      int    `db:"id"`
	Codigo  string `db:"codigo"`
	Nombre  string `db:"nombre"`
	Marca   string `db:"marca"`
	Borrado string `db:"borrado"`
}

var Marcas = new(marcas)

type marcas struct{}

func (marcas) SelectAll() ([]Marca, error) {
	mrs := []Marca{}
	rows, err := common.DB.Query("select * from marca;")
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	for rows.Next() {
		var m Marca
		if err := rows.Scan(&m.ID, &m.Codigo, &m.Nombre, &m.Marca, &m.Borrado); err != nil {
			return nil, err
		}
		mrs = append(mrs, m)
	}
	return mrs, nil
}

func (marcas) SelectOne(id string) (Marca, error) {
	var m Marca
	row := common.DB.QueryRow("select * from marca where id = " + id + " and borrado = 0")
	return m, row.Scan(&m.ID, &m.Codigo, &m.Nombre, &m.Marca, &m.Borrado)
}

func (marcas) Insert(m Marca) error {
	stmt, err := common.DB.Prepare("insert into marca (marca, nombre, codigo) values(?,?,?);")
	if err != nil {
		return err
	}
	_, err = stmt.Exec(m.Marca, m.Nombre, m.Codigo)
	defer stmt.Close()
	return err
}

func (marcas) Update(id int, mr Marca) (Marca, error) {
	stmt, err := common.DB.Prepare("update marca set marca = ? where id = ? and borrado = 0;")
	if err != nil {
		return mr, err
	}
	_, err = stmt.Exec(mr.Marca, id)
	defer stmt.Close()
	return mr, err
}

func (marcas) Delete(id int) error {
	stmt, err := common.DB.Prepare("update marca set borrado = 1 where id = ?;")
	if err != nil {
		return err
	}
	_, err = stmt.Exec(id)
	defer stmt.Close()
	return err
}

package models

import "github.com/colombia9503/RESTful-Mysql/common"

type Concentracion struct {
	ID            int    `db:"id"`
	Codigo        string `db:"codigo"`
	Concentracion string `db:"concentracion"`
	Borrado       int    `db:"borrado"`
}

var Concentraciones = new(concentraciones)

type concentraciones struct{}

func (concentraciones) SelectAll() ([]Concentracion, error) {
	cts := []Concentracion{}
	rows, err := common.DB.Query("select * from concentracion;")
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	for rows.Next() {
		var c Concentracion
		if err := rows.Scan(&c.ID, &c.Codigo, &c.Concentracion, &c.Borrado); err != nil {
			return nil, err
		}
		cts = append(cts, c)
	}
	return cts, nil
}

func (concentraciones) SelectOne(id string) (Concentracion, error) {
	var c Concentracion
	row := common.DB.QueryRow("select * from concentracion where id = " + id + " and borrado = 0")
	return c, row.Scan(&c.ID, &c.Codigo, &c.Concentracion, &c.Borrado)
}

func (concentraciones) Insert(c Concentracion) error {
	stmt, err := common.DB.Prepare("insert into concentracion (codigo, concentracion) values(?,?);")
	if err != nil {
		return err
	}
	_, err = stmt.Exec(c.Codigo, c.Concentracion)
	defer stmt.Close()
	return err
}

func (concentraciones) Update(id int, c Concentracion) (Concentracion, error) {
	stmt, err := common.DB.Prepare("update concentracion set Concentracion = ? where id = ? and borrado = 0;")
	if err != nil {
		return c, err
	}
	_, err = stmt.Exec(c.Concentracion, id)
	defer stmt.Close()
	return c, err
}

func (concentraciones) Delete(id int) error {
	stmt, err := common.DB.Prepare("update concentracion set borrado = 1 where id = ?;")
	if err != nil {
		return err
	}
	_, err = stmt.Exec(id)
	defer stmt.Close()
	return err
}

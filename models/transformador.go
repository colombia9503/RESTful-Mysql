package models

import (
	"io"
	"net/http"
	"os"

	"github.com/colombia9503/RESTful-Mysql/common"
)

type Transformador struct {
	ID               int     `db:"id"`
	Codigo           string  `db:"codigo"`
	Cliente          string  `db:"cliente"`
	Marca            string  `db:"marca"`
	KV               float32 `db:"kv"`
	Peso             float32 `db:"peso"`
	Concentracion    int     `db:"concentracion"`
	Descripcion      string  `db:"descripcion"`
	ImgTransformador string  `db:"img_transformador"`
	ImgPlaqueta      string  `db:"img_plaqueta"`
	ImgBascula       string  `db:"img_bascula"`
	Borrado          int     `db:"borrado"`
}

const MAX_MEMORY = 15 * 1024 * 1024
const SERVER_PATH = "ServerImages/"

var Transformadores = new(transformadores)

type transformadores struct{}

func (transformadores) SelectAll() ([]Transformador, error) {
	transfs := []Transformador{}
	rows, err := common.DB.Query("Select * from transformador;")
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	for rows.Next() {
		var t Transformador
		if err := rows.Scan(&t.ID, &t.Codigo, &t.Cliente, &t.Marca, &t.KV, &t.Peso,
			&t.Concentracion, &t.Descripcion, &t.ImgTransformador, &t.ImgBascula, &t.ImgPlaqueta,
			&t.Borrado); err != nil {
			return nil, err
		}
		transfs = append(transfs, t)
	}
	return transfs, nil
}

func (transformadores) SelectOne(id string) (Transformador, error) {
	var t Transformador
	row := common.DB.QueryRow("select * from transformador where id=? and borrado = 0", id)
	return t, row.Scan(&t.ID, &t.Codigo, &t.Cliente, &t.Marca, &t.KV, &t.Peso,
		&t.Concentracion, &t.Descripcion, &t.ImgTransformador, &t.ImgBascula, &t.ImgPlaqueta,
		&t.Borrado)
}

func (transformadores) Insert(t Transformador) error {
	stmt, err := common.DB.Prepare("insert into concentracion (codigo, cliente, marca, kv, peso, concentracion, " +
		"descripcion, img_transformador, img_bascula, img_plaqueta) " +
		"values(?,?,?,?,?,?,?,?,?,?);")
	if err != nil {
		return err
	}
	_, err = stmt.Exec(t.Codigo, t.Cliente, t.Marca, t.KV, t.Peso,
		t.Concentracion, t.Descripcion, t.ImgTransformador, t.ImgBascula, t.ImgPlaqueta)
	defer stmt.Close()
	return err
}

func (transformadores) Update(t Transformador, id int) error {
	stmt, err := common.DB.Prepare("update concentracion set codigo = ?, cliente = ?, marca = ?, kv = ?, peso = ?, concentracion = ?, " +
		"descripcion = ?, img_transformador = ?, img_bascula = ?, img_plaqueta = ? " +
		"where id = ? and borrado = 0;")
	if err != nil {
		return err
	}
	_, err = stmt.Exec(t.Codigo, t.Cliente, t.Marca, t.KV, t.Peso,
		t.Concentracion, t.Descripcion, t.ImgTransformador, t.ImgBascula, t.ImgPlaqueta,
		id)
	defer stmt.Close()
	return err
}

func (transformadores) Delete(id int) error {
	stmt, err := common.DB.Prepare("update concentracion set borrado = 1 where id = ? and borrado = 0;")
	if err != nil {
		return err
	}

	_, err = stmt.Exec(id)
	defer stmt.Close()
	return err
}

func (transformadores) UploadImage(r *http.Request) error {
	if err := r.ParseMultipartForm(MAX_MEMORY); err != nil {
		return err
	}

	m := r.MultipartForm

	files := m.File["images"]
	clientid := m.Value["cliente"]
	transid := m.Value["transformador"]
	for i, _ := range files {
		file, err := files[i].Open()
		//defer file.Close()
		if err != nil {
			return err
		}
		os.MkdirAll(SERVER_PATH+clientid[0]+"/"+transid[0], os.ModePerm)
		dst, err := os.Create(SERVER_PATH + clientid[0] + "/" + files[i].Filename)
		defer dst.Close()
		if err != nil {
			return err
		}

		if _, err := io.Copy(dst, file); err != nil {
			return err
		}
	}
	return nil
}

func (transformadores) GetImagePath() {

}

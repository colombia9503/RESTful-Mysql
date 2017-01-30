package models

import (
	"io"
	"net/http"
	"os"

	"github.com/colombia9503/RESTful-Mysql/common"
)

type Transformador struct {
	ID               int     `db:"id" json:"ID"`
	Codigo           string  `db:"codigo" json:"Codigo"`
	Cliente          string  `db:"cliente" json:"Cliente"`
	Marca            string  `db:"marca" json:"Marca"`
	KV               float32 `db:"kv" json:"KV"`
	Peso             float32 `db:"peso" json:"Peso"`
	Concentracion    int     `db:"concentracion" json:"Concentracion"`
	Observaciones    string  `db:"observaciones" json:"Observaciones"`
	ImgTransformador string  `db:"img_transformador" json:"ImgTransformador"`
	ImgPlaqueta      string  `db:"img_plaqueta" json:"ImgPlaqueta"`
	ImgBascula       string  `db:"img_bascula" json:"ImgBascula"`
	Borrado          int     `db:"borrado" json:"Borrado"`
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
			&t.Concentracion, &t.Observaciones, &t.ImgTransformador, &t.ImgBascula, &t.ImgPlaqueta,
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
		&t.Concentracion, &t.Observaciones, &t.ImgTransformador, &t.ImgBascula, &t.ImgPlaqueta,
		&t.Borrado)
}

func (transformadores) Insert(t Transformador) error {
	stmt, err := common.DB.Prepare("insert into concentracion (codigo, cliente, marca, kv, peso, concentracion, " +
		"Observaciones, img_transformador, img_bascula, img_plaqueta) " +
		"values(?,?,?,?,?,?,?,?,?,?);")
	if err != nil {
		return err
	}
	_, err = stmt.Exec(t.Codigo, t.Cliente, t.Marca, t.KV, t.Peso,
		t.Concentracion, t.Observaciones, t.ImgTransformador, t.ImgBascula, t.ImgPlaqueta)
	defer stmt.Close()
	return err
}

func (transformadores) Update(t Transformador, id int) error {
	stmt, err := common.DB.Prepare("update concentracion set codigo = ?, cliente = ?, marca = ?, kv = ?, peso = ?, concentracion = ?, " +
		"Observaciones = ?, img_transformador = ?, img_bascula = ?, img_plaqueta = ? " +
		"where id = ? and borrado = 0;")
	if err != nil {
		return err
	}
	_, err = stmt.Exec(t.Codigo, t.Cliente, t.Marca, t.KV, t.Peso,
		t.Concentracion, t.Observaciones, t.ImgTransformador, t.ImgBascula, t.ImgPlaqueta,
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

	codigo := m.Value["Codigo"]
	cliente := m.Value["Cliente"]
	marca := m.Value["Marca"]
	kv := m.Value["KV"]
	peso := m.Value["Peso"]
	concentracion := m.Value["Concentracion"]
	observaciones := m.Value["Observaciones"]

	var imgPaths [3]string

	files := m.File["images"]
	for i, _ := range files {
		file, err := files[i].Open()
		//defer file.Close()
		if err != nil {
			return err
		}
		os.MkdirAll(SERVER_PATH+cliente[0]+"/"+codigo[0], os.ModePerm)
		dst, err := os.Create(SERVER_PATH + cliente[0] + "/" + files[i].Filename)
		defer dst.Close()
		if err != nil {
			return err
		}
		imgPaths[i] = SERVER_PATH + cliente[0] + "/" + files[i].Filename
		if _, err := io.Copy(dst, file); err != nil {
			return err
		}
	}

	stmt, err := common.DB.Prepare("insert into concentracion (codigo, cliente, marca, kv, peso, concentracion, " +
		"Observaciones, img_transformador, img_bascula, img_plaqueta) " +
		"values(?,?,?,?,?,?,?,?,?,?);")
	if err != nil {
		return err
	}
	_, err = stmt.Exec(codigo, cliente, marca, kv, peso,
		concentracion, observaciones, imgPaths[0], imgPaths[1], imgPaths[2])
	defer stmt.Close()
	return err
}

func (transformadores) GetImagePath() {

}

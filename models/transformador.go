package models

import (
	"io"
	"net/http"
	"os"
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
	borrado          int     `db:"borrado"`
}

const MAX_MEMORY = 1 * 1024 * 1024

var Transformadores = new(transformadores)

type transformadores struct{}

func (transformadores) SelectAll() {

}

func (transformadores) SelectOne() {

}

func (transformadores) Insert() {

}

func (transformadores) Update() {

}

func (transformadores) Delete() {

}

func (transformadores) UploadImage(r *http.Request) error {
	if err := r.ParseMultipartForm(MAX_MEMORY); err != nil {
		return err
	}

	m := r.MultipartForm

	files := m.File["image"]
	for i, _ := range files {
		file, err := files[i].Open()
		defer file.Close()
		if err != nil {
			return err
		}

		dst, err := os.Create("/home/Documents/ServerImages" + files[i].Filename)
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

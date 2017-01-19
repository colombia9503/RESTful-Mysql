package models

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

var Transformadores = new(transformadores)

type transformadores struct{}

func SelectAll() {

}

func SelectOne() {

}

func Insert() {

}

func Update() {

}

func Delete() {

}

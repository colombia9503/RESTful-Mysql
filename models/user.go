package models

import (
	"github.com/colombia9503/RESTful-Mysql/common"
)

type User struct {
	Id         int
	First_Name string
	Last_Name  string
}

var Users = new(users)

type users struct{}

func (users) SelectAll() ([]*User, error) {
	var users []*User
	//i don't know how i call the db var from de mysql.go file
	rows, err := common.db.Query("select id, first_name, last_name from person;")
	if err != nil {
		panic(err)
	}

	for rows.Next() {

	}
	return nil, nil
}

func (users) SelectOne(id string) (*User, error) {
	return nil, nil
}

func (users) Insert(name, last_name string) (*User, error) {
	return nil, nil
}

func (users) Update(id, name, last_name string) error {
	return nil
}

func (users) Delete(id string) error {
	return nil
}

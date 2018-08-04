package repository

import (
	"errors"
	"fmt"

	"github.com/neocxf/go-exercises/bookcase/initial"
	"github.com/neocxf/go-exercises/bookcase/models"
)

type SQL struct {
	initial.SQLInstance
}

func (s SQL) InitSchema() {
	fmt.Println("called")
	result, err := s.DB.Exec("create table if not exists user (id INTEGER PRIMARY KEY AUTOINCREMENT, name varchar(255) ); ")

	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(result)
}

func (s SQL) CreateUser(u *models.User) error {

	//Validate the input.
	if u == nil {
		return errors.New("user required")
	} else if u.Name == "" {
		return errors.New("name required")
	}

	fmt.Println(u)

	// Perform the actual insert and return any errors.
	_, err := s.DB.Exec(`INSERT INTO user (name) VALUES (?)`, u.Name)

	return err

}

func (s SQL) CreateUsers(users []*models.User) error {

	//Validate the input.
	if users == nil {
		return errors.New("user required")
	}

	s.Begin()

	for _, u := range users {
		s.CreateUser(u)
	}

	return s.Commit()
}

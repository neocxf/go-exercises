package repository

import (
	"github.com/neocxf/go-exercises/bookcase/initial"
	"github.com/neocxf/go-exercises/bookcase/models"
	"errors"
)

type SQL struct {
	initial.SQLInstance
}

func (s SQL) CreateUser(u *models.User) error {

	//Validate the input.
	if u == nil {
		return errors.New("user required")
	} else if u.Name == "" {
		return errors.New("name required")
	}


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

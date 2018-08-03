package repository

import (
	"context"
	"database/sql"
	"github.com/neocxf/go-exercises/bookcase/author"
	"github.com/neocxf/go-exercises/bookcase/models"
	"github.com/sirupsen/logrus"
)

type mysqlAuthorRepo struct {
	DB *sql.DB
}

func NewMysqlAuthorRepository(db *sql.DB) author.AuthorRepository {
	return &mysqlAuthorRepo{
		DB: db,
	}
}

func (m *mysqlAuthorRepo) getOne(ctx context.Context, query string, args ...interface{}) (*models.Author, error) {
	stmt, err := m.DB.PrepareContext(ctx, query)

	defer m.DB.Close()

	row := stmt.QueryRowContext(ctx, args...)

	a := &models.Author{}

	err = row.Scan(
		&a.ID,
		&a.Name,
		&a.CreatedAt,
		&a.UpdatedAt,
	)

	if err != nil {
		logrus.Error(err)
		return nil, err
	}

	return a, nil

}

func (m *mysqlAuthorRepo) GetById(ctx context.Context, id int64) (*models.Author, error) {
	query := `SELECT id, name, created_at, updated_at from author where id=?`
	return m.getOne(ctx, query, id)
}

package repository

import (
	"context"
	"database/sql"
	"fmt"
	"github.com/neocxf/go-exercises/bookcase/acticle"
	"github.com/neocxf/go-exercises/bookcase/models"
	"github.com/sirupsen/logrus"
)

type mysqlArticleRepo struct {
	DB *sql.DB
}

func NewMysqlAuthorRepository(conn *sql.DB) article.ArticleRepository {
	return &mysqlArticleRepo{conn}
}

func (m *mysqlArticleRepo) fetch(ctx context.Context, query string, args ...interface{}) ([]*models.Article, error) {
	rows, err := m.DB.QueryContext(ctx, query, args...)

	if err != nil {
		logrus.Error(err)
		return nil, err
	}

	defer rows.Close()

	result := make([]*models.Article, 0)

	for rows.Next() {
		t := new(models.Article)
		authorId := int64(0)
		err := rows.Scan(
			&t.ID,
			&t.Title,
			&t.Content,
			&authorId,
			&t.CreatedAt,
			&t.UpdatedAt,
		)

		if err != nil {
			logrus.Error(err)
			return nil, err
		}

		t.Author = models.Author{ID: authorId}

		result = append(result, t)

	}

	return result, nil
}

func (m *mysqlArticleRepo) GetByID(ctx context.Context, id int64) (*models.Article, error) {
	query := `SELECT id, title, content, author_id created_at, updated_at from article where id=?`
	list, err := m.fetch(ctx, query, id)

	if err != nil {
		return nil, err
	}

	ar := &models.Article{}

	if len(list) > 0 {
		ar = list[0]
	} else {
		return nil, models.NOT_FOUND_ERROR
	}

	return ar, nil
}

func (m *mysqlArticleRepo) Fetch(ctx context.Context, cursor string, num int64) ([]*models.Article, error) {
	query := `SELECT id, title, content, author_id created_at, updated_at from article where id=? limit ?`

	return m.fetch(ctx, query, cursor, num)

}

func (m *mysqlArticleRepo) GetByTitle(ctx context.Context, title string) (*models.Article, error) {
	query := `SELECT id, title, content, author_id created_at, updated_at from article where title=?`

	list, err := m.fetch(ctx, query, title)

	if err != nil {
		return nil, err
	}

	ar := &models.Article{}

	if len(list) > 0 {
		ar = list[0]
	} else {
		return nil, models.NOT_FOUND_ERROR
	}

	return ar, nil
}

func (m *mysqlArticleRepo) Update(ctx context.Context, a *models.Article) (*models.Article, error) {
	query := `update article set title=?, content=?, author_id=?,  updated_at=?`

	stmt, err := m.DB.PrepareContext(ctx, query)

	if err != nil {
		return nil, err
	}

	logrus.Debug("Update at ", a.UpdatedAt)

	res, err := stmt.ExecContext(ctx, a.Title, a.Content, a.Author.ID, a.UpdatedAt)

	if err != nil {
		return nil, err
	}

	affect, err := res.RowsAffected()

	if err != nil {
		return nil, err
	}

	if affect != 1 {
		err = fmt.Errorf("weird Behavior.Total affected: %d", affect)
		logrus.Error(err)
		return nil, err
	}

	return a, nil
}

func (m *mysqlArticleRepo) Store(ctx context.Context, a *models.Article) (int64, error) {
	query := `insert into article set title=?, content=?, author_id=?, created_at=?, updated_at=?`

	stmt, err := m.DB.PrepareContext(ctx, query)

	if err != nil {
		return 0, err
	}

	logrus.Debug("Create at ", a.CreatedAt)

	res, err := stmt.ExecContext(ctx, a.Title, a.Content, a.Author.ID, a.CreatedAt, a.UpdatedAt)

	if err != nil {
		return 0, err
	}

	return res.LastInsertId()
}

func (m *mysqlArticleRepo) Delete(ctx context.Context, id int64) (bool, error) {
	query := `DELETE from article where id=?`

	stmt, err := m.DB.PrepareContext(ctx, query)

	if err != nil {
		return false, err
	}

	res, err := stmt.ExecContext(ctx, id)

	if err != nil {
		return false, err
	}

	rowAffected, err := res.RowsAffected()

	if err != nil {
		return false, err
	}

	if rowAffected != 1 {
		err = fmt.Errorf("weird Behavior.Total affected: %d", rowAffected)
		logrus.Error(err)
		return false, err
	}

	return true, nil

}

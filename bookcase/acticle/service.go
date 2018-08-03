package article

import (
	"context"
	"github.com/neocxf/go-exercises/bookcase/models"
)

type ArticleService interface {
	Fetch(ctx context.Context, cursor string, num int64) ([]*models.Article, string, error)
	GetById(ctx context.Context, id int64) (*models.Article, error)
	Update(ctx context.Context, ar *models.Article) (*models.Article, error)
	GetByTitle(ctx context.Context, title string) (*models.Article, error)
	Store(ctx context.Context, ar *models.Article) (*models.Article, error)
	Delete(ctx context.Context, id int64) (bool, error)
}

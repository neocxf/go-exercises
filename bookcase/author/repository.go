package author

import (
	"context"
	"github.com/neocxf/go-exercises/bookcase/models"
)

type AuthorRepository interface {
	GetById(ctx context.Context, id int64) (*models.Author, error)
}

package repo

import (
	"context"
	"lecture10/models"
)

type Repo interface {
	Get(ctx context.Context, title string) (*models.Book, error)
	Create(ctx context.Context, b *models.Book) (int, error)
}

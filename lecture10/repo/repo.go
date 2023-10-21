package repo

import (
	"context"
	"lecture10/models"
)

type Repo interface {
	Get(ctx context.Context, title string) (*models.Book, error)
}

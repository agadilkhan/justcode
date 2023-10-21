package pg

import (
	"context"
	"lecture10/models"
)

func (p *Postgres) Get(ctx context.Context, title string) (*models.Book, error) {
	var b models.Book

	row := p.DB.QueryRow("SELECT * FROM books WHERE title=?", title)
	if err := row.Scan(&b.ID, &b.Genre, &b.Title, &b.Author); err != nil {
		return nil, err

	}

	return &b, nil
}

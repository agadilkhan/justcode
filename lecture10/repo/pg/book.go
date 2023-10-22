package pg

import (
	"context"
	_ "github.com/lib/pq"
	"lecture10/models"
)

func (p *Postgres) Get(ctx context.Context, title string) (*models.Book, error) {
	var b models.Book

	row := p.DB.QueryRow("SELECT * FROM books WHERE title=$1", title)
	if err := row.Scan(&b.ID, &b.Genre, &b.Title, &b.Author); err != nil {
		return nil, err

	}

	return &b, nil
}

func (p *Postgres) Create(ctx context.Context, b *models.Book) (int, error) {
	var id int
	err := p.DB.QueryRow("INSERT INTO books (genre, title, author) VALUES ($1, $2, $3) RETURNING id",
		b.Genre, b.Title, b.Author).Scan(&id)

	if err != nil {
		return 0, err
	}

	return id, nil
}

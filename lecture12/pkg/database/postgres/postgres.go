package postgres

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Postgres struct {
	*gorm.DB
}

func New(url string) (*Postgres, error) {
	db, err := gorm.Open(postgres.Open(url), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	return &Postgres{db}, nil
}

func (p *Postgres) Close() error {
	db, _ := p.DB.DB()

	return db.Close()
}

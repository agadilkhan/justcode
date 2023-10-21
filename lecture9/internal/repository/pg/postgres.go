package pg

import (
	"fmt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
)

type Postgres struct {
	Host     string
	Port     string
	Username string
	Password string
	SSLMode  string
	DB       *gorm.DB
}

func NewPostgres(opts ...Option) (*Postgres, error) {
	p := new(Postgres)

	for _, opt := range opts {
		opt(p)
	}

	dsn := fmt.Sprintf("host=%s user=%s port=%s password=%s sslmode=%s",
		p.Host, p.Username, p.Port, p.Password, p.SSLMode)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	return &Postgres{DB: db}, nil
}

func (p *Postgres) Close() {
	db, err := p.DB.DB()
	if err != nil {
		log.Printf("error closing database: %s", err.Error())
	}

	db.Close()
}

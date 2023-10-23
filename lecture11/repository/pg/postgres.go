package pg

import (
	"fmt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
)

type Postgres struct {
	host     string
	username string
	password string
	port     string
	dbName   string
	sslmode  string
	DB       *gorm.DB
}

func New(opts ...Option) (*Postgres, error) {
	p := new(Postgres)

	for _, opt := range opts {
		opt(p)
	}

	dsn := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=%s",
		p.host, p.port, p.username, p.password, p.dbName, p.sslmode)

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

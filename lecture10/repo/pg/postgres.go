package pg

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
)

type Postgres struct {
	Host     string
	Port     string
	Username string
	Password string
	DBName   string
	SSLMode  string
	*sql.DB
}

func New(opts ...Option) (*Postgres, error) {
	p := new(Postgres)

	for _, opt := range opts {
		opt(p)
	}

	dsn := fmt.Sprintf("postgres://%s:%s@%s/%s?sslmode=%s",
		p.Username, p.Password, p.Host, p.DBName, p.SSLMode)

	db, err := sql.Open("postgres", dsn)
	if err != nil {
		return nil, err
	}

	err = db.Ping()

	return &Postgres{DB: db}, err
}

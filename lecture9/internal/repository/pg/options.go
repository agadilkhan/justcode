package pg

type Option func(p *Postgres)

func WithHost(host string) Option {
	return func(p *Postgres) {
		p.Host = host
	}
}

func WithPort(port string) Option {
	return func(p *Postgres) {
		p.Port = port
	}
}

func WithUsername(username string) Option {
	return func(p *Postgres) {
		p.Username = username
	}
}

func WithPassword(password string) Option {
	return func(p *Postgres) {
		p.Password = password
	}
}

func WithSSLMode(sslmode string) Option {
	return func(p *Postgres) {
		p.SSLMode = sslmode
	}
}

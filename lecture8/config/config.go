package config

import (
	"github.com/ilyakaznacheev/cleanenv"
	"time"
)

type Config struct {
	HTTP HTTPConfig `yaml:"httpserver"`
	DB   DBConfig   `yaml:"db"`
}

type HTTPConfig struct {
	Port            string        `yaml:"port"`
	ShutdownTimeout time.Duration `yaml:"shutdown_timeout"`
}

type DBConfig struct {
	Host     string `yaml:"host"`
	Port     string `yaml:"port"`
	Username string `yaml:"username"`
	Password string `env-required:"true" env:"DB_PASSWORD"`
	DBName   string `yaml:"db_name"`
	SSLMode  string `yaml:"sslmode"`
}

func Init(path string) (*Config, error) {
	cfg := new(Config)

	err := cleanenv.ReadConfig(path, cfg)
	if err != nil {
		return nil, err
	}

	err = cleanenv.ReadEnv(cfg)

	return cfg, err
}

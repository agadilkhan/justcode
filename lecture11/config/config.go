package config

import (
	"github.com/ilyakaznacheev/cleanenv"
	"github.com/spf13/viper"
	"time"
)

type Config struct {
	HTTP HTTPConfig `yaml:"http"`
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

func NewViperConfig() (*Config, error) {
	cfg := Config{}

	viper.SetConfigName("config")
	viper.SetConfigType("yml")
	viper.AddConfigPath("./config/")
	err := viper.ReadInConfig()
	if err != nil {
		return nil, err
	}

	err = viper.Unmarshal(&cfg)

	return &cfg, err
}

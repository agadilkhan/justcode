package config

import (
	"github.com/spf13/viper"
	"time"
)

type Config struct {
	HTTP `yaml:"http"`
	PG   `yaml:"pg"`
	JWT  `yaml:"jwt"`
}

type HTTP struct {
	Port            string        `yaml:"port"`
	ShutdownTimeout time.Duration `mapstructure:"shutdown_timeout"`
	ReadTimeout     time.Duration `mapstructure:"read_timeout"`
	WriteTimeout    time.Duration `mapstructure:"write_timeout"`
}

type PG struct {
	URL string `env-required:"true" env:"PG_URL"`
}

type JWT struct {
	SecretKey string `mapstructure:"secret_key"`
}

func New(path string) (*Config, error) {
	cfg := &Config{}

	viper.AddConfigPath(path)
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")

	err := viper.ReadInConfig()
	if err != nil {
		return nil, err
	}

	viper.AutomaticEnv()
	dbURL := viper.GetString("PG_URL")

	cfg.PG.URL = dbURL

	err = viper.Unmarshal(&cfg)

	return cfg, err
}

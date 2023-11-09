package config

import (
	"log"
	"time"

	"github.com/ilyakaznacheev/cleanenv"
)

type config struct {
	Env        string `yaml:"env" env-default:"local"`
	HttpServer struct {
		Port        string        `yaml:"port"`
		Timeout     time.Duration `yaml:"timeout"`
		IdleTimeout time.Duration `yaml:"idle_timeout"`
	} `yaml:"http_server"`
}

func MustLoad() *config {
	var cfg config

	if err := cleanenv.ReadConfig("config.yaml", &cfg); err != nil {
		log.Fatalf("failed to read config: %v", err)
	}

	return &cfg
}

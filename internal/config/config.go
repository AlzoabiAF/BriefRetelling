package config

import (
	"github.com/ilyakaznacheev/cleanenv"
	"log"
	"os"
	"time"
)

type Config struct {
	Env string `yaml:"env"`
	StorageParam
	HTTPServer
}

type HTTPServer struct {
	Address      string        `yaml:"address"`
	Timeout      time.Duration `yaml:"timeout"`
	Idle_Timeout time.Duration `yaml:"idle_timeout"`
}
type StorageParam struct {
	Username string `yaml:"username"`
	Password string `yaml:"password"`
	Dbname   string `yaml:"dbname"`
	Sslmode  string `yaml:"sslmode"`
}

func MustLoad() *Config {
	const op = "./internal/config"
	configPath := os.Getenv("CONFIG_PATH")
	if _, err := os.Stat(configPath); os.IsNotExist(err) {
		log.Fatalf("Env file not found: %s: %w", op, err)
	}
	var cfg Config
	if err := cleanenv.ReadConfig(configPath, &cfg); err != nil {
		log.Fatalf("cannot read config: %s", op)
	}
	return &cfg
}

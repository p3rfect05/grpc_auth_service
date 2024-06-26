package config

import (
	"errors"
	"os"
	"time"

	"github.com/ilyakaznacheev/cleanenv"
)

type Config struct {
	Env      string        `yaml:"env" env-default:"local"`
	Storage  string        `yaml:"storage_path" env-required:"true"`
	GRPC     GRPCConfig    `yaml:"grpc"`
	TokenTTL time.Duration `yaml:"token_ttl" env-default:"1h"`
}

type GRPCConfig struct {
	Port    int           `yaml:"port" env-required:"true"`
	Timeout time.Duration `yaml:"timeout" env-required:"true"`
}

type DatabaseURLConfig struct {
	Host     string
	User     string
	Password string
	DBName   string
	Port     string
}

func MustLoadDatabaseURL() DatabaseURLConfig {
	host := os.Getenv("POSTGRES_HOST")
	user := os.Getenv("POSTGRES_USER")
	password := os.Getenv("POSTGRES_PASSWORD")
	db_name := os.Getenv("POSTGRES_DB")
	port := os.Getenv("POSTGRES_PORT")
	
	return DatabaseURLConfig{
		Host:     host,
		User:     user,
		Password: password,
		DBName:   db_name,
		Port:     port,
	}
}
func MustLoad() *Config {
	configPath := fetchConfigPath()
	if configPath == "" {
		panic("no config path provided")
	}
	if _, err := os.Stat(configPath); errors.Is(err, os.ErrNotExist) {
		panic("config file does not exist: " + configPath)
	}
	var cfg Config

	if err := cleanenv.ReadConfig(configPath, &cfg); err != nil {
		panic("cannot read config: " + err.Error())
	}

	return &cfg

}

func fetchConfigPath() string {
	res := os.Getenv("CONFIG_PATH")
	return res
}

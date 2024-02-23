package config

import (
	"fmt"
	"os"

	"github.com/caarlos0/env/v10"
	"github.com/joho/godotenv"
)

type appEnv string

const (
	EnvDev appEnv = "dev"
	EnvPrd appEnv = "prd"
)

type (
	mysql struct {
		Database string `env:"MYSQL_DATABASE" envDefault:"go_database"`
		Port     int    `env:"MYSQL_PORT" envDefault:"3306"`
		Host     string `env:"MYSQL_HOST"`
		User     string `env:"MYSQL_USER"`
		Password string `env:"MYSQL_PASSWORD"`
		TZ       string `env:"TZ" envDefault:"Asia/Tokyo"`
	}

	redis struct {
		Host     string `env:"REDIS_HOST"`
		Port     int    `env:"REDIS_PORT" envDefault:"6379"`
		Password string `env:"REDIS_PASSWORD"`
		DB       int    `env:"REDIS_DB"`
	}
)

var (
	MySQL mysql
	Reids redis
)

func getEnv() appEnv {
	switch env := os.Getenv("ENV"); env {
	case "prod":
		return EnvPrd
	default:
		return EnvDev
	}
}

func init() {
	Load()
}

func IsDev() bool {
	return getEnv() == EnvDev
}

func Load() {
	if IsDev() {
		err := godotenv.Load("env/.env.dev")
		if err != nil {
			panic(fmt.Sprintf("Error loading .env file: %v", err))
		}
	}
	loadMySQLCfg()
	loadRedisCfg()
}

func loadMySQLCfg() {
	if err := env.Parse(&MySQL); err != nil {
		panic(err)
	}
}

func loadRedisCfg() {
	if err := env.Parse(&Reids); err != nil {
		panic(err)
	}
}

package config

import (
	"fmt"
	"os"

	"github.com/caarlos0/env/v10"
	"github.com/joho/godotenv"
)

type appEnv string

type mysqlCfg struct {
	Database string `env:"MYSQL_DATABASE" envDefault:"go_database"`
	Port     int    `env:"MYSQL_PORT" envDefault:"3306"`
	Host     string `env:"MYSQL_HOST"`
	User     string `env:"MYSQL_USER"`
	Password string `env:"MYSQL_PASSWORD"`
	TZ       string `env:"TZ" envDefault:"Asia/Tokyo"`
}

var MySQLCfg mysqlCfg

const (
	EnvDev appEnv = "dev"
	EnvPrd appEnv = "prd"
)

func getEnv() appEnv {
	switch env := os.Getenv("ENV"); env {
	case "prd":
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
	if err := env.Parse(&MySQLCfg); err != nil {
		panic(err)
	}
}

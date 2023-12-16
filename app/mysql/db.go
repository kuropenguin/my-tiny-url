package mysql

import (
	"database/sql"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

var (
	db *sql.DB
)

func init() {
	var err error
	db, err = sql.Open("mysql", "admin:password@tcp(tiny-url.cgssfxuiz3lu.ap-northeast-1.rds.amazonaws.com:3306)/go_database")
	if err != nil {
		panic(err)
	}
	db.SetConnMaxLifetime(time.Minute * 3)
	db.SetMaxOpenConns(10)
	db.SetMaxIdleConns(10)

	if err := db.Ping(); err != nil {
		panic(err)
	}
}

func GetDB() *sql.DB {
	return db
}

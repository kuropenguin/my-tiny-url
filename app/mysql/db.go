package mysql

import (
	"database/sql"
	"fmt"
	"time"

	"log"

	_ "github.com/go-sql-driver/mysql"
	"github.com/kuropenguin/my-tiny-url/app/config"
)

var (
	db *sql.DB
)

func init() {
	// TODO DELETE ME
	for {
		fmt.Println("Connecting to MySQL ...")
		time.Sleep(5 * time.Second)
	}
	var err error
	// db, err = sql.Open("mysql", "admin:password@tcp(tiny-url.cgssfxuiz3lu.ap-northeast-1.rds.amazonaws.com:3306)/go_database")
	db, err = sql.Open("mysql", fmt.Sprintf("%s:%s@tcp(%s:%d)/%s", config.MySQL.User, config.MySQL.Password, config.MySQL.Host, config.MySQL.Port, config.MySQL.Database))
	if err != nil {
		panic(err)
	}
	db.SetConnMaxLifetime(time.Minute * 3)
	db.SetMaxOpenConns(10)
	db.SetMaxIdleConns(10)

	if err := db.Ping(); err != nil {
		log.Printf("Error connecting to the database: %s", err)
		// panic(err)
	}
}

func GetDB() *sql.DB {
	return db
}

package users_db

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"log"
	"os"
)

var (
	Db *sql.DB
)

func init() {
	var err error
	dataSource := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8",
		os.Getenv("MYSQL_USER"),
		os.Getenv("MYSQL_PWD"),
		os.Getenv("MYSQL_HOST"),
		"suboi",
	)

	Db, err = sql.Open("mysql", dataSource)

	if err != nil {
		panic(err)
	}

	if err = Db.Ping(); err != nil {
		panic(err)
	}

	log.Println("database successfully connected.")
}

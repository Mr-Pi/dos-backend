package pgsql

import (
	"database/sql"
	"github.com/Mr-Pi/dos-backend/config"
	_ "github.com/lib/pq"
	"log"
)

var db *sql.DB

func testErr(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func testWarn(err error) {
	if err != nil {
		log.Println(err)
	}
}

func Connect(cfg config.Config) {
	var err error
	db, err = sql.Open("postgres", cfg.PGsql.Url)
	testErr(err)
	db.Ping()
}

func ListUsers() []string {
	var users []string
	rows, err := db.Query(`SELECT username FROM customer;`)
	testWarn(err)
	for rows.Next() {
		var user string
		rows.Scan(&user)
		users = append(users, user)
	}
	return users
}

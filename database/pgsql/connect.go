package pgsql

import (
	"database/sql"
	"github.com/Mr-Pi/dos-backend/config"
	"github.com/Mr-Pi/dos-backend/types"
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

func GETUser(username string) types.User {
	var user types.User
	var permID int64
	db.QueryRow(`SELECT username, firstname, lastname, perms, credit FROM customer WHERE username=$1;`, username).Scan(&user.Username, &user.FirstName, &user.LastName, &permID, &user.Credit)
	db.QueryRow(`SELECT type, patchdrinkeveryone, modsupplier, moddrink, moduser, setownpassword FROM perms WHERE id=$1;`, permID).Scan(&user.Permissions.Type, &user.Permissions.PatchDrinkEveryone, &user.Permissions.ModSuppliers, &user.Permissions.ModDrink, &user.Permissions.ModUser, &user.Permissions.SetOwnPassword)
	return user
}

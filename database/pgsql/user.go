package pgsql

import (
	"github.com/Mr-Pi/dos-backend/types"
	"log"
)

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

func TestUser(username string) bool {
	err := db.QueryRow(`SELECT FROM customer WHERE username=$1`, username).Scan()
	testWarn(err)
	if err != nil {
		log.Println("User not found", username)
		return false
	} else {
		return true
	}
}

func GETUser(username string) types.User {
	var user types.User
	var permID int64
	err := db.QueryRow(`SELECT username, firstname, lastname, perms, credit FROM customer WHERE username=$1;`, username).Scan(&user.Username, &user.FirstName, &user.LastName, &permID, &user.Credit)
	testWarn(err)
	db.QueryRow(`SELECT type, patchdrinkeveryone, modsupplier, moddrink, moduser, setownpassword FROM perms WHERE id=$1;`, permID).Scan(&user.Permissions.Type, &user.Permissions.PatchDrinkEveryone, &user.Permissions.ModSuppliers, &user.Permissions.ModDrink, &user.Permissions.ModUser, &user.Permissions.SetOwnPassword)
	testWarn(err)
	return user
}

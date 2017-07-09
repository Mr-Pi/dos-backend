package pgsql

import (
	"github.com/Mr-Pi/dos-backend/types"
	"log"
)

func CreateUser(user types.User) (err error) {
	_, err = db.Query(`INSERT INTO customer ( username, firstname, lastname, perms, password, salt ) VALUES ($1,$2,$3,$4,$5,$6);`, user.Username, user.FirstName, user.LastName, user.Permissions, user.Password, user.Salt)
	return err
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
	err := db.QueryRow(`SELECT username, firstname, lastname, perms, credit, password, salt FROM customer WHERE username=$1;`, username).Scan(
		&user.Username,
		&user.FirstName,
		&user.LastName,
		&user.Permissions,
		&user.Credit,
		&user.Password,
		&user.Salt,
	)
	testWarn(err)
	return user
}

func DecrementUserCredit(username string, amount float64) (err error) {
	stmt, err := db.Prepare(`UPDATE customer set credit = credit - $1 WHERE username = $2;`)
	if err != nil {
		return
	}
	defer stmt.Close()
	_, err = stmt.Exec(amount, username)
	return

}

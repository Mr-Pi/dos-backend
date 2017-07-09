package pgsql

import (
	"github.com/Mr-Pi/dos-backend/types"
	"log"
)

func ListPermissions() []string {
	var permissions []string
	rows, err := db.Query(`SELECT type FROM perms;`)
	testWarn(err)
	for rows.Next() {
		var permType string
		rows.Scan(&permType)
		permissions = append(permissions, permType)
	}
	return permissions
}

func TestPermissions(permType string) bool {
	err := db.QueryRow(`SELECT FROM perms WHERE type=$1`, permType).Scan()
	testWarn(err)
	if err != nil {
		log.Println("Permission type not found", permType)
		return false
	} else {
		return true
	}
}

func GETPermissions(permType string) types.UserPermissions {
	var permissions types.UserPermissions
	err := db.QueryRow(`SELECT type, patchdrinkeveryone, modsupplier, moddrink, moduser, setownpassword FROM perms WHERE type=$1;`, permType).Scan(
		&permissions.Type,
		&permissions.PatchDrinkEveryone,
		&permissions.ModSuppliers,
		&permissions.ModDrink,
		&permissions.ModUser,
		&permissions.SetOwnPassword,
	)
	testWarn(err)
	return permissions
}

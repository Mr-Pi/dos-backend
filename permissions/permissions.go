package permissions

import (
	"github.com/Mr-Pi/dos-backend/types"
	"github.com/Mr-Pi/dos-backend/database/pgsql"
)

func CheckUserPermissions(username string, reqPerm types.UserPermissions) (rc int) {
	rc = 200
	perm := pgsql.GETUser(username).Permissions
	checkPerm(perm.PatchDrinkEveryone, reqPerm.PatchDrinkEveryone, &rc)
	if rc != 200 {
		return
	}
	checkPerm(perm.ModUser, reqPerm.ModUser, &rc)
	if rc != 200 {
		return
	}
	checkPerm(perm.ModSuppliers, reqPerm.ModSuppliers, &rc)
	if rc != 200 {
		return
	}
	checkPerm(perm.ModDrink, reqPerm.ModDrink, &rc)
	if rc != 200 {
		return
	}
	checkPerm(perm.SetOwnPassword, reqPerm.SetOwnPassword, &rc)
	return
}

func checkPerm(val, req bool, rc *int) {
	if req {
		if val {
			*rc = 200
		} else {
			*rc = 401
		}
	}
}

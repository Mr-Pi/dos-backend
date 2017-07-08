package types

type User struct {
	Username string
	FullName string
	Credit float32
	Password string
	Permissions UserPermissions
}
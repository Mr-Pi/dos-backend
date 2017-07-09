package types

type UserPermissions struct {
	Type               string
	PatchDrinkEveryone bool
	ModSuppliers       bool
	ModDrink           bool
	ModUser            bool
	SetOwnPassword     bool
}

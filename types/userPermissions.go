package types

type UserPermissions struct {
	PatchDrinkEveryone bool
	ModSuppliers       bool
	ModDrink           bool
	ModUser            bool
	SetOwnPass         bool
}

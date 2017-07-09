package types

type UserPermissions struct {
	Type               string `json:"name"`
	PatchDrinkEveryone bool `json:"patchDrinkEveryone"`
	ModSuppliers       bool `json:"modifySuppliers"`
	ModDrink           bool `json:"modifyDrinks"`
	ModUser            bool `json:"modifyUsers"`
	SetOwnPassword     bool `json:"setOwnPassword"`
}

package types

type User struct {
	Username    string `json:"uid"`
	FirstName   string `json:"givenName"`
	LastName    string `json:"surname"`
	Credit      float64 `json:"credit"`
	Password    string `json:"-"`
	RFIDTag     []string `json:"rfidTags"`
	Permissions UserPermissions `json:"permissions"`
}

package types

type User struct {
	Username    string
	FirstName   string
	LastName    string
	Credit      float64
	Password    string `json:"-"`
	RFIDTag     []string
	Permissions string
	Salt        string
}

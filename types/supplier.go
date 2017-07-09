package types

type Supplier struct {
	ID          int64 `json:"id"`
	Name        string `json:"name"`
	Address     string `json:"address"`
	DeliverTime int64 `json:"deliveryTime"`
}

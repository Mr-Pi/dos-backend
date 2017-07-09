package types

type Drink struct {
	EAN             string `json:"ean"`
	Name            string `json:"name"`
	PriceOrder      float64 `json:"priceOrder"`
	PriceResell     float64 `json:"priceResell"`
	Amount          int64 `json:"amount"`
	Supplier        Supplier `json:"supplier"`
	RedeliverAmount int64 `json:"redeliverAmount"`
	ImgUrl          string `json:"imageUrl"`
}

package types

type Drink struct {
	EAN             string
	Name            string
	PriceOrder      float64
	PriceResell     float64
	Amount          int64
	Supplier        Supplier
	RedeliverAmount int64
	ImgUrl          string
}

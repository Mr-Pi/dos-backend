package types

type Drink struct {
	Name string
	PriceOrder float32
	PriceResell float32
	Amount uint32
	Supplier Supplier
	RedeliverAmount uint32
}

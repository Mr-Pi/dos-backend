package pgsql

import (
	"github.com/Mr-Pi/dos-backend/types"
	"log"
)

func ListDrinks() []string {
	var drinks []string
	rows, err := db.Query(`SELECT EAN FROM drink;`)
	testWarn(err)
	for rows.Next() {
		var drinkEAN string
		rows.Scan(&drinkEAN)
		drinks = append(drinks, drinkEAN)
	}
	return drinks
}

func TestDrink(drink string) bool {
	err := db.QueryRow(`SELECT FROM drink WHERE ean=$1`, drink).Scan()
	testWarn(err)
	if err != nil {
		log.Println("Drink not found", drink)
		return false
	} else {
		return true
	}
}

func GETDrink(drinkEAN string) types.Drink {
	var drink types.Drink
	err := db.QueryRow(`SELECT ean, name, amount, redeliveramount, priceorder, priceresell, imgurl FROM drink WHERE ean=$1;`, drinkEAN).Scan(
		&drink.EAN,
		&drink.Name,
		&drink.Amount,
		&drink.RedeliverAmount,
		&drink.PriceOrder,
		&drink.PriceResell,
		&drink.ImgUrl,
	)
	testWarn(err)
	return drink
}

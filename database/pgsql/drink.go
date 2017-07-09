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
	var supplierID int64
	err := db.QueryRow(`SELECT ean, name, amount, redeliveramount, priceorder, priceresell, imgurl, supplier FROM drink WHERE ean=$1;`, drinkEAN).Scan(
		&drink.EAN,
		&drink.Name,
		&drink.Amount,
		&drink.RedeliverAmount,
		&drink.PriceOrder,
		&drink.PriceResell,
		&drink.ImgUrl,
		&supplierID,
	)
	drink.Supplier = GETSupplier(supplierID)
	testWarn(err)
	return drink
}

// TODO Test
func CreateDrink(drink types.Drink) (success bool) {
	success = false
	stmt, err := db.Prepare(`INSERT INTO drink (ean, name, amount, supplier, redeliveramount, priceorder, priceresell, imgurl)
		VALUES ($1, $2, $3, $4, $5, $6, $7, $8);`)
	if err != nil {
		return
	}
	defer stmt.Close()
	_, err = stmt.Exec(drink.EAN, drink.Name, drink.Amount, drink.Supplier, drink.RedeliverAmount, drink.PriceOrder, drink.PriceResell, drink.ImgUrl)
	if err != nil {
		return
	}
	success = true
	return
}

func DecrementDrinkAmount(drinkEAN string, amount uint32) (err error) {
	stmt, err := db.Prepare(`UPDATE drink SET amount = amount - $1 WHERE ean = $2;`)
	if err != nil {
		return
	}
	defer stmt.Close()
	_, err = stmt.Exec(amount, drinkEAN)
	return
}


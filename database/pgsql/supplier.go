package pgsql

import (
	"github.com/Mr-Pi/dos-backend/types"
	"log"
)

func ListSuppliers() []int64 {
	var suppliers []int64
	rows, err := db.Query(`SELECT id FROM supplier;`)
	testWarn(err)
	for rows.Next() {
		var supplier int64
		rows.Scan(&supplier)
		suppliers = append(suppliers, supplier)
	}
	return suppliers
}

func TestSupplier(supplier int64) bool {
	err := db.QueryRow(`SELECT FROM supplier WHERE id=$1`, supplier).Scan()
	testWarn(err)
	if err != nil {
		log.Println("Supplier not found", supplier)
		return false
	} else {
		return true
	}
}

func GETSupplier(supplierID int64) types.Supplier {
	var supplier types.Supplier
	err := db.QueryRow(`SELECT id, address, name, delivertime FROM supplier WHERE id=$1;`, supplierID).Scan(
		&supplier.ID,
		&supplier.Address,
		&supplier.Name,
		&supplier.DeliverTime,
	)
	testWarn(err)
	return supplier
}

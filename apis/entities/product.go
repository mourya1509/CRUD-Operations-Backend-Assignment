package entities

import (
	"fmt"
)

type Product struct {
	Id       int
	Name     string
	Price    float32
	Quantity int
}

func (product Product) ToString() string {

	return fmt.Sprintf("id: %d\nname: %s\nprice: %0.1f\nquantity: %d", product.Id, product.Name, product.Price, product.Quantity)
}

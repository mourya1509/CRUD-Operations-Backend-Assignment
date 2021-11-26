package models

import (
	"apis/entities"
	"database/sql"
	"fmt"
	"strconv"
)

type ProductModel struct {
	Db *sql.DB
}

func (productModel ProductModel) FindAll() (product []entities.Product, err error) {
	rows, err := productModel.Db.Query("select * from product")
	if err != nil {
		return nil, err
	} else {
		var products []entities.Product
		for rows.Next() {
			var id int
			var name string
			var price float32
			var quantity int
			err2 := rows.Scan(&id, &name, &price, &quantity)
			if err2 != nil {
				return nil, err2
			} else {
				product := entities.Product{
					Id:       id,
					Name:     name,
					Price:    price,
					Quantity: quantity,
				}
				products = append(products, product)

			}
		}
		return products, nil
	}
}

func (productModel ProductModel) Create(product *entities.Product) (err error) {
	result, err := productModel.Db.Exec("insert into product(id, name, price, quantity) values(" + strconv.Itoa(product.Id) + ",'" + product.Name + "', " + fmt.Sprintf("%f", product.Price) + ", " + strconv.Itoa(product.Quantity) + ")")
	if err != nil {
		return err
	} else {
		fmt.Println(result)
		//product.Id, _ = result.LastInsertId()
		return nil

	}
}

func (productModel ProductModel) Update(product *entities.Product) (int64, error) {
	result, err := productModel.Db.Exec("update product set name = '" + product.Name + "', price = " + fmt.Sprintf("%f", product.Price) + ", quantity =" + strconv.Itoa(product.Quantity) + " where id = " + strconv.Itoa(product.Id) + "")
	if err != nil {
		return 0, err
	} else {
		return result.RowsAffected()

	}

}

func (productModel ProductModel) Delete(id int64) (int64, error) {
	var id1 = strconv.FormatInt(id, 10)
	result, err := productModel.Db.Exec("delete from product where id= " + id1 + "")
	if err != nil {
		return 0, err
	} else {
		return result.RowsAffected()

	}

}

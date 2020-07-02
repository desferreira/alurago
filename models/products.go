package models

import (
	db "github.com/desferreira/alurago/db"
)

type Product struct {
	Id int
	Name string
	Description string
	Price float64
	Quantity int
}

func GetAll() []Product {
	conn := db.CreateConnection()
	defer db.CloseConnection(conn)

	allProducts, err := conn.Query("select * from products")
	if err != nil {
		panic(err.Error())
	}

	products := []Product{}

	for allProducts.Next() {
		var id, quantity int
		var name, description string
		var price float64
		var p Product
		err = allProducts.Scan(&id, &name, &description, &price, &quantity)
		p.Id = id
		p.Name = name
		p.Description = description
		p.Price = price
		p.Quantity = quantity
		products = append(products, p)

	}
	return products
}

func Create(name, description string, price float64, quantity int){
	conn := db.CreateConnection()
	defer db.CloseConnection(conn)
	insertString, err := conn.Prepare("insert into products (Name, Description, Price, Quantity) values ($1, $2, $3, $4)")
	if err != nil {
		panic(err)
	}
	insertString.Exec(name, description, price, quantity)
}

func Delete(id int){
	conn := db.CreateConnection()
	defer db.CloseConnection(conn)
	deleteString, err := conn.Prepare("delete from products where id = $1")
	if err != nil {
		panic(err)
	}
	deleteString.Exec(id)
}
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

func Edit(id int, name, description string, price float64, quantity int){
	conn := db.CreateConnection()
	defer db.CloseConnection(conn)

	updateString, err := conn.Prepare("update products set name=$1, description=$2, price=$3, quantity=$4 where id=$5")
	if err != nil {
		panic(err)
	}
	updateString.Exec(name, description, price, quantity, id)
}

func FindById(id int) Product{
	conn := db.CreateConnection()
	defer db.CloseConnection(conn)

	selectString, err := conn.Prepare("select * from products where id=$1")
	if err != nil {
		panic(err)
	}

	rows, err := selectString.Query(id)
	if err != nil {
		panic(err)
	}
	var p Product
	for rows.Next() {

		var id, quantity int
		var name, description string
		var price float64
		rows.Scan(&id, &name, &description, &price, &quantity)
		p.Id = id
		p.Name = name
		p.Description = description
		p.Price = price
		p.Quantity = quantity
	}
	return p
}
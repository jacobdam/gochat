package models

type Product struct {
	Name  string `json:"name"`
	Price int    `json:"price"`
}

var AllProducts = []Product{
	Product{"Car", 10000},
	Product{"Pen", 10},
}

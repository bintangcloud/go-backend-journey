package main

import (
	"encoding/json"
	"fmt"
)

type Customer struct {
	Name  string `json:"name"`
	Email string `json:"email"`
}

type Items struct {
	ProductName string `json:"product_name"`
	Price       int    `json:"price"`
	Quantity    int    `json:"quantity"`
}

type Order struct {
	ID       int      `json:"id"`
	Customer Customer `json:"customer"`
	Items    []Items  `json:"items"`
}

func main() {
	p1 := Order{
		ID: 101,
		Customer: Customer{
			Name:  "Bintang",
			Email: "bintang@gmail.com",
		},
		Items: []Items{
			{ProductName: "Novel", Price: 200000, Quantity: 1},
			{ProductName: "Pencil", Price: 2000, Quantity: 2},
		},
	}

	jsonData, _ := json.Marshal(p1)
	fmt.Println("Hasil JSON:", string(jsonData))
}

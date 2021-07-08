package main

import "fmt"

type Product struct {
	Id       int
	Name     string
	Cost     float32
	Units    int
	Category string
}

/*
	type PerishableProduct struct {
		Prod   Product
		Expiry string
	}
*/

type PerishableProduct struct {
	Product
	Expiry string
	Cost   float32 //override
}

func main() {
	// var p Product
	/*
		grapes := PerishableProduct{Prod: Product{102, "Grapes", 10, 20, "Fruits"}, Expiry: "5 days"}
		fmt.Println(grapes)
		fmt.Println("Cost of grapes =>", grapes.Prod.Cost)
	*/
	/*
		//grapes := PerishableProduct{Product: Product{102, "Grapes", 10, 20, "Fruits"}, Expiry: "5 days"}
		grapes := PerishableProduct{Product{102, "Grapes", 10, 20, "Fruits"}, "5 days", 70}
		fmt.Println(grapes)
		fmt.Println("Cost of grapes =>", grapes.Cost)
	*/

	grapes := NewPersiableProduct(102, "Grapes", 100, 20, "Fruits", "5 days")
	fmt.Println(grapes)
	fmt.Println("Cost of grapes =>", grapes.Cost)
}

func NewPersiableProduct(id int, name string, cost float32, units int, category string, expiry string) *PerishableProduct {
	return &PerishableProduct{Product{id, name, cost, units, category}, expiry, cost}
}

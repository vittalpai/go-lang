package main

import "fmt"

type Product struct {
	Id       int
	Name     string
	Cost     float32
	Units    int
	Category string
}

func main() {
	// var p Product
	//var p Product = Product{Id: 100, Name: "Pen", Cost: 21321, Units: 12321, Category: "Stationary"}
	var p Product = Product{102, "Pen", 10, 20, "Stationary"}
	pAddr := &p
	fmt.Println(p)
	fmt.Println("Cost of Pen =>", p.Cost)
	fmt.Println("Id of Pen =>", (*pAddr).Id)
	fmt.Println("Id of Pen =>", pAddr.Id)

	fmt.Println("Product cost before discount =>", p.Cost)
	applyDiscount(&p, 5)
	fmt.Println("Product cost after discount =>", p.Cost)

}

func applyDiscount(p *Product, discount int) {
	p.Cost = p.Cost - float32(discount)
}

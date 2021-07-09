package main

import (
	"fmt"
	"sort"
)

type Product struct {
	Id       int
	Name     string
	Cost     float32
	Units    int
	Category string
}

/*
func (product *Product) Format() string {
	return fmt.Sprintf("Id = %d, Name = %s, Cost = %v, Units = %d, Category = %s\n", product.Id, product.Name, product.Cost, product.Units, product.Category)
}
*/

/* implementing the Stringer interface */
func (product Product) String() string {
	return fmt.Sprintf("Id = %d, Name = %s, Cost = %v, Units = %d, Category = %s\n", product.Id, product.Name, product.Cost, product.Units, product.Category)
}

//implement the applyDiscount "method" for the product
func (product *Product) ApplyDiscount(discount float32) {
	product.Cost = product.Cost * ((100 - discount) / 100)
}

type PerishableProduct struct {
	Product
	Expiry string
}

func (pp PerishableProduct) String() string {
	//return fmt.Sprintf("Id = %d, Name = %s, Cost = %v, Units = %d, Category = %s, Expiry = %s\n", pp.Id, pp.Name, pp.Cost, pp.Units, pp.Category, pp.Expiry)
	//using the Product.Format()
	return fmt.Sprintf("%s, Expiry = %s\n", pp.Product, pp.Expiry)
}

//products

type Products []Product

func (products Products) Len() int {
	return len(products)
}

func (products Products) Less(i, j int) bool {
	return products[i].Id < products[j].Id
}

func (products Products) Swap(i, j int) {
	products[i], products[j] = products[j], products[i]
}

/*
	func (products Products) Sort() {
		sort.Sort(products)
	}
*/

// To Sort by Bame
type ByName struct {
	Products
}

func (products ByName) Less(i, j int) bool {
	return products.Products[i].Name < products.Products[j].Name
}

/* Products Methods */

func (products Products) Sort(attrName string) {
	switch attrName {
	case "Id":
		sort.Sort(products)
	case "Name":
		sort.Sort(ByName{products})
	default:
		sort.Sort(products)
	}
}

func main() {
	pencil := Product{107, "Pencil", 2, 100, "Stationary"}
	//fmt.Print(Format(pencil))
	fmt.Print(pencil)
	pencil.ApplyDiscount(10)
	fmt.Print(pencil)

	//manipulating products

	products := Products{
		Product{105, "Pen", 5, 50, "Stationary"},
		Product{107, "Pencil", 2, 100, "Stationary"},
		Product{103, "Marker", 50, 20, "Utencil"},
		Product{102, "AB", 5000, 5, "Utencil"},
		Product{101, "Kettle", 2500, 10, "Stationary"},
		Product{104, "Scribble Pad", 20, 20, "Stationary"},
	}
	fmt.Println("Initial List")
	fmt.Println(products)

	fmt.Println("Default Sort")
	products.Sort("Id")
	fmt.Println(products)

	fmt.Println("Sort by name")
	products.Sort("Name")
	fmt.Println(products)

}

package main

import "fmt"

type Product struct {
	Id       int
	Name     string
	Cost     float32
	Units    int
	Category string
}

func (product *Product) Format() string {
	return fmt.Sprintf("Id = %d, Name = %s, Cost = %v, Units = %d, Category = %s\n", product.Id, product.Name, product.Cost, product.Units, product.Category)
}

func (product *Product) ApplyDiscount(discount int) {
	product.Cost = product.Cost - ((float32(discount) / 100) * product.Cost)
}

type PerishableProduct struct {
	Product
	Expiry string
}

func (pp *PerishableProduct) Format() string {
	//return fmt.Sprintf("Id = %d, Name = %s, Cost = %v, Units = %d, Category = %s, Expiry = %s\n", pp.Id, pp.Name, pp.Cost, pp.Units, pp.Category, pp.Expiry)
	//using the Product.Format()
	return fmt.Sprintf("%s, Expiry = %s\n", pp.Product.Format(), pp.Expiry)
}

//products

type Products []Product

func (products *Products) PrintList() {
	for _, p := range *products {
		fmt.Print(p.Format())
	}
}

func (products *Products) IndexOf(product Product) int {
	for idx, p := range *products {
		if p == product {
			return idx
		}
	}
	return -1
}

func (products *Products) Includes(product Product) bool {
	return products.IndexOf(product) != -1
}

func (products *Products) Any(criteria func(Product) bool) bool {
	for _, product := range *products {
		if criteria(product) {
			return true
		}
	}
	return false
}

func (products *Products) All(criteria func(Product) bool) bool {
	for _, product := range *products {
		if !criteria(product) {
			return false
		}
	}
	return true
}

func (products *Products) Filter(criteria func(Product) bool) *Products {
	result := &Products{}
	for _, product := range *products {
		if criteria(product) {
			*result = append(*result, product)
		}
	}
	return result
}
func main() {
	pencil := Product{107, "Pencil", 200, 100, "Stationary"}
	fmt.Println(pencil.Format())
	pencil.ApplyDiscount(10)
	fmt.Println(pencil.Format())

	products := Products{
		Product{105, "Pen", 5, 50, "Stationary"},
		Product{107, "Pencil", 2, 100, "Stationary"},
		Product{103, "Marker", 50, 20, "Utencil"},
		Product{102, "Stove", 5000, 5, "Utencil"},
		Product{101, "Kettle", 2500, 10, "Stationary"},
		Product{104, "Scribble Pad", 20, 20, "Stationary"},
	}

	dummyProduct := Product{Id: -100}
	fmt.Println("Index of Pencil => ", products.IndexOf(pencil))
	fmt.Println("Includes Pencil ? => ", products.Includes(pencil))
	fmt.Println("Includes dummyProduct ? => ", products.Includes(dummyProduct))

	areThereCostlyProducts := products.Any(func(p Product) bool {
		return p.Cost > 50
	})
	fmt.Println("Are there any costly products ? => ", areThereCostlyProducts)

	stationaryProductCriteria := func(p Product) bool {
		return p.Category == "Stationary"
	}

	areThereStationaryProducts := products.Any(stationaryProductCriteria)
	fmt.Println("Are there any stationary products ? => ", areThereStationaryProducts)

	areThereElectronicsProducts := products.Any(func(p Product) bool {
		return p.Category == "Electronics"
	})
	fmt.Println("Are there any electronic products ? => ", areThereElectronicsProducts)

	areAllStationaryProducts := products.All(stationaryProductCriteria)

	fmt.Println("Are all the products stationary products => ", areAllStationaryProducts)

	stationaryProducts := products.Filter(stationaryProductCriteria)
	fmt.Println("Stationary Products")
	stationaryProducts.PrintList()

	fmt.Println()
	fmt.Println("Perishable Product")
	grapes := PerishableProduct{Product{150, "Grapes", 70, 50, "Food"}, "2 Days"}
	//Inheriting the methods from the composed type (Product)
	fmt.Print(grapes.Format())
	grapes.ApplyDiscount(10)
	fmt.Println("After applying discount")
	fmt.Print(grapes.Format())
}

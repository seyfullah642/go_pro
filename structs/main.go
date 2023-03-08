package main

import (
	"fmt"
	//"encoding/json"
	//"strings"
)

// Structs can be defined outside of the main function
// Using Pointer types for struct fields
type Product struct {
	name, category string
	price float64
	*Supplier
	//otherNames []string
}

type Supplier struct {
	name, city string
}

func main() {

	// Structs are custom data types which allow you to use all types
	// type Product struct {
	// 	name, category string
	// 	price float64
	// 	//otherNames []string
	// }

	type Item struct {
		category, name string
		//category string
		price float64
	}
	// Embedded field
	type StockLevel struct {
		Product
		Alternate Product
		count int
	}
	// kayak := Product {
	// 	name: "Kayak",
	// 	category: "Watersports",
	// 	price: 275,
	// }

	// kayak := Product {
	// 	name: "Kayak",
	// 	category: "Watersports",
	// }

	// fmt.Println(kayak.name, kayak.category, kayak.price)
	// kayak.price = 300
	// fmt.Println("Changed price:", kayak.price)

	// var lifejacket Product
	// fmt.Println("Name is zero value:", lifejacket.name == "")
	// fmt.Println("Category is zero value:", lifejacket.category == "")
	// fmt.Println("Price is zero value:", lifejacket.price == 0)

	// var kayak = Product { "Kayak", "Watersports", 275.00 }
	// fmt.Println("Name:", kayak.name)
	// fmt.Println("Category:", kayak.category)
	// fmt.Println("Price:", kayak.price)

	// stockItem := StockLevel {
	// 	Product: Product{ "Kayak", "Watersports", 275.00 },
	// 	Other: Product{ "Lifejacket", "Watersports", 48.95 },
	// 	count: 100,
	// }
	// fmt.Println("Name:", stockItem.Product.name)
	// fmt.Println("Alt Name:", stockItem.Other.name)
	// fmt.Println("Count:", stockItem.count)

	// Comparing structs
	// p1 := Product{ "Kayak", "Watersports", 275.00 }
	// p2 := Product{ "Kayak", "Watersports", 275.00 }
	// p3 := Product{ "Kayak", "Boats", 275.00 }

	// fmt.Println("p1 == p2:", p1 == p2)
	// fmt.Println("p1 == p3:", p1 == p3)

	// Converting between structs
	// prod := Product{ "Kayak", "Watersports", 275.00 }
	// item := Product{ "Kayak", "Watersports", 275.00 }
	// fmt.Println("prod == item:", prod == Product(item))

	// Anonymous structs
	// prod := Product{ "Kayak", "Watersports", 275.00 }
	// item := Product{ "Stadium", "Soccer", 75000 }
	// writeName(prod)
	// writeName(item)
	// var builder strings.Builder
	// json.NewEncoder(&builder).Encode(struct {
	// 	ProductName string
	// 	ProductPrice float64
	// }{
	// 	ProductName: prod.name,
	// 	ProductPrice: prod.price,
	// })
	// fmt.Println(builder.String())

	// Creating Arrays, Slices, and Maps containing struct values
	// array := [1]StockLevel{
	// 	{
	// 		Product: Product{"Kayak", "Watersports", 275.00},
	// 		Alternate: Product{"Lifejacket", "Watersports", 48.95},
	// 		count: 100,
	// 	},
	// }
	// fmt.Println("Array:", array[0].Product.name)

	// slice := []StockLevel{
	// 	{
	// 		Product: Product{"Kayak", "Watersports", 275.00},
	// 		Alternate: Product{"Lifejacket", "Watersports", 48.95},
	// 		count: 100,
	// 	},
	// 	{
	// 		Product: Product{"Boat", "Watersports", 1000.00},
	// 		Alternate: Product{"Lifejacket", "Watersports", 48.95},
	// 		count: 100,
	// 	},
	// }
	// fmt.Println("Slice:", slice)

	// kvp := map[string]StockLevel {
	// 	"kayak": {
	// 		Product: Product{"Kayak", "Watersports", 275.00},
	// 		Alternate: Product{"Lifejacket", "Watersports", 48.95},
	// 		count: 100,
	// 	},
	// 	"boat": {
	// 		Product: Product{"Boat", "Watersports", 1000.00},
	// 		Alternate: Product{"Lifejacket", "Watersports", 48.95},
	// 		count: 100,
	// 	},
	// }
	// fmt.Println("Map:", kvp)

	// Structs and Pointers
	// p1 := Product {
	// 	name: "Kayak",
	// 	category: "Watersports",
	// 	price: 275,
	// }
	// // p2 := p1 // Creates an entire new value (struct) and copies the field values
	// p2 := &p1 // memory address is assigned as value
	// p1.name = "Original Kayak"
	// fmt.Println("P1:", p1.name)
	// fmt.Println("P2:", (*p2).name) // * is a Pointer to dereference to access value at the memory address

	// kayak := &Product {
	// 	name: "Kayak",
	// 	category: "Watersports",
	// 	price: 275,
	// }
	//calcTax(&kayak)
	// calcTax(kayak)
	// fmt.Println("Name:", kayak.name, "Category:", kayak.category, "Price:", kayak.price)

	// kayak := calcTax(&Product {
	// 	name: "Kayak",
	// 	category: "Watersports",
	// 	price: 275,
	// })
	// fmt.Println("Name:", kayak.name, "Category:", kayak.category, "Price:", kayak.price)

	// Struct constructor functions
	// products := [2]*Product {
	// 	newProduct("Kayak", "Watersports", 275),
	// 	newProduct("Hat", "Skiing", 42.50),
	// }
	// for _, p := range products {
	// 	fmt.Println("Name:", p.name, "Category:", p.category, "Price:", p.price)
	// }
	
	// Using Pointer types for structs fields
	// acme := &Supplier {"Acme Co", "New York"}
	// products := [2]*Product {
	// 	newProduct("Kayak", "Watersports", 275, acme),
	// 	newProduct("Hat", "Skiing", 42.50, acme),
	// }
	// for _, p := range products {
	// 	fmt.Println("Name:", p.name, "Supplier:", p.Supplier.name, p.Supplier.city)
	// }

	// Pointer field copying
	acme := &Supplier {"Acme Co", "New York"}
	p1 := newProduct("Kayak", "Watersports", 275, acme)
	p2 := copyProduct(p1)
	p1.name = "Original Kayak"
	p1.Supplier.name = "BoatCo"
	for _, p := range []Product {*p1, p2} {
		fmt.Println("Name:", p.name, "Supplier:", p.Supplier.name, p.Supplier.city)
	}

	// Zero value for structs and pointers to structs
	//var prod Product
	// var prod Product = Product{ Supplier: &Supplier{}}
	// var prodPtr *Product
	// fmt.Println("Value:", prod.name, prod.category, prod.price, prod.Supplier.name)
	// fmt.Println("Pointer:", prodPtr)
}

// Pointer field copying

func copyProduct(product *Product) Product {
	p := *product
	s := *product.Supplier
	p.Supplier = &s
	return p
}

// Using Pointer types for structs fields

func newProduct(name, category string, price float64, supplier *Supplier) *Product {
	return &Product{name, category, price, supplier}
}

// Struct constructor functions

// func newProduct(name, category string, price float64) *Product {
// 	return &Product{name, category, price}
// }

// Structs and Pointers

// func calcTax(product *Product){
// 	if (*product).price > 100 {
// 		(*product).price += (*product).price * 0.2
// 	}
// }

// func calcTax(product *Product){
// 	if product.price > 100 {
// 		product.price += product.price * 0.2
// 	}
// }

// func calcTax(product *Product) *Product {
// 	if product.price > 100 {
// 		product.price += product.price * 0.2
// 	}
// 	return product
// }

// Anonymous structs

// func writeName(st struct{
// 	name, category string
// 	price float64
// }) {
// 	fmt.Println("Name:", st.name)
// }
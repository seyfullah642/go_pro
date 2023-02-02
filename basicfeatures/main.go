package main

import (
	"fmt"
	//"math/rand"
	"sort"
)

func main() {
	//fmt.Println("Value:", rand.Int())
	// fmt.Println("Hello, Go")
	// fmt.Println(20 + 20)
	// fmt.Println(20 + 30)
	//const price, tax float32 = 275.00, 27.50
	//const tax float32 = 27.50
	//const quantity, inStock = 2, true
	//const quantity = 2
	// fmt.Println("Total:", 2 * quantity * (price + tax))
	// fmt.Println("In stock: ", inStock)
	// var price float32 = 275.00
	// var tax float32 = 27.50
	// fmt.Println(price + tax)
	// price = 300
	// fmt.Println(price + tax)
	// var price = 275.00
	// var price2 = price
	// fmt.Println(price)
	// fmt.Println(price2)
	// var tax float32 = 27.50
	// fmt.Println(price + tax)
	// var price float32
	// fmt.Println(price)
	// price := 275.00
	// fmt.Println(price)
	// price, tax, inStock := 275.00, 27.50, true
	// fmt.Println("Total:", price + tax)
	// fmt.Println("In stock:", inStock)
	// price, tax, inStock := 275.00, 27.50, true
	// fmt.Println("Total:", price + tax)
	// fmt.Println("In stock:", inStock)
	//var price2, tax = 200.00, 25.00
	// price2, tax := 200.00, 25.00
	// fmt.Println("Total 2:", price2 + tax)

	// Pointers in Go

	// '&' symbol points to the address of the stored value

	// '*' declare a pointer and to dereference a pointer. 
	//	By derefrencing you get the value which the pointer is pointing to

	// first := 100
	// second := first
	// first++
	// fmt.Println("First:", first)
	// fmt.Println("Second:", second)
	// first := 100
	// var second *int = &first
	// first++
	// fmt.Println("First:", first)
	// fmt.Println("Second:", second)
	// first := 100
	// second := &first
	// first++
	// fmt.Println("First:", first)
	// fmt.Println("Second:", *second)
	// first := 100
	// second := &first
	// first++
	// *second++
	// var myNewPointer *int
	// myNewPointer = second
	// *myNewPointer++
	// fmt.Println("First:", first)
	// fmt.Println("Second:", *second)
	// first := 100
	// var second *int
	// fmt.Println(second)
	// second = &first
	// fmt.Println(second)
	// first := 100
	// second := &first
	// third := &second
	// fmt.Println(first)
	// fmt.Println(*second)
	// fmt.Println(**third)
	names := [3]string {"Alice", "Charlie", "Bob"}
	secondName := &names[1]
	fmt.Println(*secondName)
	sort.Strings(names[:])
	fmt.Println(*secondName)
}
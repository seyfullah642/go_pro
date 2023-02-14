package main

import (
	"fmt"
)

func main() {
	// fmt.Println("About to call function")
	// printPrice()
	// fmt.Println("Function complete")

	// printPrice("Kayak", 275, 0.2)
	// printPrice("Lifejacket", 48.95, 0.2)
	// printPrice("Soccer Ball", 19.50, 0.15)

	// Variadic Parameters
	
	// printSuppliers("Kayak", []string {"Acme Kayaks", "Bob's Boats", "Crazy Canoes"})
	// printSuppliers("Lifejacket", []string {"Sail Safe Co"})

	// printSuppliers("Kayak", "Acme Kayaks", "Bob's Boats", "Crazy Canoes")
	// printSuppliers("Lifejacket", "Sail Safe Co")
	// printSuppliers("Soccer Ball")

	// names := []string {"Acme Kayaks", "Bob's Boats", "Crazy Canoes"}
	// printSuppliers("Kayak:", names...)
	// printSuppliers("Lifejacket", "Sail Safe Co")
	// printSuppliers("Soccer Ball")

	// Pointers as Function Parameters

	// val1, val2 := 10, 20
	// fmt.Println("Before calling function", val1, val2)
	// swapValues(&val1, &val2)
	// fmt.Println("After calling function", val1, val2)

	// products := map[string] float64 {
	// 	"Kayak": 275,
	// 	"Lifejacket": 48.95,
	// }
	// for product, price := range products {
	// 	priceWithTax := calcTax(price)
	// 	fmt.Println("Product:", product, ", Price:", priceWithTax)
	// }

	// for product, price := range products {
	// 	//priceWithTax := calcTax(price)
	// 	fmt.Println("Product:", product, ", Price:", calcTax(price))
	// }

	// val1, val2 := 10, 20
	// fmt.Println("Before calling function", val1, val2)
	// val1, val2 = swapValues(val1, val2)
	// fmt.Println("After calling function", val1, val2)

	products := map[string]float64 {
		"Kayak": 275,
		"Lifejacket": 48.95,
	}
	// for product, price := range products {
	// 	tax := calcTax(price)
	// 	if tax != -1 {
	// 		fmt.Println("Product:", product, "Tax:", tax)
	// 	} else {
	// 		fmt.Println("Product:", product, "No tax due")
	// 	}
	// }

	// for product, price := range products {
	// 	taxAmount, taxDue := calcTax(price)
	// 	if taxDue {
	// 		fmt.Println("Product:", product, "Tax:", taxAmount)
	// 	} else {
	// 		fmt.Println("Product:", product, "No tax due")
	// 	}
	// }

	// total1, tax1 := calcTotalPrice(products, 10)
	// fmt.Println("Total 1:", total1, "Tax 1:", tax1)
	// total2, tax2 := calcTotalPrice(nil, 10)
	// fmt.Println("Total 2:", total2, "Tax 2:", tax2)

	_, total := calcTotalPrice(products)
	fmt.Println("Total:", total)

}

// Defer keyword

func calcTotalPrice(products map[string]float64) (count int, total float64) {
	fmt.Println("Function started")
	defer fmt.Println("First defer call")
	count = len(products)
	for _, price := range products {
		total += price
	}
	defer fmt.Println("Second defer call")
	fmt.Println("Function about to return")
	return
}

// func calcTotalPrice(products map[string]float64) (count int, total float64) {
// 	count = len(products)
// 	for _, price := range products {
// 		total += price
// 	}
// 	return
// }

// func calcTax(price float64) (float64, bool) {
// 	if price > 100 {
// 		return price * 0.2, true
// 	}
// 	return 0, false
// }

// func calcTotalPrice(products map[string]float64, minSpend float64) (total, tax float64) {
// 	total = minSpend
// 	for _, price := range products {
// 		if taxAmount, due := calcTax(price); due {
// 			total += taxAmount
// 			tax += taxAmount
// 		} else {
// 			total += price
// 		}
// 	}
// 	return
// }

// func calcTax(price float64) float64 {
// 	if price > 100 {
// 		return price * 0.2
// 	}
// 	return -1
// }

// func swapValues(first, second int) (int, int) {
// 	return second, first
// }

// func calcTax(price float64) float64 {
// 	return price + (price * 0.2)
// }

// Pointers as Function Parameters

// func swapValues(first, second *int) {
// 	fmt.Println("Before swap:", *first, *second)
// 	temp := *first
// 	*first = *second
// 	*second = temp
// 	fmt.Println("After swap:", *first, *second)
// }

// func swapValues(first, second int) {
// 	fmt.Println("Before swap:", first, second)
// 	temp := first
// 	first = second
// 	second = temp
// 	fmt.Println("After swap:", first, second)
// }

// Variadic Parameters

// func printSuppliers(product string, suppliers ...string) {
// 	if len(suppliers) == 0 {
// 		fmt.Println("Product:", product, "Supplier: (none)")
// 	} else {
// 		for _, supplier := range suppliers {
// 			fmt.Println("Product:", product, "Supplier:", supplier)
// 		}
// 	}
// }

// func printSuppliers(product string, supplies ...string) {
// 	for _, supplier := range supplies {
// 		fmt.Println("Product:", product, "Supplier:", supplier)
// 	}
// }

// func printSuppliers(product string, supplies []string) {
// 	for _, supplier := range supplies {
// 		fmt.Println("Product:", product, "Supplier:", supplier)
// 	}
// }

// func printPrice(string, float64, float64) {
// 	//taxAmount := price * taxRate
// 	fmt.Println("No parameters")
// }

// Functions with/without parameters

// func printPrice(product string, price float64, _ float64) {
// 	taxAmount := price * 0.25
// 	fmt.Println(product, "Price:", price, "Tax:", taxAmount)
// }

// func printPrice(product string, price, taxRate float64) {
// 	taxAmount := price * taxRate
// 	fmt.Println(product, "Price:", price, "Tax:", taxAmount)
// }

// func printPrice(product string, price float64, taxRate float64) {
// 	taxAmount := price * taxRate
// 	fmt.Println(product, "Price:", price, "Tax:", taxAmount)
// }

// func printPrice() {
// 	kayakPrice := 275.00
// 	kayakTax := kayakPrice * 0.2
// 	fmt.Println("Price:", kayakPrice, "Tax:", kayakTax)
// }
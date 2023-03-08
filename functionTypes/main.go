package main

import (
	"fmt"
)

// Creating function type aliases
type calcFunc func(float64) float64

func main() {
	// products := map[string]float64 {
	// 	"Kayak": 275,
	// 	"Lifejacket": 48.95,
	// }
	
	// Functions

	// for product, price := range products {
	// 	// The line below is a function type
	// 	var calcFunc func(float64) float64
	// 	fmt.Println("Function assigned:", calcFunc == nil)
	// 	if (price > 100) {
	// 		calcFunc = calcWithTax
	// 	} else {
	// 		calcFunc = calcWithoutTax
	// 	}
	// 	fmt.Println("Function assigned:", calcFunc == nil)
	// 	totalPrice := calcFunc(price)
	// 	fmt.Println("Product:", product, "Price:", totalPrice)
	// }

	// Functions as arguments

	// for product, price := range products {
	// 	if (price > 100) {
	// 		printPrice(product, price, calcWithTax)
	// 	} else {
	// 		printPrice(product, price, calcWithoutTax)
	// 	}
	// }

	// Functions as results

	// for product, price := range products {
	// 	printPrice(product, price, selectCalculator(price))
	// }
	
	// Order of func calls:
	// printPrice function is called.
	// With selectCalculator as a func arg, that is then called.
	// Since selectCalculator func returns a function, 
	//	calcWithTax or calcWithoutTax is called and the price value is ultimately returned.

	// Function closure

	watersportsProducts := map[string]float64 {
		"Kayak": 275,
		"Lifejacket": 48.95,
	}
	soccerProducts := map[string]float64 {
		"Soccer Ball": 19.50,
		"Stadium": 79500,
	}
	waterCalc := priceCalcFactory(100, 0.2)
	// calc := func(price float64) float64 {
	// 	if price > 100 {
	// 		return price + (price * 0.2)
	// 	}
	// 	return price
	// }
	for product, price := range watersportsProducts {
		printPrice(product, price, waterCalc)
	}
	soccerCalc := priceCalcFactory(50, 0.1)
	// calc = func(price float64) float64 {
	// 	if price > 50 {
	// 		return price + (price * 0.1)
	// 	}
	// 	return price
	// }
	for product, price := range soccerProducts {
		printPrice(product, price, soccerCalc)
	}

}

// Functions

// func calcWithTax(price float64) float64 {
// 	return price + (price * 0.2)
// }
// func calcWithoutTax(price float64) float64 {
// 	return price
// }

// Functions as arguments

// func printPrice(product string, price float64, calculator func(float64) float64) {
// 	fmt.Println("Product:", product, "Price:", calculator(price))
// }

// Functions as results

// func selectCalculator(price float64) func(float64) float64 {
// 	if price > 100 {
// 		return calcWithTax
// 	}
// 	return calcWithoutTax
// }

// Function type aliases

// func printPrice(product string, price float64, calculator calcFunc) {
// 	fmt.Println("Product:", product, "Price:", calculator(price))
// }

// func selectCalculator(price float64) calcFunc {
// 	if price > 100 {
// 		return calcWithTax
// 	}
// 	return calcWithoutTax
// }

// Literal function syntax aka anonymous functions

// func selectCalculator(price float64) calcFunc {
// 	if price > 100 {
// 		withTax := func (price float64) float64 {
// 			return price + (price * 0.2)
// 		}
// 		return withTax
// 	}
// 	withoutTax := func (price float64) float64 {
// 		return price
// 	}
// 	return withoutTax
// }

// Using functions values directly

// func selectCalculator(price float64) calcFunc {
// 	if price > 100 {
// 		return func (price float64) float64 {
// 			return price + (price * 0.2)
// 		}
// 	}
// 	return func (price float64) float64 {
// 		return price
// 	}
// }

// Function closures

func printPrice(product string, price float64, calculator calcFunc) {
	fmt.Println("Product:", product, "Price:", calculator(price))
}

func priceCalcFactory(threshold, rate float64) calcFunc {
	return func(price float64) float64 {
		if price > threshold {
			return price + (price * rate)
		}
		return price
	}
}
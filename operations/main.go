package main

import (
	"fmt"
	"math"
)

func main() {
	// price, tax := 275.00, 27.40
	// sum := price + tax
	// difference := price - tax
	// product := price * tax
	// quotient := price / tax
	// fmt.Println(sum)
	// fmt.Println(difference)
	// fmt.Println(product)
	// fmt.Println(quotient)

	var intVal = math.MaxInt64
	var floatVal = math.MaxFloat64
	fmt.Println(intVal * 2)
	fmt.Println(floatVal * 2)
	fmt.Println(math.IsInf((floatVal * 2), 0))
}
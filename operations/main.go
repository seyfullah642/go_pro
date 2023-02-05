package main

import (
	"fmt"
	//"math"
	"strconv"
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

	// var intVal = math.MaxInt64
	// var floatVal = math.MaxFloat64
	// fmt.Println(intVal * 2)
	// fmt.Println(floatVal * 2)
	// fmt.Println(math.IsInf((floatVal * 2), 0))

	// val1 := "true"
	// val2 := "false"
	// val3 := "not true"
	// bool1, b1err := strconv.ParseBool(val1)
	// bool2, b2err := strconv.ParseBool(val2)
	// bool3, b3err := strconv.ParseBool(val3)
	// fmt.Println("Bool 1", bool1, b1err)
	// fmt.Println("Bool 2", bool2, b2err)
	// fmt.Println("Bool 3", bool3, b3err)

	// val1 := "0"
	// bool1, b1err := strconv.ParseBool(val1)
	// if b1err == nil {
	// 	fmt.Println("Parsed value:", bool1)
	// } else {
	// 	fmt.Println("Cannot parse", val1)
	// }

	// val1 := "0"
	// if bool1, b1err := strconv.ParseBool(val1); b1err == nil {
	// 	fmt.Println("Parsed value:", bool1)
	// } else {
	// 	fmt.Println("Cannot parse", val1)
	// }

	// val1 := "100"
	// int1, int1err := strconv.ParseInt(val1, 0, 8)
	// if int1err == nil {
	// 	fmt.Println("Parsed value:", int1)
	// } else {
	// 	fmt.Println("Cannot parse", val1)
	// }

	// val1 := "500"
	// int1, int1err := strconv.ParseInt(val1, 0, 16)
	// if int1err == nil {
	// 	fmt.Println("Parsed value:", int1)
	// } else {
	// 	fmt.Println("Cannot parse", val1, int1err)
	// }

	// val1 := "100"
	// int1, int1err := strconv.ParseInt(val1, 10, 0)
	// if int1err == nil {
	// 	var intResult int = int(int1)
	// 	fmt.Println("Parsed value:", intResult)
	// } else {
	// 	fmt.Println("Cannot parse", val1, int1err)
	// }

	val1 := "100"
	int1, int1err := strconv.Atoi(val1)
	if int1err == nil {
		var intResult int = int1
		fmt.Println("Parsed value:", intResult)
	} else {
		fmt.Println("Cannot parse", val1, int1err)
	}
}
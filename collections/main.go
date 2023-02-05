package main

import (
	"fmt"
)

func main() {
	// var names [3]string
	// names[0] = "Kayak"
	// names[1] = "Lifejacket"
	// names[2] = "Paddle"
	// fmt.Println(names)

	// names := [3]string {"Kayak", "Lifejacket", "Paddle"}
	// fmt.Println(names)

	// names := [3]string {"Kayak", "Lifejacket", "Paddle"}
	// var otherArray [4]string = names
	// fmt.Println(names)

	// names := [3]string {"Kayak", "Lifejacket", "Paddle"}
	// otherArray := names
	// names[0] = "Canoe"
	// fmt.Println("names:", names)
	// fmt.Println("otherArray:", otherArray)

	// names := [3]string {"Kayak", "Lifejacket", "Paddle"}
	// otherArray := &names
	// names[0] = "Canoe"
	// fmt.Println("names:", names)
	// fmt.Println("otherArray:", *otherArray)

	// names := [3]string {"Kayak", "Lifejacket", "Paddle"}
	// moreNames := [3]string {"Kayak", "Lifejacket", "Paddle"}
	// same := names == moreNames
	// fmt.Println("comparision:", same)

	// names := [3]string {"Kayak", "Lifejacket", "Paddle"}
	// for index, value := range names {
	// 	fmt.Println("Index:", index, "Value:", value)
	// }

	// SLICIES

	// names := make([]string, 3)
	// names[0] = "Kayak"
	// names[1] = "Lifejacket"
	// names[2] = "Paddle"
	// fmt.Println(names)

	// names := []string {"Kayak", "Lifejacket", "Paddle"}
	// fmt.Println(names)

	// names := []string {"Kayak", "Lifejacket", "Paddle"}
	// names = append(names, "Hat", "Gloves")
	// fmt.Println(names)

	names := []string {"Kayak", "Lifejacket", "Paddle"}
	appendedNames := append(names, "Hat", "Gloves")
	names[0] = "Canoe"
	
}
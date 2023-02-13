package main

import (
	"fmt"
	//"sort"
	//"reflect"
	//"strconv"
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

	// names := []string {"Kayak", "Lifejacket", "Paddle"}
	// appendedNames := append(names, "Hat", "Gloves")
	// names[0] = "Canoe"
	// fmt.Println("names:", names)
	// fmt.Println("appendedNames:", appendedNames)

	// names := make([]string, 3, 6)
	// names[0] = "Kayak"
	// names[1] = "Lifejacket"
	// names[2] = "Paddle"
	// fmt.Println("len:", len(names))
	// fmt.Println("cap:", cap(names))

	// names := make([]string, 3, 6)
	// names[0] = "Kayak"
	// names[1] = "Lifejacket"
	// names[2] = "Paddle"
	// appendedNames := append(names, "Hat", "Glove")
	// names[0] = "Canoe"
	// fmt.Println("names:", names)
	// fmt.Println("appendedNames:", appendedNames)

	// names := make([]string, 3, 6)
	// names[0] = "Kayak"
	// names[1] = "Lifejacket"
	// names[2] = "Paddle"
	// moreNames := []string {"Hat Gloves"}
	// appendedNames := append(names, moreNames...)
	// fmt.Println(appendedNames)

	// products := [4]string {"Kayak", "Lifejacket", "Paddle", "Hat"}
	// //someNames := products[1:3]
	// someNames := products[1:3:3] // This form allows you to specify the capacity of the slice
	// allNames := products[:]
	// someNames = append(someNames, "Gloves")
	// //someNames = append(someNames, "Boots")
	// fmt.Println("someNames:", someNames)
	// fmt.Println("len:", len(someNames), "cap:", cap(someNames))
	// fmt.Println("allNames:", allNames)
	// fmt.Println("len:", len(allNames), "cap:", cap(allNames))

	// products := [4]string {"Kayak", "Lifejacket", "Paddle", "Hat"}
	// allNames := products[1:]
	// someNames := allNames[1:3]
	// allNames = append(allNames, "Glove")
	// allNames[0] = "Canoe"
	// fmt.Println("someNames:", someNames)
	// fmt.Println("allNames:", allNames)

	// products := [4]string {"Kayak", "Lifejacket", "Paddle", "Hat"}
	// allNames := products[1:]
	// //someNames := make([]string, 2)
	// var someNames []string
	// copy(someNames, allNames)
	// fmt.Println("someNames:", someNames)
	// fmt.Println("allNames:", allNames)

	// products := [4]string {"Kayak", "Lifejacket", "Paddle", "Hat"}
	// allNames := products[1:]
	// someNames := []string {"Boots", "Gloves"}
	// copy(someNames[1:], allNames[2:3])
	// fmt.Println("someNames:", someNames)
	// fmt.Println("allNames:", allNames)

	// products := []string {"Kayak", "Lifejacket", "Paddle", "Hat"}
	// replacementProducts := []string {"Canoe", "Boots"}
	// //copy(products, replacementProducts)
	// copy(products[0:1], replacementProducts)
	// fmt.Println("products:", products)

	// products := [4]string {"Kayak", "Lifejacket", "Paddle", "Hat"}
	// delete := append(products[:2], products[3:]...)
	// fmt.Println("Deleted:", delete)

	// products := [4]string {"Kayak", "Lifejacket", "Paddle", "Hat"}
	// sort.Strings(products[:])
	// for index, value := range products {
	// 	fmt.Println("Index:", index, "Value:", value)
	// }

	// p1 := []string {"Kayak", "Lifejacket", "Paddle", "Hat"}
	// p2 := p1
	// //fmt.Println("Equal:", p1 == p2)
	// fmt.Println("Equal:", reflect.DeepEqual(p1,p2))

	// p1 := []string {"Kayak", "Lifejacket", "Paddle", "Hat"}
	// arrayPtr := (*[3]string)(p1)
	// array := *arrayPtr
	// fmt.Println(array)

	// MAPS

	// products := make(map[string]float64, 10)
	// products["Kayak"] = 279
	// products["Lifejacket"] = 48.95

	// products := map[string]float64 {
	// 	"Kayak": 279, 
	// 	"Lifejacket": 48.9,
	// 	"Hat": 0,
	// }
	// fmt.Println("Map size:", len(products))
	// fmt.Println("Price:", products["Kayak"])
	//fmt.Println("Price:", products["Hat"])

	// value, ok := products["Hat"]
	// if ok {
	// 	fmt.Println("Stored value:", value)
	// } else {
	// 	fmt.Println("No stored value")
	// }

	// if value, ok := products["Hat"]; ok {
	// 	fmt.Println("Stored value:", value)
	// } else {
	// 	fmt.Println("No stored value")
	// }

	// delete(products, "Hat")
	// if value, ok := products["Hat"]; ok {
	// 	fmt.Println("Stored value:", value)
	// } else {
	// 	fmt.Println("No stored value")
	// }

	// for key, value := range products {
	// 	fmt.Println("key:", key, "value:", value)
	// }

	// keys := make([]string, 0, len(products))
	// for key,_ := range products {
	// 	keys = append(keys, key)
	// }
	// sort.Strings(keys)
	// for _,key := range keys {
	// 	fmt.Println("Key:", key, "value:", products[key])
	// }

	// STRINGS
	
	// var price string = "$48.95"
	// //var currency byte = price[0]
	// var currency string = string(price[0])
	// var amountString string = price[1:]
	// amount, parseErr := strconv.ParseFloat(amountString, 64)
	// fmt.Println("Currency:", currency)
	// if parseErr == nil {
	// 	fmt.Println("Amount:", amount)
	// } else {
	// 	fmt.Println("Error:", parseErr)
	// }

	var price = "$48.95"
	for index, char := range price {
		fmt.Println(index, char, string(char))
	}

}
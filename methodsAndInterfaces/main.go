package main

import (
	"fmt"
)

// Interfaces
type Expense interface {
	getName() string
	getCost(annual bool) float64
}

type Person struct {
	name, city string
}

// Using an Interface for Struct fields
type Account struct {
	accountNumber int
	expenses []Expense
}

func main() {
	// kayak := Product{"Kayak", "Watersports", 275}
	// insurance := Service{"Boat Cover", 12, 89.50}
	// insPrice := insurance.monthlyFee * float64(insurance.durationMonths)

	// fmt.Println("Product:", kayak.name, "Price:", kayak.price)
	// fmt.Println("Service:", insurance.description, "Price:", insPrice)

	// Implement interface
	// expenses := []Expense {
	// 	Product{"Kayak", "Watersports", 275},
	// 	Service{"Boat Cover", 12, 89.50},
	// }
	// for _, expense := range expenses {
	// 	fmt.Println("Expense:", expense.getName(), "Cost:", expense.getCost(true))
	// }

	// Using Interface in a function
	// fmt.Println("Total:", calcTotal(expenses))

	// Using an Interface for Struct fields
	// account := Account{
	// 	accountNumber: 12345,
	// 	expenses: [] Expense{
	// 		Product{"Kayak", "Watersports", 275},
	// 		Service{"Boat Cover", 12, 89.50},
	// 	},
	// }

	// for _, expense := range account.expenses {
	// 	fmt.Println("Expense:", expense.getName(), "Cost:", expense.getCost(true))
	// }
	// fmt.Println("Total:", calcTotal(account.expenses))

	// Understanding effect of pointer method receivers
	// product := Product {"Kayak", "Watersports", 275}
	// var expense Expense = &product
	// var expense Expense = &product
	// product.price = 100
	// fmt.Println("Product field value:", product.price)
	// fmt.Println("Expense method result:", expense.getCost(false))

	// Comparing interface values
	// var e1 Expense = &Product{name: "Kayak"}
	// var e2 Expense = &Product{name: "Kayak"}
	// var e3 Expense = Service{description: "Boat Cover"}
	// var e4 Expense = Service{description: "Boat Cover"}
	// fmt.Println("e1 == e2", e1 == e2)
	// fmt.Println("e3 == e4", e3 == e4)

	// Performing type assertions
	// expenses := []Expense {
	// 	Service {"Boat Cover", 12, 89.50, []string{}},
	// 	Service {"Boat Cover", 12, 8, []string{}},
	// }

	// for _, expense := range expenses {
	// 	s := expense.(Service) // here is the type assertion
	// 	fmt.Println("Service:", s.description, "Price:", s.monthlyFee * float64(s.durationMonths))
	// }

	// Testing before performing a type assertion
	// expenses := []Expense {
	// 	Service {"Boat Cover", 12, 89.50, []string{}},
	// 	Service {"Boat Cover", 12, 8, []string{}},
	// 	&Product {"Kayak", "Watersports", 275},
	// }

	// for _, expense := range expenses {
	// 	if s, ok := expense.(Service); ok {
	// 		//s := expense.(Service) // here is the type assertion
	// 		fmt.Println("Service:", s.description, "Price:", s.monthlyFee * float64(s.durationMonths))
	// 	} 
	// 	if p, ok := expense.(*Product); ok {
	// 		fmt.Println("Expense:", p.getName(), "Cost:", p.getCost(false))
	// 	}
	// }

	// Switching on dynamic types
	// expenses := []Expense {
	// 	Service {"Boat Cover", 12, 89.50, []string{}},
	// 	Service {"Boat Cover", 12, 8, []string{}},
	// 	&Product {"Kayak", "Watersports", 275},
	// }
	// for _, expense := range expenses {
	// 	switch value := expense.(type) {
	// 		case Service:
	// 			fmt.Println("Service:", value.description, "Price:", value.monthlyFee * float64(value.durationMonths))
	// 		case *Product:
	// 			fmt.Println("Expense:", value.getName(), "Cost:", value.getCost(false))
	// 		default:
	// 			fmt.Println("Expense:", expense.getName(), "Cost:", expense.getCost(true))
	// 	}
	// }

	// Using an empty interface
	// Go allows use of empty interface, which means an interface that defines no methods, to represent any type
	// This is useful way to group disparate types that share no common features
	// interface{} -> empty interface declaration
	// var expense Expense = &Product{"Kayak", "Watersports", 275}
	// data := []interface{} {
	// 	expense,
	// 	Product{"Lifejacket", "Watersports", 48.95},
	// 	Service{"Boat Cover", 12, 89.50, []string{}},
	// 	Person{"Alice", "London"},
	// 	&Person{"Bob", "New York"},
	// 	"This is a string",
	// 	100,
	// 	true,
	// }
	// for _, item := range data {
	// 	switch value := item.(type) {
	// 		case Product:
	// 			fmt.Println("Product:", value.name, "Price:", value.price)
	// 		case *Product:
	// 			fmt.Println("Product Pointer:", value.name, "Price:", value.price)
	// 		case Service:
	// 			fmt.Println("Service:", value.description, "Price:", value.monthlyFee * float64(value.durationMonths))
	// 		case Person:
	// 			fmt.Println("Person:", value.name, "City:", value.city)
	// 		case *Person:
	// 			fmt.Println("Pointer Person:", value.name, "City:", value.city)
	// 		case string, bool, int:
	// 			fmt.Println("Built-in type:", value)
	// 		default:
	// 			fmt.Println("Default:", value)
	// 	}
	// }

	// Using empty interface for function parameters
	var expense Expense = &Product{"Kayak", "Watersports", 275}
	data := []interface{} {
		expense,
		Product{"Lifejacket", "Watersports", 48.95},
		Service{"Boat Cover", 12, 89.50, []string{}},
		Person{"Alice", "London"},
		&Person{"Bob", "New York"},
		"This is a string",
		100,
		true,
	}
	for _, item := range data {
		proccessItem(item)
	}
}

// Using empty interface for function parameters
func proccessItem(item interface{}){
	switch value := item.(type) {
	case Product:
		fmt.Println("Product:", value.name, "Price:", value.price)
	case *Product:
		fmt.Println("Product Pointer:", value.name, "Price:", value.price)
	case Service:
		fmt.Println("Service:", value.description, "Price:", value.monthlyFee * float64(value.durationMonths))
	case Person:
		fmt.Println("Person:", value.name, "City:", value.city)
	case *Person:
		fmt.Println("Pointer Person:", value.name, "City:", value.city)
	case string, bool, int:
		fmt.Println("Built-in type:", value)
	default:
		fmt.Println("Default:", value)
}
}

// Using Interface in a function
func calcTotal(expense []Expense) (total float64) {
	for _, item := range expense {
		total += item.getCost(true)
	}
	return total
}
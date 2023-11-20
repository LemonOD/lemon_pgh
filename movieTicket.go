package main

import "fmt"

func main() {
	age := 0
	fmt.Print("Enter the person's age: ")
	fmt.Scan(&age)

	ticketPrice := calculatePriceOfTicket(age)

	fmt.Println("The ticket price is: N", ticketPrice)
}

func calculatePriceOfTicket(age int) int {
	switch {
	case age <= 12:
		return 5
	case age > 12 && age <= 64:
		return 10
	case age >= 65:
		return 7
	default:
		return 0
	}
}

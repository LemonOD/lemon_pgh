package main

import "fmt"

func main() {
	age := 0
	fmt.Print("Enter the person's age: ")
	fmt.Scan(&age)

	ticket_price := calculatePriceOfTicket(age)

	fmt.Printf("The ticket price will be: N%d\n", ticket_price)
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

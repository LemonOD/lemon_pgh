package main

import (
	"fmt"
)

func main() {
	score := 0
	fmt.Print("Enter student's score: ")
	fmt.Scan(&score)

	grade := convertScoreToGrade(score)

	fmt.Println("The student's grade is: ", grade)
}

func convertScoreToGrade(score int) string {
	switch {
	case score >= 90 && score <= 100:
		return "A"
	case score >= 80 && score < 90:
		return "B"
	case score >= 70 && score < 80:
		return "C"
	case score >= 60 && score < 70:
		return "D"
	case score >= 0 && score < 60:
		return "F"
	default:
		return "enter a score between 0 and 100."
	}
}

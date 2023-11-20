package main

import (
	"fmt"
	"unicode"
)

func main() {
	var inputChar rune
	fmt.Print("Enter a single character: ")
	fmt.Scanf("%c", &inputChar)

	isVowel := checkVowel(inputChar)

	if isVowel {
		fmt.Printf("%c is a vowel.\n", inputChar)
	} else {
		fmt.Printf("%c is a consonant.\n", inputChar)
	}
}

func checkVowel(char rune) bool {
	char = unicode.ToLower(char)
	switch char {
	case 'a', 'e', 'i', 'o', 'u':
		return true
	default:
		return false
	}
}

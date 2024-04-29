package main

import (
	"fmt"
	"unicode"
)

func main() {
	var symbol rune
	fmt.Scan(&symbol)

	if unicode.IsUpper(symbol) {
		fmt.Println("UPPER")
	}
	if unicode.IsLower(symbol) {
		fmt.Println("lower")
	}
	if unicode.IsTitle(symbol) {
		fmt.Println("Title")
	}

	// put your conditions next
}

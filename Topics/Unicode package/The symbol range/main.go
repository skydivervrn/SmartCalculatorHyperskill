package main

import (
	"fmt"
	"unicode"
)

func main() {
	var symbol rune
	fmt.Scan(&symbol)

	switch {
	case unicode.In(symbol, unicode.Latin):
		fmt.Println("Latin")
	case unicode.In(symbol, unicode.Greek):
		fmt.Println("Greek")
		// write the cases for the `Egyptian_Hieroglyphs` and `other` symbols below:
	case unicode.In(symbol, unicode.Egyptian_Hieroglyphs):
		fmt.Println("Egyptian")
	default:
		fmt.Println("other")
	}

}

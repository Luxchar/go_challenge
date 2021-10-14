package main

import (
	"github.com/01-edu/z01"
)

func IsNegative(nb int) {
	if nb >= 0 {
		z01.PrintRune(rune(70))
	} else {
		z01.PrintRune(rune(84))

	}
	z01.PrintRune(rune('\n'))
}

/*
func main() {

	IsNegative(1)
	IsNegative(-1)
}
*/

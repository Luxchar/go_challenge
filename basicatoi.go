package main

import (
	"fmt"

	"github.com/01-edu/z01"
)

func BasicAtoi(s string) int {
	a := []rune(s)
	for b := range a {
		if b != 0 && 48+b <= 57 {
			z01.PrintRune(rune(48 + b))
		}
	}
	return 0
}

func main() {
	fmt.Println(BasicAtoi("12345"))
	fmt.Println(BasicAtoi("0000000012345"))
	fmt.Println(BasicAtoi("000000"))
}

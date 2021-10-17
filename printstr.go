package main

import (
	"github.com/01-edu/z01"
)

func PrintStr(s string) {
	var char string
	for i := 0; i < len(char); i++ {
		runes := []rune(char)

		var r []int

		for i := 0; i < len(runes); i++ {
			z01.PrintRune(int(runes[i]))

		}
	}
}
func main() {
	PrintStr("Hello World!")
}

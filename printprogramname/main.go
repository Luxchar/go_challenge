package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	s := os.Args[0]
	for _, r := range s {
		z01.PrintRune(r)
	}
	z01.PrintRune(rune('\n'))
}

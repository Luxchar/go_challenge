package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	s := ""
	a := os.Args[1:]
	for _, arg := range a {
		s += arg
		for _, r := range s {
			z01.PrintRune(r)
		}
		s = ""
		z01.PrintRune('\n')
	}
}

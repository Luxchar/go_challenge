package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	arg := os.Args
	for i := len(arg) - 1; i >= 1; i-- {
		for j := len(arg) - 1; j >= 1; j-- {
			if arg[int(i)] > arg[int(j)] {
				arg[i], arg[j] = arg[j], arg[i]
			}
		}
	}
	for i := 1; i < len(arg); i++ {
		for _, k := range arg[i] {
			z01.PrintRune(k)
		}
		z01.PrintRune('\n')
	}
}

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
				print(arg[i])
				arg[i], arg[j] = arg[j], arg[i]
			}
		}
	}
	for _, k := range arg {
		print(k)
		z01.PrintRune('\n')
	}
}

/*
func main() {
	arg := os.Args
	j := rune(0)
	for i := len(arg) - 1; i >= 1; i-- {
		for i := len(arg) - 1; i >= 1; i-- {
			if j > i {
				break
			}
			if arg[int(i)] > arg[int(j)] {
				arg[j] = arg[i]
			}
			j += 1
		}
		for _, j := range arg[i] {
			z01.PrintRune(j)
			z01.PrintRune('\n')
		}
		j += 1
	}
}
*/

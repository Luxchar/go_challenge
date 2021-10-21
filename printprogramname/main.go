package main

import (
	"os"

	"github.com/01-edu/z01"
)

func osname(s string) {
	for _, r := range s {
		z01.PrintRune(r)
	}
	z01.PrintRune(rune('\n'))
}
func main() {
	osname(os.Args[0])

}

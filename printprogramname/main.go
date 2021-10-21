package main

import (
	"os"
	"path/filepath"

	"github.com/01-edu/z01"
)

func main() {
	path := os.Args[0]
	path = filepath.Base(path)
	for _, r := range path {
		z01.PrintRune(r)
	}

	z01.PrintRune(rune('\n'))

}

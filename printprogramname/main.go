package main

import (
	"os"
	"path/filepath"

	"github.com/01-edu/z01"
)

func osname(s string) {
	for _, r := range s {
		z01.PrintRune(r)
	}
	z01.PrintRune(rune('\n'))
}
func main() {
	s := ""
	path := os.Args[0]
	s = filepath.Base(path)
	osname(s)

}

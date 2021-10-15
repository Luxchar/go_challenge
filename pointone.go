package piscine

import (
	"github.com/01-edu/z01"
)

func PointOne(n *int) {
	var a, b int
	a = 1
	p := &a
	b = *p
	z01.PrintRune(rune(48 + b))
	z01.PrintRune(rune('\n'))
}

/*
func main() {
	n := 0
	PointOne(&n)
}
*/

package main

import (
	"github.com/01-edu/z01"
)

func PrintNbr(n int) {
	order := []int{}
	if n == 0 {
		z01.PrintRune(rune('0'))
	} else {
		for n > 0 {
			val := n % 10
			order = append(order, val)
			n /= 10
		}
		for i := 0; i < len(order); i++ {
			for j := 0; j < len(order); j++ {
				if order[i] > order[j] {
					order[i], order[j] = order[j], order[i]
				}
			}
		}
		for i := 0; i < len(order); i++ {
			z01.PrintRune(rune(48 + order[i]))
		}
	}
}

type point struct {
	x int
	y int
}

func setPoint(ptr *point) {
	ptr.x = 42
	ptr.y = 21
}

func main() {
	points := &point{}

	setPoint(points)

	z01.PrintRune(rune('x'))
	z01.PrintRune(rune(' '))
	z01.PrintRune(rune('='))
	z01.PrintRune(rune(' '))
	PrintNbr(points.x)
	z01.PrintRune(rune(' '))
	z01.PrintRune(rune('y'))
	z01.PrintRune(rune(' '))
	z01.PrintRune(rune('='))
	z01.PrintRune(rune(' '))
	PrintNbr(points.y)
}

package piscine

import "github.com/01-edu/z01"

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
			z01.PrintRune(rune(48 + order[i]))
		}
	}
}

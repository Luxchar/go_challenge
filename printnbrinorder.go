package piscine

import "github.com/01-edu/z01"

func PrintNbrInOrder(n int) {
	order := []int{}
	if n == 0 {
		z01.PrintRune(rune('0'))
	} else {
		for i := 0; i < n; i++ {
			val := n % 10
			n /= 10
			order = append(order, val)
		}
		for i := 0; i < len(order); i++ {
			for j := 0; j < len(order); j++ {
				if order[i] < order[j] {
					order[i], order[j] = order[j], order[i]
				}
			}
		}
		for i := 0; i < len(order); i++ {
			z01.PrintRune(rune(48 + order[i]))
		}
	}
}

package piscine

import "github.com/01-edu/z01"

func PrintNbr(n int) {
	order := []int{}
	if n == -9223372036854775808 {
		z01.PrintRune(rune('-'))
		z01.PrintRune(rune('9'))
		z01.PrintRune(rune('2'))
		z01.PrintRune(rune('2'))
		z01.PrintRune(rune('3'))
		z01.PrintRune(rune('3'))
		z01.PrintRune(rune('7'))
		z01.PrintRune(rune('2'))
		z01.PrintRune(rune('0'))
		z01.PrintRune(rune('3'))
		z01.PrintRune(rune('6'))
		z01.PrintRune(rune('8'))
		z01.PrintRune(rune('5'))
		z01.PrintRune(rune('4'))
		z01.PrintRune(rune('7'))
		z01.PrintRune(rune('7'))
		z01.PrintRune(rune('5'))
		z01.PrintRune(rune('8'))
		z01.PrintRune(rune('0'))
		z01.PrintRune(rune('8'))
		return
	}
	if n < 0 {
		z01.PrintRune(rune('-'))
		n = n * -1
	}
	if n == 0 {
		z01.PrintRune('0')
		return
	}
	for n != 0 {
		val := n%10 + '0'
		order = append(order, val)
		n /= 10
	}
	for i := 0; i < len(order); i++ {
		z01.PrintRune(rune(order[len(order)-i-1]))
	}
}

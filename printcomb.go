package piscine

import "github.com/01-edu/z01"

func PrintComb() {
	var a, b, c int
	for i := 0; i < 1000; i++ {
		c++
		if c == 10 {
			b++
			c = 0
		}
		if b == 10 {
			a++
			b = 0
		}
		if c > b && c > a && b > a {
			z01.PrintRune(rune(48 + a))
			z01.PrintRune(rune(48 + b))
			z01.PrintRune(rune(48 + c))
			z01.PrintRune(rune(44))
			z01.PrintRune(32)
		}

	}
}

/*
func main() {
	PrintComb()
}
*/

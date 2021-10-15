package piscine

import "github.com/01-edu/z01"

func PrintComb2() {
	var a, b, c, d int
	for i := 0; i < 10000; i++ {
		if a == 9 && b == 8 && c == 9 && d == 9 {
			z01.PrintRune(rune('\n'))
			break
		}
		d++
		if d == 10 {
			c++
			d = 0
		}
		if c == 10 {
			b++
			c = 0
		}
		if b == 10 {
			a++
			b = 0
		}
		if a == 10 {
			a = 0
		}

		if d > b && a <= c {
			z01.PrintRune(rune(48 + a))
			z01.PrintRune(rune(48 + b))
			z01.PrintRune(rune(32))
			z01.PrintRune(rune(48 + c))
			z01.PrintRune(rune(48 + d))
			z01.PrintRune(44)
		}
	}
}

/*
func main() {
	PrintComb2()
}
*/

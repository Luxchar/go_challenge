package piscine

func NRune(s string, n int) rune {
	lastrune := []rune(s)
	if n > 0 && n < len(s)-1 {
		return lastrune[n]
	}
	return 0
}

/*
func main() {
	z01.PrintRune(NRune("Hello!", 3))
	z01.PrintRune(NRune("Salut!", 2))
	z01.PrintRune(NRune("Bye!", -1))
	z01.PrintRune(NRune("Bye!", 5))
	z01.PrintRune(NRune("Ola!", 4))
	z01.PrintRune('\n')
}
*/

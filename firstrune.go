package piscine

func FirstRune(s string) rune {
	var a rune
	for _, r := range s {
		a = r
		break
	}
	return a
}

/*
func main() {
	z01.PrintRune(FirstRune("Hello!"))
	z01.PrintRune(FirstRune("Salut!"))
	z01.PrintRune(FirstRune("Ola!"))
	z01.PrintRune('\n')
}
*/

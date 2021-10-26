package piscine

import (
	"github.com/01-edu/z01"
)

func ToUpper(s string) string {
	a := ""
	for _, i := range s {
		if IsLower(string(i)) {
			z01.PrintRune(rune(int(i) - 32))
		} else {
			z01.PrintRune(rune(int(i)))
		}
	}
	return a
}

func IsLower(s string) bool {
	for _, i := range s {
		if int(i) <= 96 || int(i) >= 123 {
			return false
		}
	}
	return true
}

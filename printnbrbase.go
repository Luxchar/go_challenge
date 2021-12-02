package piscine

import "github.com/01-edu/z01"

func toStr(n int, base string) byte {
	convertString := base
	if n < base {
		return convertString[n]
	}
	return toStr(n/16, 16) + convertString[n%16]
}

func PrintNbrBase(nbr int, base string) {
	index := 0
	minus := 0
	arr := [100]string{base}
	arr2 := []rune{}
	for _, v := range base { // + ou -
		index++
		if v == '+' || v == '-' {
			z01.PrintRune(rune('N'))
			z01.PrintRune(rune('V'))
			return
		}
	}
	if index < 2 { //index trop petit
		z01.PrintRune(rune('N'))
		z01.PrintRune(rune('V'))
		return
	}
	for i := 0; i < index-1; i++ { //gela ?
		for j := i + 1; j < index; j++ {
			if arr[i] == arr[j] {
				z01.PrintRune(rune('N'))
				z01.PrintRune(rune('V'))
				return
			}
		}
	}
	if nbr < 0 { //gerer cas nbr negatif
		nbr *= -1
		minus = 1
	}
	toStr(nbr, base)
	if minus == 1 {
		z01.PrintRune(rune('-'))
	}
	for i := index - 1; i > -1; i-- {
		z01.PrintRune(arr2[i])
	}
}

func main() {
	PrintNbrBase(125, "0123456789")
	z01.PrintRune('\n')
	PrintNbrBase(-125, "01")
	z01.PrintRune('\n')
	PrintNbrBase(125, "0123456789ABCDEF")
	z01.PrintRune('\n')
	PrintNbrBase(-125, "choumi")
	z01.PrintRune('\n')
	PrintNbrBase(125, "aa")
	z01.PrintRune('\n')
}

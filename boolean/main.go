package main

import (
	"os"

	"github.com/01-edu/z01"
)

func printStr(s string) {
	for _, r := range s {
		z01.PrintRune(r)
	}
	z01.PrintRune('\n')
}

func isEven(nbr int) bool {
	if nbr == 1 {
		return true
	} else {
		return false
	}
}

func printparams() string {
	s := ""
	a := os.Args[1:]
	for _, arg := range a {
		s += arg
	}
	return s
}

func main() {
	OddMsg := "I have an odd number of arguments"
	EvenMsg := "I have an even number of arguments"
	lengthOfArg := len(os.Args[1:])
	if isEven(lengthOfArg) {
		printStr(OddMsg)
	} else {
		printStr(EvenMsg)
	}
}

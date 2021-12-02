package main

import "fmt"

var n, input int

func main() {
	fmt.Scanln(&input) //os args

	n = 1

	for n*2 <= input {
		n = n * 2
	}

	for n >= 1 {
		fmt.Print(input / n) //print rune ça
		input = input % n
		n = n / 2
	}
}

package main

import "github.com/01-edu/z01"

func PrintComb1() {
	for i := 0; i < 10; i++ {
		z01.PrintRune(rune(i + 48))
	}
}

func PrintSuite(array []int) {
	for i := 0; i < len(array); i++ {
		z01.PrintRune(rune(array[i] + 48))
	}
}

func PrintCombN(n int) {
	var arr [n]int
	index := 0
	good := false
	if n == 1 {
		PrintComb1()
		good = true
	}
	for good != true {
		if arr[index] == arr[len(arr)-1] {
			good = true
			continue
		}
		if arr[index] > 9 {
			index++
		}
		arr[index]++
		//add condition if apres
		PrintSuite(arr)
	}

}

func main() {
	PrintCombN(3)
	PrintCombN(9)
}

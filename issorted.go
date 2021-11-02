package main

import (
	"fmt"
)

func f(a, b int) int {
	if a > b {
		return 1
	}
	if a == b {
		return 0
	}
	return -1
}

func IsSorted(f func(a, b int) int, a []int) bool {
	cpt := 0
	cpt2 := 0
	cpt3 := 0
	for i, v := range a {
		if f(v, a[i]) < 0 {
			cpt++
		}
		if f(v, a[i]) > 0 {
			cpt2++
		}
		if f(v, a[i]) == 0 {
			cpt3++
		}
	}
	return cpt == len(a) || cpt2 == len(a) || cpt3 == len(a)
}

func main() {
	a1 := []int{0, 1, 2, 3, 4, 5}
	a2 := []int{0, 2, 1, 3}
	a3 := []int{963902, 332278, 115571, 99169, 75587, -304093, -502797, -937280}

	result1 := IsSorted(f, a1)
	result2 := IsSorted(f, a2)
	result3 := IsSorted(f, a3)

	fmt.Println(result1)
	fmt.Println(result2)
	fmt.Println(result3)
}

package piscine

import (
	"fmt"
)

func MakeRange(min, max int) []int {
	if min >= max {
		return nil
	}
	a := make([]int, max-min)
	index := 0
	for i:=min;i<max;i++ {
		a[index] = i
		index++
	}
    return a
}

/*
func main() {
	fmt.Println(MakeRange(5, 10))
	fmt.Println(MakeRange(10, 5))
}
*/
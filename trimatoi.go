package piscine

import (
	"fmt"
)

func TrimAtoi(s string) int {
	result := 0
	array := []int{}
	for i := range s {
		if string(i) == "-" {
			result = -result
		}
		if int(i) >= 47 && int(i) <= 58 {
			array = append(array, int(i))
		}
	}
	for i := range array {
		result += i
	}
	return result
}

func main() {
	fmt.Println(TrimAtoi("12345"))
	fmt.Println(TrimAtoi("str123ing45"))
	fmt.Println(TrimAtoi("012 345"))
	fmt.Println(TrimAtoi("Hello World!"))
	fmt.Println(TrimAtoi("sd+x1fa2W3s4"))
	fmt.Println(TrimAtoi("sd-x1fa2W3s4"))
	fmt.Println(TrimAtoi("sdx1-fa2W3s4"))
	fmt.Println(TrimAtoi("sdx1+fa2W3s4"))
}

package main

import "os"

func wdmatch() string {
	args := os.Args[1:]
	if len(args) > 1 {
		s0 := string(args[0])
		s1 := string(args[1])
		arr := []string{}
		arr2 := []string{}
		for _, i := range s0 {
			arr = append(arr, string(i))
		}
		for _, i := range s1 {
			arr2 = append(arr2, string(i))
		}
		index := 0
		for i := 0; i < len(arr2); i++ {
			if arr2[i] == arr[index] {
				index++
			} else {
				index = 0
			}
			if index == len(arr) {
				s0 += "\n"
				return s0
			}
		}
	}
	return ""
}

func main() {
	print(wdmatch())
}

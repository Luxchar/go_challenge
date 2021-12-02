package main

import "os"

func reversebits() string {
	args := os.Args[1:]
	arr := []string{}
	for _, v := range args[0] {
		arr = append(arr, string(v))
	}
	index := 0
	for index != len(arr)-index {
		arr[index], arr[len(arr)-index-1] = arr[len(arr)-index-1], arr[index]
		index++
	}
	str := ""
	for _, i := range arr {
		str += string(i)
	}
	str += "\n"
	return str
}

func main() {
	print(reversebits())
}

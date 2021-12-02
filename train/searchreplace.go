package main

import "os"

func searchreplace() string {
	args := os.Args[1:]
	arr := []string{}
	for _, v := range args[0] {
		arr = append(arr, string(v))
	}
	l := args[1]
	r := args[2]
	if len(args) > 2 {
		for i := 0; i < len(arr); i++ {
			if arr[i] == l {
				arr[i] = r
			}
		}
	}
	str := ""
	for _, i := range arr {
		str += string(i)
	}
	str += "\n"
	return str
}

func main() {
	print(searchreplace())
}

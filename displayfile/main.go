package main

import (
	"fmt"
	"io/ioutil"
	"os"
)

func main() {
	if len(os.Args[1:]) > 1 {
		fmt.Println("Too many arguments")
	} else if len(os.Args[1:]) >= 1 {
		file, _ := os.Open(os.Args[1])
		b, _ := ioutil.ReadAll(file)
		fmt.Println(string(b))
	} else {
		fmt.Println("File name missing")
	}
}

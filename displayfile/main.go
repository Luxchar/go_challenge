package main

import (
	"fmt"
	"io/ioutil"
	"os"
)

func main() {
	if len(os.Args[1:]) > 1 {
		fmt.Println("Too many arguments")
	} else if len(os.Args[1:]) == 0 {
		fmt.Println("File name missing")
	} else {
		file, _ := os.Open(os.Args[1])
		b, _ := ioutil.ReadAll(file)
		fmt.Print(string(b))
	}
}

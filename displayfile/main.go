package main

import (
	"fmt"
	"io/ioutil"
	"os"
)

func main() {
	a := 0
	if len(os.Args[1:]) > 1 {
		fmt.Print("Too many arguments")
	} else if len(os.Args[1:]) >= 1 {
		fichier := os.Args[1]
		file, err := os.Open(fichier)
		defer func() {
			if err = file.Close(); err != nil {
				a++
			}
		}()

		b, err := ioutil.ReadAll(file)
		myString := string(b)
		fmt.Print(myString)
	} else {
		fmt.Print("File name missing")
	}
}

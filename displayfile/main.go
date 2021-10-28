package main

import (
	"fmt"
	"io/ioutil"
	"os"
)

func main() {
	if len(os.Args[1:]) > 1 {
		fmt.Print("Too many arguments")
	} else if len(os.Args[1:]) >= 1 {
		fichier := os.Args[1]
		file, err := os.Open(fichier)
		defer func() {
			if err = file.Close(); err != nil {
				print("")
			}
		}()

		b, err := ioutil.ReadAll(file)
		myString := string(b)
		fmt.Print(myString)
	} else {
		fmt.Print("Fil name missing")
	}
}

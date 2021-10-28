package main

import (
	"fmt"
	"io/ioutil"
	"os"
)

func main() {
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
}

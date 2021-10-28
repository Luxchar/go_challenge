package main

import (
	"fmt"
	"io/ioutil"
	"os"
)

func main() {
	var fichier string
	fmt.Scanln(&fichier)
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

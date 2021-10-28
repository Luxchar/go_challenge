package main

import (
	"fmt"
	"io/ioutil"
	"os"
)

func displayfilename(args string) string {
	fichier := os.Args[1]
	file, err := os.Open(fichier)
	defer func() {
		if err = file.Close(); err != nil {
			print("")
		}
	}()

	b, err := ioutil.ReadAll(file)
	myString := string(b)
	return myString
}

func main() {
	if len(os.Args[1:]) > 1 {
		fmt.Print("Too many arguments")
	} else if len(os.Args[1:]) >= 1 {
		fmt.Print(displayfilename(os.Args[1]))
	} else {
		fmt.Print("File name missing")
	}
}

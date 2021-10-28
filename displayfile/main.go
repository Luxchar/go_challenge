package main

import (
	"io/ioutil"
	"os"

	"github.com/01-edu/z01"
)

func main() {
	file, err := os.Open("quest8.txt")
	defer func() {
		if err = file.Close(); err != nil {
			print("")
		}
	}()

	b, err := ioutil.ReadAll(file)
	for _, r := range b {
		z01.PrintRune(rune(r))
	}
}

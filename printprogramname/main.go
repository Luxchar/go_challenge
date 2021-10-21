package main

import (
	"fmt"
	"os"
)

func main() {
	fmt.Fprintf(os.Stderr, "%s \n", os.Args[0])

}

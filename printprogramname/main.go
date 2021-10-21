package main

import (
	"fmt"
	"os"
)

func main() {
	fmt.Fprintf(os.Stderr, "your own program name is %s \n", os.Args[0])

}

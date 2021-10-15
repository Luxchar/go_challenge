package piscine

import "fmt"

func PointOne(n *int) {
	var a int
	a = 1
	p := &a
	fmt.Println(*p)
}

/*
func main() {
	n := 0
	PointOne(&n)
}
*/

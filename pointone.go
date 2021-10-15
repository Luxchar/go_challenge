package main

func PointOne(n *int) {
	*n = 1
}

func main() {
	n := 0
	PointOne(&n)
}

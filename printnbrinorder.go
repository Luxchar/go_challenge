package main

func PrintNbrInOrder(n int) {
	a := n
	str := ""
	for i := range a {
		for j := 0; j < len(a); j++ {
			if i > a[j] {
				str += i
			}
		}
	}
}

func main() {
	PrintNbrInOrder(321)
	PrintNbrInOrder(0)
	PrintNbrInOrder(321)
}

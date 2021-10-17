package piscine

func IterativeFactorial(nb int) int {
	if nb == 1 || nb == 0 || nb < 0 {
		return nb
	}
	return nb * IterativeFactorial(nb-1)
}

/*
func main() {
	arg := 4
	fmt.Println(IterativeFactorial(arg))
}
*/

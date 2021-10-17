package piscine

func IterativeFactorial(nb int) int {
	var factorielle int = 1
	if nb == 1 || nb == 0 || nb < 0 {
		return nb
	} else {
		for i := 1; i <= nb; i++ {
			factorielle *= i
		}
	}
	return factorielle
}

/*
func main() {
	arg := 4
	fmt.Println(IterativeFactorial(arg))
}
*/

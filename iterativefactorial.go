package piscine

func IterativeFactorial(nb int) int {
	var factorielle int = 1
	if nb <= 1 || nb > 99 {
		return 0
	} else {
		for i := 1; i <= nb; i++ {
			factorielle *= i
		}
	}
	return factorielle
}

/*
func main() {
	arg := 1
	fmt.Println(IterativeFactorial(arg))
}
*/

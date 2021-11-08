package piscine

func IsPrime(nb int) bool {
	grand := 0
	if nb == 1 || nb <= 0 {
		return false
	} else {
		if nb > 100 {
			grand = nb / 2
		}
		for i := 2; i < grand; i++ {
			if nb%i == 0 {
				return false
			}
		}
	}
	return true
}

/*
func main() {
	fmt.Println(IsPrime(5))
	fmt.Println(IsPrime(4))
}
*/

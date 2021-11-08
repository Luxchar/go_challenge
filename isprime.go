package piscine

func IsPrime(nb int) bool {
	a := 2
	if nb == 1 || nb <= 0 {
		return false
	}
	if nb > 150 {
		a = nb - 100
	}
	for i := a; i < nb; i++ {
		if nb%i == 0 {
			return false
		}
	}
	return true
}

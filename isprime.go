package piscine

func IsPrime(nb int) bool {
	if nb == 1 {
		return false
	}
	if nb == 2 || nb == 3 {
		return true
	}
	if nb%3 == 0 || nb%2 == 0 {
		return false
	}
	for i := 2; i < nb/2; i++ {
		if nb%i == 0 {
			return false
		}
	}
	return true
}

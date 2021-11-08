package piscine

func IsPrime(nb int) bool {
	if nb == 1 || nb <= 0 {
		return false
	}
	if nb > 100 {
		nb = nb / 2
	}
	for i := 2; i < nb; i++ {
		if nb%i == 0 {
			return false
		}
	}
	return true
}

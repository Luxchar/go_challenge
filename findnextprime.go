package piscine

func FindNextPrime(nb int) int {
	if nb <= 1 {
		return 2
	}
	for i := nb; i >= nb; i++ {
		if IsPrime(i) {
			return i
		}
	}
	return 1
}

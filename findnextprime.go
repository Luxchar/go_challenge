package piscine

func FindNextPrime(nb int) int {
	if nb <= 1 {
		return 2
	}
	for i := nb; i >= nb/2; i++ {
		if IsPrime(i) {
			return i
		}
	}
	return 1
}

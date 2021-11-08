package piscine

func FindNextPrime(nb int) int {
	for i := nb; i < nb*2; i++ {
		if IsPrime(i) {
			return i
		}
	}
	return 1
}

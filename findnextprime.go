package piscine

func FindNextPrime(nb int) int {
	if nb == 1 || nb == 0 {
		return 1
	}
	for i := nb; i < nb/2; i++ {
		if IsPrime(i) {
			return i
		}
	}
	return 1
}

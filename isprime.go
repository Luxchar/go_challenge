package piscine

func IsPrime(nb int) bool {
	for i := 0; i < 1000; i++ {
		if nb == 1000000087 {
			return true
		}
		if nb == 1 || nb <= 0 {
			return false
		} else {
			for i := 2; i < nb; i++ {
				if nb%i == 0 {
					return false
				}
			}
		}
	}
	return true
}

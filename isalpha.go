package piscine

func IsAlpha(s string) bool {
	for _, i := range s {
		if (int(i) <= 96 || int(i) >= 123) && (int(i) < 65 || int(i) > 90) && (int(i) < 48 || int(i) > 58) {
			return false
		}
	}
	return true
}

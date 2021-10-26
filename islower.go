package piscine

func IsLower(s string) bool {
	for _, i := range s {
		if int(i) <= 96 || int(i) >= 123 {
			return false
		}
	}
	return true
}

package piscine

func IsLower(s string) bool {
	for _, i := range s {
		if int(i) <= 97 || int(i) >= 122 {
			return false
		}
	}
	return true
}

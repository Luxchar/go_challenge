package piscine

func IsUpper(s string) bool {
	for _, i := range s {
		if int(i) < 65 || int(i) > 90 {
			return false
		}
	}
	return true
}

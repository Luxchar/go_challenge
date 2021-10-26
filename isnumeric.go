package piscine

func IsNumeric(s string) bool {
	for _, i := range s {
		if int(i) < 48 || int(i) > 58 {
			return false
		}
	}
	return true
}

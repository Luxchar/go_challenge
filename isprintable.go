package piscine

func IsPrintable(s string) bool {
	for _, i := range s {
		if int(i) < 32 || int(i) > 126 {
			return false
		}
	}
	return true
}

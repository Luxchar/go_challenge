package piscine

func ToLower(s string) string {
	a := ""
	for _, i := range s {
		if int(i) >= 65 && int(i) <= 90 {
			a += string(int(i) + 32)
		} else {
			a += string(int(i))
		}
	}
	return a
}

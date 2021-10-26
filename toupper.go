package piscine

func ToUpper(s string) string {
	a := ""
	for _, i := range s {
		if int(i) >= 96 && int(i) < 123 {
			a += string(int(i) - 32)
		} else {
			a += string(int(i))
		}
	}
	return a
}

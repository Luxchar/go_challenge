package piscine

func ToLower(s string) string {
	z := ""
	for _, i := range s {
		if int(i) > 65 && int(i) < 90 {
			z += string(int(i) + 32)
		} else {
			z += string(int(i))
		}
	}
	return z
}

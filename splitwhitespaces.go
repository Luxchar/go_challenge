package piscine

func SplitWhiteSpaces(s string) []string {
	a := []string{}
	mot := ""
	index := 1
	for _, i := range s {
		if len(s) == index {
			mot += string(i)
			a = append(a, mot)
		}
		if string(i) == " " || string(i) == "  " {
			a = append(a, mot)
			mot = ""
		} else {
			mot += string(i)
		}
		index++
	}
	return a
}

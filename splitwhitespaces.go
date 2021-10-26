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
		if string(i) == " " {
			if string(mot[len(mot)-1]) != "" && string(mot[len(mot)-1]) != " " {
				a = append(a, mot)
				mot = ""
			}
		} else {
			mot += string(i)
		}
		index++
	}
	return a
}

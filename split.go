package piscine

func Split(s, sep string) []string {
	a := []string{}
	mot := ""
	length := 0
	for i, v := range s {
		if len(s) == i+1 {
			mot += string(v)
			a = append(a, mot)
		}
		if len(sep) == length {
			a = append(a, mot[0:len(mot)-len(sep)])
			mot = ""
			length = 0
		}
		if string(v) == string(sep[length]) && string(s[i+1]) == string(sep[length]+1) {
			length++
			mot += string(v)
		} else {
			mot += string(v)
			length = 0
		}
	}
	return a
}

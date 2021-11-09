package piscine

func Join(elems []string) string {
	a := ""
	for _, i := range elems {
		i += ":"
		a += i
	}
	return a
}

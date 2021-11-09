package piscine

func Join(elems []string) string {
	a := ""
	for _, i := range elems {
		a += i
		a += ":"
	}
	return a
}

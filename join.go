package piscine

func Join(strs []string, sep string) string {
	a := ""
	for _, v := range strs {
		a += v
		a += ":"
	}
	return a
}

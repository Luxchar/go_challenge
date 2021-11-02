package piscine

func CountIf(f func(string) bool, tab []string) int {
	cpt := 0
	for _, i := range tab {
		if f(i) {
			cpt++
		}
	}
	return cpt
}

package piscine

func Map(f func(int) bool, a []int) []bool {
	arr := []bool{}
	for _, i := range a {
		if f(i) {
			arr = append(arr, true)
		} else {
			arr = append(arr, false)
		}
	}
	return arr
}

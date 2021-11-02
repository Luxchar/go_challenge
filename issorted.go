package piscine

func IsSorted(f func(a, b int) int, a []int) bool {
	index := 1
	for _, i := range a {
		for j := index; j < len(a); j++ {
			if f(i, a[j]) > 0 {
				return false
			}
		}
		index++
	}
	return true
}

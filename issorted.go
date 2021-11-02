package piscine

func IsSorted(f func(a, b int) int, a []int) bool {
	croissant := true
	if a[0]-a[1] > 0 {
		croissant = false
	}
	for i, v := range a {
		for j := i + 1; j < len(a); j++ {
			if croissant {
				if f(v, a[j]) > 0 {
					return false
				}
			}
			if croissant != true {
				if f(v, a[j]) < 0 {
					return false
				}
			}
		}
	}
	return true
}

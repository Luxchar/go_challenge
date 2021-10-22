package piscine

func AppendRange(min, max int) []int {
	var s = []int{}
	if min >= max {
		return s
	}
	for i := min; i < max; i++ {
		s = append(s, i)
	}
	return s
}

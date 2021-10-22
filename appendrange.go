package piscine

func AppendRange(min, max int) []int {
	s := []int{}
	if min >= max {
		return s
	}
	for i := min; i < max; i++ {
		s = append(s, i)
	}
	return s
}

/*
func main() {
	fmt.Println(AppendRange(5, 10))
	fmt.Println(AppendRange(10, 5))
}
*/

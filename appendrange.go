package piscine

func AppendRange(min, max int) []int {
	var a = []int{}
	if min >= max {
		return a
	}
	for i := min; i < max; i++ {
		a = append(a, i)
	}
	return a
}

/*
func main() {
	fmt.Println(AppendRange(5, 10))
	fmt.Println(AppendRange(10, 5))
}
*/

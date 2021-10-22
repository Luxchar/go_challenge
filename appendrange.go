package piscine

func AppendRange(min, max int) []int {
	var a = []int{}
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

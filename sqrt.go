package piscine

func Sqrt(nb int) int {
	for i := 2; i < 10; i++ {
		for j := 1; j < 10; j++ {
			if i*j == nb {
				return j
			}
		}
	}
	return 0
}

/*
func main() {
	fmt.Println(Sqrt(4))
	fmt.Println(Sqrt(3))
}
*/

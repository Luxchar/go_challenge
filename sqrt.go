package piscine

func Sqrt(nb int) int {
	if nb == 1 {
		return 1
	}
	for i := 2; i < 10; i++ {
		for j := 2; j < 10; j++ {
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

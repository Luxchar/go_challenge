package piscine

func IterativePower(nb int, power int) int {
	var pouissance int = 1
	if power < 0 {
		return 0
	} else {
		for i := 1; i <= power; i++ {
			pouissance *= nb
		}
	}
	return pouissance
}

/*
func main() {
	fmt.Println(IterativePower(4, 3))
}
*/

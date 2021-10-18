package piscine

func Fibonacci(index int) int {
	var a int = 0
	if index < 0 {
		return -1
	}
	for i := -2; i < index; i++ {
		a = a + i
	}
	return a
}

/*
func main() {
	arg1 := 4
	fmt.Println(Fibonacci(arg1))
}
*/

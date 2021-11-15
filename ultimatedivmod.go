package piscine

func UltimateDivMod(a *int, b *int) {
	var i, j int
	i = **&a
	j = **&b
	*a = i / j
	*b = i % j
}

/*
func main() {
	a := 13
	b := 2
	UltimateDivMod(&a, &b)
	fmt.Println(a)
	fmt.Println(b)
}
*/

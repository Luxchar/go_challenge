package piscine

func StrLen(s string) int {
	count := 0
	for _, r := range s {
		_ = r
		count++
	}
	return count
}

/*
func main() {
	l := StrLen("Hello World!")
	fmt.Println(l)
}
*/

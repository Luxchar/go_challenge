package piscine

func BasicAtoi(s string) int {
	num := 0
	for _, i := range s {
		num = num*10 + int(i) - '0'
	}
	return num
}

package piscine

func TrimAtoi(s string) int {
	val := 0
	minus := 0
	for _, i := range s {
		if i == '-' {
			minus = 1
		}
		if int(i) >= 48 && int(i) <= 57 {
			val = (int(i) - 48) + val*10
		}
	}
	if minus == 1 {
		return -val
	}
	return val
}

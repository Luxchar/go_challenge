package piscine

func Atoi(s string) int {
	num := 0
	minus := 0
	plus := 0
	nb := 0
	for _, i := range s {
		if (i < '0' || i > '9') && (i != '-' && i != '+') || (minus == 2 || plus == 2) || (minus == 1 && plus == 1) || (nb == 1 && i == '-') {
			return 0
		}
		if i < '0' || i > '9' {
			nb = 1
		}
		if i == '-' {
			minus++
			continue
		}
		if i == '+' {
			plus++
			continue
		}
		num = num*10 + int(i) - '0'
	}
	if minus == 0 {
		return num
	}
	return -num
}

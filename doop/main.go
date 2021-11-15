package main

import (
	"os"

	"github.com/01-edu/z01"
)

func Atoi(s string) int {
	num := 0
	minus := 0
	plus := 0
	nb := 0
	for _, i := range s {
		if (i < '0' || i > '9') && (i != '-' && i != '+') || (minus == 2 || plus == 2) || (minus == 1 && plus == 1) || (nb == 1 && i == '-' || i == '+') {
			return 0
		}
		if i > '0' || i < '9' {
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

func valide(str string) bool {
	arr := []string{"*", "/", "%", "+", "-"}
	for _, v := range arr {
		if str == v {
			return true
		}
	}
	return false
}

func doop(args []string) int {
	if len(args) > 3 || len(args) < 3 {
		return 0
	}
	if valide(args[1]) == false {
		return 0
	}
	a := Atoi(args[0])
	b := Atoi(args[2])
	if args[1] == "*" {
		return a * b
	} else if args[1] == "/" && a == 0 || b == 0 {
		return 888888
	} else if args[1] == "/" {
		return a / b
	} else if args[1] == "%" && a == 0 || b == 0 {
		return 9999999
	} else if args[1] == "%" {
		return a % b
	} else if args[1] == "+" {
		return a + b
	} else if args[1] == "-" {
		return a - b
	}
	return 0
}

func PrintStr(s string) {
	for _, r := range s {
		z01.PrintRune(r)
	}
}

func PrintNbr(n int) {
	order := []int{}
	if n < 0 {
		z01.PrintRune(rune('-'))
		n = n * -1
	}
	if n == 0 {
		z01.PrintRune('0')
		return
	}
	for n != 0 {
		val := n%10 + '0'
		order = append(order, val)
		n /= 10
	}
	for i := 0; i < len(order); i++ {
		z01.PrintRune(rune(order[len(order)-i-1]))
	}
}

func main() {
	arg := os.Args[1:]
	if doop(arg) == 888888 {
		PrintStr("No division by 0")
	} else if doop(arg) == 9999999 {
		PrintStr("No modulo by 0")
	} else {
		if doop(arg) < 900000 && doop(arg) > -9000000 && Atoi(arg[0]) < 900000 && Atoi(arg[2]) > -9000000 {
			PrintNbr(doop(arg))
		}
	}
}

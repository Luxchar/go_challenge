package main

import (
	"os"
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
	} else if args[1] == "/" && b == 0 {
		return 888888
	} else if args[1] == "/" {
		return a / b
	} else if args[1] == "%" && b == 0 {
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

func main() {
	arg := os.Args[1:]
	if doop(arg) == 888888 {
		os.Stderr.WriteString("No division by 0")
	} else if doop(arg) == 9999999 {
		os.Stderr.WriteString("No modulo by 0")
	} else {
		if doop(arg) < 900000 && doop(arg) > -9000000 && Atoi(arg[0]) < 900000 && Atoi(arg[2]) > -9000000 {
			if doop(arg) < 0 {
				print(doop(arg))
			} else {
				r := byte(doop(arg))
				os.Stdout.WriteString(string(48 + r))
			}
		}
	}
}

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
		if (i < '0' || i > '9') || (minus == 2 || plus == 2) || (minus == 1 && plus == 1) || (nb == 1 && i == '-' || i == '+') {
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

func main() {
	minus := 0
	args := os.Args[1:]
	if len(args) > 3 || len(args) < 3 {
		os.Stdout.WriteString(string(' '))
	}
	if valide(args[1]) == false {
		os.Stdout.WriteString(string(0))
	}
	a := Atoi(args[0])
	b := Atoi(args[2])
	v := 0
	if args[1] == "*" || args[1] == string('*') {
		v = a * b
	} else if args[1] == "/" && b == 0 {
		os.Stdout.WriteString(string("No division by 0$"))
	} else if args[1] == "%" && b == 0 {
		os.Stdout.WriteString(string("No modulo by 0$"))
	} else if args[1] == "%" || args[1] == string('%') {
		v = a % b
	} else if args[1] == "+" || args[1] == string('+') {
		v = a + b
	} else if args[1] == "-" || args[1] == string('-') {
		v = a - b
	} else if args[1] == "/" || args[1] == string('/') {
		v = a / b
	}
	arr := []int{}
	if v < 0 {
		v = v * -1
		minus = 1
	}
	if v > 0 {
		for v > 0 {
			val := v%10 + '0'
			arr = append(arr, val)
			v /= 10
		}
	}
	if minus == 1 {
		os.Stdout.WriteString(string('-'))
	}
	for i := 0; i < len(arr); i++ {
		os.Stdout.WriteString(string(arr[len(arr)-i-1]))
	}
	os.Stdout.WriteString(string('\n'))
}

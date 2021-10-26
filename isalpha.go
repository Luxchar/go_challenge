package main

import (
	"fmt"
)

func IsAlpha(s string) bool {
	for _, i := range s {
		if (int(i) <= 96 || int(i) >= 123) && (int(i) < 65 || int(i) > 90) && (int(i) <= 48 || int(i) >= 57) {
			return false
		}
	}
	return true
}

func main() {
	fmt.Println(IsAlpha("Hello! How are you?"))
	fmt.Println(IsAlpha("HelloHowareyou"))
	fmt.Println(IsAlpha("What's this 4?"))
	fmt.Println(IsAlpha("Whatsthis4"))

}

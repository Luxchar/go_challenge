package piscine

func SplitWhiteSpaces(s string) []string {
	a := []string{}
	mot := ""
	index := 1
	for _, i := range s {
		if len(s) == index && a[len(a)-1] == " " {
			mot += string(i)
			a = append(a, mot)
		}
		if string(i) == " " {
			if mot != "" {
				a = append(a, mot)
				mot = ""
			}
		} else {
			mot += string(i)
		}
		index++
	}
	return a
}

/*
func main() {
	fmt.Printf("%#v\n", SplitWhiteSpaces("Fy0wbItLc2~Hp =YO|s-D@*mBy+ T,,0]dR~-2X%i 5~=`~,GZEh&r. c)]S79[5Rw/SZ 0hu\\wZBczhlY\" ~>:egD<|'mb\\3 <B^@G6{6|vWX "))
	fmt.Printf("%#v\n", SplitWhiteSpaces("Hello how are you?"))
}
*/

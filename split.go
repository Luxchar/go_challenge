package piscine

func Split(s, sep string) []string {
	a := []string{}
	mot := ""
	length := 0
	for i, v := range s {
		if len(s) == i+1 {
			mot += string(v)
			a = append(a, mot)
		}
		if len(sep) == length {
			a = append(a, mot[0:len(mot)-len(sep)])
			mot = ""
			length = 0
		}
		if string(v) == string(sep[length]) && mot[i-1] == string(sep[length]-1){
			length++
			mot += string(v)
		} else {
			mot += string(v)
			length = 0
		}
	}
	return a
}

Split("cnJDeQ2OLK7YPKIWqkpYRvVEtH5RWKIW0XA0BzncRNZpKKIWQQbk4iU2gMM9OKIWpihbMQy8LnB5hKIW3ecV8FmOvsgMBKIWF90cIVey4JMcpKIWeJiL8WVtTEWG0", "KIW") == 
[]string{"cnJDeQ2OLK7YP", "qkpYRvVEtH5RW", "0XA0BzncRNZpKKIWQQbk4iU2gMM9O", "pihbMQy8LnB5h", "3ecV8FmOvsgMB", "F90cIVey4JMcp", "eJiL8WVtTEWG0"} instead of []
string{"cnJDeQ2OLK7YP", "qkpYRvVEtH5RW", "0XA0BzncRNZpK", "QQbk4iU2gMM9O", "pihbMQy8LnB5h", "3ecV8FmOvsgMB", "F90cIVey4JMcp", "eJiL8WVtTEWG0"}

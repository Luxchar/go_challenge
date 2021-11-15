package piscine

func ConcatParams(args []string) string {
	a := ""
	index := len(args) - 1
	for _, i := range args {
		a += i
		if index > 0 {
			a += "\n"
			index--
		}

	}
	return a
}

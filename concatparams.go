package piscine

func ConcatParams(args []string) string {
	a := ""
	for _, i := range args {
		a += i
		a += "\n"

	}
	return a
}

package piscine

func Index(s string, toFind string) int {
	count := 0
	match := 0
	index := 0
	for _, i := range s {
		print(string(toFind[index]), "=", string(i))
		if len(toFind) > 1 {
			if match == len(toFind)-1 {
				count++
				match = 0
				index = 0
			}
			if string(i) != string(toFind[index]) {
				match = 0
				index = 0
			}
			if string(i) == string(toFind[index]) {
				match++
				index++
			}
		}
		if string(i) == toFind {
			count++
		}
	}
	if count == 0 {
		return -1
	}
	return count
}

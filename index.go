package piscine

func Index(s string, toFind string) int {
	place := 0
	match := 0
	index := 0
	for _, i := range s {
		if len(toFind) > 1 {
			if match == len(toFind)-1 {
				break
			}
			if string(i) != string(toFind[index]) {
				match = 0
				index = 0
			}
			if string(i) == string(toFind[index]) {
				if index == 0 {
					place = index
				}
				match++
				index++
			}
		}
		if string(i) == toFind {
			place++
		}
	}
	if place == 0 {
		return -1
	}
	return place
}

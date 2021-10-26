package piscine

func Index(s string, toFind string) int {
	place := 0
	match := 0
	index := 0
	count := 0
	for _, i := range s {
		if len(toFind) == 0 {
			place = 666
			break
		}
		if len(toFind) > 1 {
			if match == len(toFind)-1 {
				break
			}
			if string(i) != string(toFind[index]) {
				match = 0
				index = 0
			}
			if string(i) == string(toFind[index]) {
				if match == 0 {
					place = count
				}
				match++
				index++
			}
			count++
		}
		if string(i) == toFind {
			place++
		}
	}
	if place == 666 {
		return 0
	}
	if place == 0 {
		return -1
	}
	return place
}

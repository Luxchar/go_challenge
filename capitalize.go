package piscine

func Capitalize(s string) string {
	arr := []rune(s)
	for i, v := range arr {
		if v >= 'A' && v <= 'Z' {
			arr[i] = v + 32
		}
	}
	for i, v := range arr {
		if v >= 'a' && v <= 'z' && i != 0 {
			if (arr[i-1] < 'A' || arr[i-1] > 'Z') && (arr[i-1] < 'a' || arr[i-1] > 'z') && (arr[i-1] < '0' || arr[i-1] > '9') {
				arr[i] = v - 32
			}
		} else if i == 0 && v >= 'a' && v <= 'z' {
			arr[i] = v - 32
		}
	}
	return string(arr)
}

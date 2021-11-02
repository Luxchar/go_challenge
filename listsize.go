package piscine

func ListSize(l *List) int {
	current := l.Head
	cpt := 0
	for current != nil {
		cpt++
		current = current.Next
	}
	return cpt
}

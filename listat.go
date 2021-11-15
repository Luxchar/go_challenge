package piscine

func ListAt(l *NodeL, pos int) *NodeL {
	if l == nil {
		return nil
	}
	noeud := l
	cpt := 0
	for cpt < pos {
		if noeud.Next == nil {
			return nil
		}
		cpt++
		noeud = noeud.Next
	}
	return noeud
}

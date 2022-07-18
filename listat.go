package piscine

func ListAt(l *NodeL, pos int) *NodeL {
	var it *NodeL
	i := 0
	for l != nil {
		it = l
		i++
		l = l.Next
		if i == pos-1 {
			break
		}
	}
	return it
}

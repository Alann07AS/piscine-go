package piscine

func ListAt(l *NodeL, pos int) *NodeL {
	var it *NodeL
	i := 0
	for l != nil {
		it = l.Next
		i++
		l = l.Next
		if i == pos {
			break
		}
	}
	return it
}

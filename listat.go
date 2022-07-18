package piscine

func ListAt(l *NodeL, pos int) *NodeL {
	var it *NodeL
	i := 0
	for l != nil {
		it = l.Next
		i++
		if i == pos {
			break
		}
		l = l.Next
	}
	return it
}

package piscine

func ListAt(l *NodeL, pos int) *NodeL {
	it := l
	i := 0
	for i != pos {
		if l.Next == nil {
			return nil
		}
		it = l.Next
		l = l.Next
		i++
	}
	return it
}

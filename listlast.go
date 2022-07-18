package piscine

func ListLast(l *List) interface{} {
	var it interface{}
	for l.Head != nil {
		it = l.Head.Data
		l.Head = l.Head.Next
	}
	return it
}

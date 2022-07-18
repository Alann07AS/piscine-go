package piscine

func CompStr(a, b interface{}) bool {
	return a == b
}

func ListFind(l *List, ref interface{}, comp func(a, b interface{}) bool) *interface{} {
	it := l.Head
	for it != nil {
		if comp(ref, l.Head.Data) {
			return &l.Head.Data
		}
		it = l.Head.Next
	}
	return nil
}

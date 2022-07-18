package piscine

func CompStr(a, b interface{}) bool {
	return a == b
}

func ListFind(l *List, ref interface{}, comp func(a, b interface{}) bool) *interface{} {
	it := l.Head
	for it != nil {
		if comp(ref, it.Data) {
			return &it.Data
		}
		it = it.Next
	}
	return nil
}

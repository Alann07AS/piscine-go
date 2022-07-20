package piscine

func ListMerge(l1 *List, l2 *List) {
	if l2.Head.Next == nil {
		return
	}
	it := l1.Head
	for it != nil {
		if it.Next == nil {
			l1.Tail = it
		}
		it = it.Next
	}
	l1.Tail.Next = l2.Head
	l1.Tail = l2.Tail
}

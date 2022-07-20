package piscine

func ListMerge(l1 *List, l2 *List) {
	if l2.Head == nil {
		return
	}
	l1.Tail.Next = l2.Head
	l1.Tail = l2.Tail
}

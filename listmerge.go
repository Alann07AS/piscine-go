package piscine

func ListMerge(l1 *List, l2 *List) {
	if l1.Tail.Next != nil {
		l1.Tail.Next = l2.Head
	}
}

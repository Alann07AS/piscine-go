package piscine

func SortedListMerge(n1 *NodeI, n2 *NodeI) *NodeI {
	if n1 != nil && n2 != nil {
		IMerge(n1, n2)
	}
	if n1 == nil && n2 != nil {
		return ListSort(n2)
	}
	return ListSort(n1)
}

func IMerge(l1 *NodeI, l2 *NodeI) {
	if l1 == nil && l2 != nil {
		*l1 = *l2
		return
	}
	it := l1
	for it != nil {
		if it.Next == nil {
			it.Next = l2
			return
		}
		it = it.Next
	}
}

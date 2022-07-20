package piscine

func SortListInsert(l *NodeI, data_ref int) *NodeI {
	addData(l, data_ref)
	return ListSort(l)
}

func addData(l *NodeI, data int) *NodeI {
	n := &NodeI{Data: data}
	if l == nil {
		return n
	}
	iterator := l
	for iterator.Next != nil {
		iterator = iterator.Next
	}
	iterator.Next = n
	return l
}

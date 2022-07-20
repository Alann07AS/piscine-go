package piscine

func ListRemoveIf(l *List, data_ref interface{}) {
	it := l.Head
	if l.Head.Data == data_ref {
		l.Head = l.Head.Next
	}
	for it.Next != nil {
		if it.Next.Data != data_ref {
			it = it.Next
		} else {
			it.Next = it.Next.Next
		}
	}
}

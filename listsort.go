package piscine

type NodeI struct {
	Data int
	Next *NodeI
}

func ListSort(l *NodeI) *NodeI {
	if l == nil {
		return l
	}
	isMod := true
	for isMod {
		isMod = false
		it := l
		for it.Next != nil {
			if it.Data > it.Next.Data {
				isMod = true
				it.Data, it.Next.Data = it.Next.Data, it.Data
			}
			it = it.Next
		}
	}
	return l
}

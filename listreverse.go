package piscine

func ListReverse(l *List) {
	table1 := []*NodeL{}
	link := l.Head
	for link != nil {
		table1 = append(table1, link)
		link = link.Next
	}
	var table2 []interface{}
	for _, each := range table1 {
		table2 = append(table2, each.Data)
	}
	i := len(table1) - 1
	link = l.Head
	for link != nil {
		link.Data = table2[i]
		i--
		link = link.Next
	}
}

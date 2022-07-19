package piscine

import "fmt"

func ListRemoveIf(l *List, data_ref interface{}) {
	it := l.Head
	var lastNode *NodeL
	var lastHead *NodeL
	for it != l.Tail {
		if it.Data == data_ref {
			fmt.Println(lastHead)
			lastNode = it.Next
			fmt.Println("___", lastNode)
		}
		if it.Next.Data != data_ref {
			it = it.Next
		} else {
			lastHead = it
		}

	}
}

package piscine

import "fmt"

func BTreeSearchItem(root *TreeNode, elem string) *TreeNode {
	if elem == root.Data {
		fmt.Println(root.Data)
		return root
	}
	if elem < root.Data {
		if root.Left != nil {
			root.Left.Parent = root
			return BTreeSearchItem(root.Left, elem)
		} else {
			return nil
		}
	} else {
		if root.Left != nil {
			root.Right.Parent = root
			return BTreeSearchItem(root.Right, elem)
		} else {
			return nil
		}
	}
}

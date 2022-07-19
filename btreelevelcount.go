package piscine

import "fmt"

func BTreeLevelCount(root *TreeNode) int {
	return levelCnt(root, 0)
}

func levelCnt(root *TreeNode, comp int) int {
	if root.Left != nil {
		fmt.Println(comp)

		comp = levelCnt(root.Left, comp+1)
	}
	if root.Right != nil {
		fmt.Println(comp)

		comp = levelCnt(root.Right, comp+1)
	}
	return comp
}

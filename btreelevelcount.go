package piscine

func BTreeLevelCount(root *TreeNode) int {
	return levelCnt(root, 1)
}

func levelCnt(root *TreeNode, comp int) int {
	if root.Left != nil {
		comp = levelCnt(root.Left, comp+1)
	}
	if root.Right != nil {
		comp = levelCnt(root.Right, comp+1)
	}

	return comp
}

package piscine

func BTreeLevelCount(root *TreeNode) int {
	var compL, compR int = 1, 1

	if root.Left != nil {
		compL = BTreeLevelCount(root.Left) + 1
	}
	if root.Right != nil {
		compR = BTreeLevelCount(root.Right) + 1
	}

	if compL > compR {
		return compL
	} else {
		return compR
	}
}

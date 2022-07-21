package piscine

func BTreeIsBinary(root *TreeNode) bool {
	var isBinary bool = false
	if root.Left != nil && root.Left.Data > root.Data {
		return false
	} else if root.Left != nil {
		isBinary = BTreeIsBinary(root.Left)
	} else {
		return true
	}
	if root.Right != nil && root.Right.Data <= root.Data {
		return false
	} else if root.Right != nil {
		isBinary = BTreeIsBinary(root.Right)
	} else {
		return true
	}
	return isBinary
}

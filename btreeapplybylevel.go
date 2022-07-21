package piscine

func BTreeApplyByLevel(root *TreeNode, f func(...interface{}) (int, error)) {
	if root == nil {
		return
	}
	if root.Parent == nil {
		f(root.Data)
	}

	if root.Left != nil {
		root.Left.Parent = root
		f(root.Left.Data)
	}
	if root.Right != nil {
		root.Right.Parent = root
		f(root.Right.Data)
	}
	BTreeApplyByLevel(root.Left, f)
	BTreeApplyByLevel(root.Right, f)
}

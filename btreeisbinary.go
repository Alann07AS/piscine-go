package piscine

func BTreeIsBinary(root *TreeNode) bool {
	var stat bool = true
	if root.Left != nil {
		if !(basicAtoi(root.Data) >= basicAtoi(root.Left.Data)) {
			return false
		} else {
			stat = BTreeIsBinary(root.Left)
		}
	}
	if root.Right != nil {
		if !(basicAtoi(root.Data) < basicAtoi(root.Right.Data)) {
			return false
		} else {
			stat = BTreeIsBinary(root.Right)
		}
	}
	return stat
}

func basicAtoi(s string) int {
	newInt := 0
	for _, each := range s {
		newInt = newInt*10 + (int(each - '0'))
	}
	return newInt
}

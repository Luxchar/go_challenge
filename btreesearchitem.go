package piscine

func BTreeSearchItem(root *TreeNode, elem string) *TreeNode {
	if root == elem || root == nil {
		return root
	} else if root.Data < elem {
		return BTreeSearchItem(root.Left, elem)
	}
	return BTreeSearchItem(root.Right, elem)
}

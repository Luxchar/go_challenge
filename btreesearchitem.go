package piscine

func BTreeSearchItem(root *TreeNode, elem string) *TreeNode {
	if root == elem || root == nil {
		return root
	} else if root.Data < elem {
		return BtreeSearchItem(root.Left, elem)
	}
	return BtreeSearchItem(root.Right, elem)
}

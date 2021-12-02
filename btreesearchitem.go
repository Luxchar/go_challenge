package piscine

type TreeNode struct {
	Left, Right, Parent *TreeNode
	Data                string
}

func BTreeSearchItem(root *TreeNode, elem string) *TreeNode {
	if root.Data == elem || root == nil {
		return root
	}
	if root.Data > elem {
		return BTreeSearchItem(root.Left, elem)
	}
	return BTreeSearchItem(root.Right, elem)
}

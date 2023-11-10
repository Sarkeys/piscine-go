package piscine

func BTreeDeleteNode(root, node *TreeNode) *TreeNode {
	if node.Left == nil && node.Right == nil {
		BTreeTransplant(root, node, nil)
	} else if node.Left == nil {
		BTreeTransplant(root, node, node.Right)
	} else if node.Right == nil {
		BTreeTransplant(root, node, node.Left)
	} else {
		y := BTreeMin(node.Right)
		if y != nil && y.Parent != node {
			root = BTreeTransplant(root, y, y.Right)
			y.Right = node.Right
			y.Right.Parent = y
		}
		root = BTreeTransplant(root, node, y)
		y.Left = node.Left
		y.Left.Parent = y
	}
	return root
}

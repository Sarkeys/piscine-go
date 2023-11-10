package piscine

func BTreeApplyByLevel(root *TreeNode, f func(...interface{}) (int, error)) {
	if root == nil {
		return
	}
	current := []*TreeNode{root}
	for len(current) > 0 {
		node := current[0]
		current = current[1:]
		f(node.Data)
		if node.Left != nil {
			current = append(current, node.Left)
		}
		if node.Right != nil {
			current = append(current, node.Right)
		}
	}
}

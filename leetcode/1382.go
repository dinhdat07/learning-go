package leetcode

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func balanceBST(root *TreeNode) *TreeNode {
	var arr []*TreeNode

	var inorder func(root *TreeNode)
	inorder = func(root *TreeNode) {
		if root == nil {
			return
		}
		inorder(root.Left)
		arr = append(arr, root)
		inorder(root.Right)
	}
	inorder(root)

	var build func([]*TreeNode) *TreeNode
	build = func(arr []*TreeNode) *TreeNode {
		if len(arr) == 0 {
			return nil
		}
		mid := len(arr) / 2
		node := arr[mid]
		node.Left = build(arr[:mid])
		node.Right = build(arr[mid+1:])
		return node
	}

	return build(arr)
}

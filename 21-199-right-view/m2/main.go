package main

type TreeNode struct {
	Val int
	Left *TreeNode 
	Right *TreeNode 
}

func rightSideView(root *TreeNode) []int {
	res := []int{}

	var helper func(*TreeNode, int) 
	helper = func(node *TreeNode, depth int) {
		if node == nil {
			return
		}

		if depth == len(res) {
			res = append(res, node.Val)
		}
		helper(node.Right, depth + 1)
		helper(node.Left, depth + 1)
	}

	helper(root, 0)
	return res 
}







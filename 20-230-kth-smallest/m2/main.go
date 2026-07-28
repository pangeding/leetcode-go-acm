package main

type TreeNode struct {
	Val int
	Left *TreeNode 
	Right *TreeNode 
}

func kthSmallest(root *TreeNode, k int) int {
	res := 0
	count := 0
	var inorder func(*TreeNode)
	inorder = func(node *TreeNode) {
		if node == nil {
			return 
		}

		inorder(node.Left)
		count++
		if count == k {
			res = node.Val 
		}
		inorder(node.Right)
	}
	inorder(root)
	return res
}



package main

type TreeNode struct {
	Val int
	Left *TreeNode 
	Right *TreeNode 
}


func isSymmetric(root *TreeNode) bool {
	if root == nil {
		return true 
	}
	return isMirror(root.Left, root.Right)
}

func isMirror(p, q *TreeNode) bool {
	if p == nil && q == nil {
		return true
	}

	if p == nil || q == nil {
		return false 
	}

	if p.Val != q.Val {
		return false 
	}

	return isMirror(p.Left, q.Right) && isMirror(p.Right, q.Left)
}

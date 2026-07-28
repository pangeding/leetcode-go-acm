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
	queue := []*TreeNode{root.Left, root.Right}
	for len(queue) > 0 {
		p := queue[0]
		q := queue[1]
		queue = queue[2:]
		if p == nil && q == nil {
			continue
		}
		if p == nil || q == nil {
			return false 
		}
		if p.Val != q.Val {
			return false
		} 
		queue = append(queue, p.Left, q.Right, p.Right, q.Left)
	}
	return true
}

package main

type TreeNode struct {
	Val int
	Left *TreeNode 
	Right *TreeNode 
}

func flatten(root *TreeNode) {
	if root == nil {
		return 
	}

	stack := []*TreeNode{root}
	prev := &TreeNode{}
	for len(stack) > 0 {
		node := stack[len(stack) - 1]
		stack = stack[:len(stack) - 1]

		prev.Right = node
		prev.Left = nil 

		if node.Right != nil {
			stack = append(stack, node.Right)
		}
		if node.Left != nil {
			stack = append(stack, node.Left)
		}
		
		prev = node
	}

}

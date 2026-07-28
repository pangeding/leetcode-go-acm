package main

type TreeNode struct {
	Val int
	Left *TreeNode 
	Right *TreeNode 
}

func postorderTraversal(root *TreeNode) []int {	
	// 第一个栈负责 根 右 左 
	// 第二个栈负责 反转
	if root == nil {
		return []int{} 
	}
	res := []int{}
	stack1 := []*TreeNode{root}
	stack2 := []*TreeNode{}
	for len(stack1) > 0 {
		node := stack1[len(stack1) - 1]
		stack1 = stack1[:len(stack1) - 1]
		stack2 = append(stack2, node)
		if node.Left != nil {
			stack1 = append(stack1, node.Left)
		}
		if node.Right != nil {
			stack1 = append(stack1, node.Right)
		}
	}

	for i := len(stack2) - 1; i >= 0; i-- {
		res = append(res, stack2[i].Val)
	}

	return res
}
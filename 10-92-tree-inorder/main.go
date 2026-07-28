package main

type TreeNode struct {
	Val int
	Left *TreeNode 
	Right *TreeNode 
}

func inorderTraversal(root *TreeNode) []int {
	result := make([]int, 0)
	inorder(root, &result)
	return result 
}

func inorder(node *TreeNode, result *[]int) {
	if node == nil {
		return 
	}
	inorder(node.Left, result)
	*result = append(*result, node.Val)
	inorder(node.Right, result)
}

/////

func inorderTraversalV2(root *TreeNode) []int {
	res := []int{}
	stack := []*TreeNode{}
	cur := root 

	for cur != nil || len(stack) > 0 {
		// 一直向左把压入栈
		for cur != nil {
			stack = append(stack, cur)
			cur = cur.Left 
		}

		// 弹出
		cur = stack[len(stack) - 1]
		stack = stack[:(len(stack) - 1)]
		res = append(res, cur.Val)

		// 转右
		cur = cur.Right
	}

	return res
}

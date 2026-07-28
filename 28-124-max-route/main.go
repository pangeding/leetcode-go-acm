package main

type TreeNode struct {
	Val int
	Left *TreeNode 
	Right *TreeNode 
}




func maxPathSum(root *TreeNode) int {
	// defien the maxSum
	// the smallest
	maxSum := -1 << 31 
	// dfs
	var dfs func(*TreeNode) int
	dfs = func(node *TreeNode) int {
		// corner case
		if node == nil {
			return 0
		}

		// left 
		leftGain := max(dfs(node.Left), 0)

		// right
		rightGain := max(dfs(node.Right), 0)

		// max cur max sum
		// only if you settle down you use this root as the start, 
		// can you both add left and right
		currentMax := node.Val + leftGain + rightGain
		maxSum = max(currentMax, maxSum)

		// or you can only choose one side
		return node.Val + max(leftGain, rightGain)
	}

	
 

	// revur
	dfs(root)

	// return 
	return maxSum 
	

	
}
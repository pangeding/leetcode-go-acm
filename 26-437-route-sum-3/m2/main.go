package main

type TreeNode struct {
	Val int
	Left *TreeNode 
	Right *TreeNode 
}

func pathSum(root *TreeNode, targetSum int) int {
	// define the hash map
	prefix := map[int]int{0: 1}
	
	// define the result 
	res := 0

	// dfs function
	var dfs func(node *TreeNode, curSum int) 
	dfs = func(node *TreeNode, curSum int) {
		// corner case 
		if node == nil {
			return
		}

		// curSum add now
		curSum += node.Val 

		// if the condition is satisfied, res++
		res += prefix[curSum - targetSum]

		// backtrace
		prefix[curSum]++
		dfs(node.Left, curSum)
		dfs(node.Right, curSum)
		prefix[curSum]--
	} 
	
	// start recursive 
	dfs(root, 0)

	// return 
	return res 
}
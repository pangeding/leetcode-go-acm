package main

type TreeNode struct {
	Val int
	Left *TreeNode 
	Right *TreeNode 
}

func pathSum(root *TreeNode, targetSum int) int {
	// corner case
	if root == nil {
		return 0
	}

	// start recursive

	// return dfs(root, targetSum)
	// this is wrong! not considering dfs is only the methodsNum root is the start. 
	// return dfs(root, targetSum) + pathSum(root.Left, targetSum - root.Val) + pathSum(root.Right, targetSum - root.Val)
	return dfs(root, targetSum) + pathSum(root.Left, targetSum) + pathSum(root.Right, targetSum)
}

/**
node as the root. how many methods include the targetSum 
*/
func dfs(node *TreeNode, targetSum int) int {
	// corner case 
	if node == nil {
		return 0
	}

	// define the "res"
	res := 0

	// if the condition is satisfied 
	if node.Val == targetSum {
		res++
	}

	// recursive 
	res += dfs(node.Left, targetSum - node.Val)
	res += dfs(node.Right, targetSum - node.Val)

	return res
}
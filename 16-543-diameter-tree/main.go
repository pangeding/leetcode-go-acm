package main

type TreeNode struct {
	Val int
	Left *TreeNode 
	Right *TreeNode 
}



func diameterOfBinaryTree(root *TreeNode) int {
	maxDiam := 0
    depth(root, &maxDiam)
	return maxDiam
}

func depth(node *TreeNode, maxDiam *int) int {
	if node == nil {
		return 0 
	}
	leftDepth := depth(node.Left, maxDiam)
	rightDepth := depth(node.Right, maxDiam)
	if leftDepth + rightDepth > *maxDiam {
		*maxDiam = leftDepth + rightDepth 
	}

	if leftDepth > rightDepth {
		return leftDepth + 1
	}

	return rightDepth + 1
 
}



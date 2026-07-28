package main

type TreeNode struct {
	Val int
	Left *TreeNode 
	Right *TreeNode 
}



func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
	// corner case return
	// root == nil
	// root == p
	// root == q
	if root == nil || root == p || root == q {
		return root 
	}
	
	// recursive 
	left := lowestCommonAncestor(root.Left, p, q)
	right := lowestCommonAncestor(root.Right, p, q)
	// if all have an ancestor, so we get up
	if left != nil && right != nil {
		return root
	}

	// if left have but right no return left
	if left != nil {
		return left
	}
	// return 
	return right 
	

}
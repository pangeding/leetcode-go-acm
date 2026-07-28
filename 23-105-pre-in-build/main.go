package main


type TreeNode struct {
	Val int
	Left *TreeNode 
	Right *TreeNode 
}

func buildTree(preorder []int, inorder []int) *TreeNode {
	idxMap := make(map[int]int)
	for i, v := range inorder {
		idxMap[v] = i
	}
	var build func(preStart, preEnd, inStart, inEnd int) *TreeNode 
	build = func(preStart, preEnd, inStart, inEnd int) *TreeNode {
		if preStart > preEnd || inStart > inEnd {
			return nil
		}
		rootVal := preorder[preStart]
		rootIdx := idxMap[rootVal]
		root := &TreeNode{Val: rootVal}
		leftSize := rootIdx - preStart
		root.Left= build(preStart + 1, preStart +leftSize, inStart, rootIdx - 1)
		root.Right = build(preStart + leftSize + 1, preEnd, rootIdx + 1, inEnd)
		return root
	}
	return build(0, len(preorder) - 1, 0, len(inorder) - 1)
}
package main



type TreeNode struct {
	Val int
	Left *TreeNode 
	Right *TreeNode 
}



func buildTree(inorder []int, postorder []int) *TreeNode {
	idxMap := make(map[int]int)
	for i, v := range inorder {
		idxMap[v] = i 
	}
	var build func(inStart, inEnd, postStart, postEnd int) *TreeNode
	build = func(inStart, inEnd, postStart, postEnd int) *TreeNode{
		if inStart > inEnd || postStart > postEnd {
			return nil
		}
		rootVal := postorder[postEnd]
		root := &TreeNode{Val: rootVal}
		rootIdx := idxMap[rootVal]
		leftSize := rootIdx - inStart 

		root.Left = build(inStart, rootIdx - 1, postStart, postStart + leftSize - 1)
		root.Right = build(rootIdx + 1, inEnd, postStart + leftSize, postEnd - 1)
		return root
	}

	return build(0, len(inorder) - 1, 0, len(postorder) - 1)
}
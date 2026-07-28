package main

type TreeNode struct {
	Val int
	Left *TreeNode 
	Right *TreeNode 
}


func constructFromPrePost(preorder []int, postorder []int) *TreeNode {
	postIdx := make(map[int]int)
	for i, v := range postorder {
		postIdx[v] = i 
	}


	var build func(preStart, preEnd, postStart, postEnd int) *TreeNode 
	build = func(preStart, preEnd, postStart, postEnd int) *TreeNode {
		if preStart > preEnd || postStart > postEnd {
			return nil 
		}
		rootVal := postorder[postEnd]
		// rootVal := preorder[preStart]
		// 上面 2 个 都可以的，没关系 一样的
		root := &TreeNode{Val: rootVal}

		if preStart == preEnd {
			return root
		}

		// 这个是错的，这个是最左边的数，不是左子树根leftRootVal := postorder[postStart + 1]
		leftRootVal := preorder[preStart + 1]
		leftRootIdx := postIdx[leftRootVal]
		leftSize := leftRootIdx - postStart + 1 
		
		root.Left = build(preStart + 1, preStart + leftSize, postStart, leftRootIdx)
		root.Right = build(preStart + leftSize + 1, preEnd, leftRootIdx + 1, postEnd - 1)
		return root
	}
	return build(0, len(preorder) - 1, 0, len(postorder) - 1)
}

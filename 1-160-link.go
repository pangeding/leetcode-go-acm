package main

import "fmt"

func main() {
	fmt.Println("Hello, World!")
}

func getIntersectionNode(headA, headB *ListNode) *ListNode {
	// 1. 如果 A 或 B 为空 责退出
	if headA == nil || headB == nil {
		return nil
	}

	// 2. 
	pA, pB := headA, headB
	for pA != pB {
		if pA == nil {
			pA = headB
		} else{
			pA = pA.Next
		}

		if pB == nil {
			pB = headA
		} else{
			pB = pB.Next
		}
	}

	return pA
}

type ListNode struct {
	Val int
	Next *ListNode
}
package main

// 1. sort merge 
type ListNode struct {
	Val int
	Next *ListNode
}

func sortList(head *ListNode) *ListNode {
	// 1. corner case 
	if head == nil || head.Next == nil {
		return head
	}

	// 2. split
	// 2.1 double pointer 
	slow, fast := head, head
	var prev *ListNode
	for fast != nil && fast.Next != nil {
		prev = slow
		slow = slow.Next 
		fast = fast.Next.Next 
	}
	prev.Next = nil

	// 2.2 recursive
	left := sortList(head)
	right := sortList(slow)

	// 3. return the merged list
	return merge(left, right)
}

func merge(l1 *ListNode, l2 *ListNode) *ListNode {
	// 0. dummy node
	dummy := &ListNode{}
	tail := dummy 

	// 1. compare
	for l1 != nil && l2 != nil {
		if l1.Val < l2.Val {
			tail.Next = l1 
			l1 = l1.Next 
		} else {
			tail.Next = l2 
			l2 = l2.Next 
		}
		tail = tail.Next 
	}

	// 2. the rest
	if l1 == nil {
		tail.Next = l2
	}

	if l2 == nil {
		tail.Next = l1
	}


	// 3. return
	return dummy.Next
}

// 2. 
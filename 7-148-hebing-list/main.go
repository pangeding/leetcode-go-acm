package main

// 1. sort merge 
type ListNode struct {
	Val int
	Next *ListNode
}

func sortListV1(head *ListNode) *ListNode {
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

// 2. 自顶向上归并排序

/**
规避了递归带来的 logn 空间
空间复杂度降低为 1
*/

/**
错误了？？改0为1？
*/

/**
unsolved
*/
func sortList(head *ListNode) *ListNode {
	// 1. corner case
	if head == nil || head.Next == nil {
		return head 
	}

	// 2. count the length of the List
	length := 0
	cur := head 
	for cur != nil {
		length++
		cur = cur.Next 
	}

	// 3. define the dummy node
	dummy := &ListNode{Next: head}

	// 4. cycle 
	for subLen := 1; subLen < length; subLen<<=1 {
		// 4.1 define the prev and cur
		prev := dummy 
		cur := dummy.Next 

		// 4.2 inner cycle 
		for cur != nil {
			// 4.3 find the head1
			head1 := cur 
			for i := 1; i < subLen && cur.Next != nil; i++ {
				cur = cur.Next 
			}

			// 4.4 find the head2
			head2 := cur.Next 
			// 4.5 the next of the cur is nil
			cur.Next = nil
			// 4.6 set the cur as head2
			cur = head2 

			// 4.7 移动 cur 到第二段末尾
			for i := 1; i < subLen && cur != nil && cur.Next != nil; i++ {
				cur = cur.Next
			}

			var next *ListNode 
			if cur != nil {
				next = cur.Next 
				cur.Next = nil // 断开第二段
			}

			// 合并 head1 和 head2
			merged := merge1(head1, head2) 
			// 连接到已经排序的部分
			prev.Next = merged

			// 移动 prev 到当前合并段的末尾，准备下一轮连接
			for prev.Next != nil {
				prev = prev.Next 
			}
			cur = next 


		}
	}
	return dummy.Next
}

func merge1(l1 *ListNode, l2 *ListNode) *ListNode {
	// 0. corner
	dummy := &ListNode{}
	cur := dummy

	// 1. merge
	for l1 != nil && l2 != nil {
		if l1.Val < l2.Val {
			cur.Next = l1 
			l1 = l1.Next 
		} else {
			cur.Next = l2 
			l2 = l2.Next
		}
		cur = cur.Next
	}

	// 2. tail
	if l1 != nil {
		cur.Next = l1
	}

	if l2 != nil {
		cur.Next = l2
	}

	// 3. return 
	return dummy.Next
}
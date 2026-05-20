package main

/** 
2 mistakes:

first: tail == nil then we should return, if we return at tail != nil ,we just randomly return

secondly, we should use reverse(head, tail.Next) because the reverse function cope with the [) interval

thirdly when continuing the loop, we should not use prev = nextHead, head = nextHead.Next. we should use prev = head, 
head = nextHead
*/
type ListNode struct {
	Val int
	Next *ListNode
}

func reverse(head *ListNode, tail *ListNode) *ListNode {
	var prev *ListNode
	cur := head
	for cur != tail {
		// the reverse logic
		next := cur.Next
		cur.Next = prev
		// after the reverse, the position should be switched
		prev = cur
		cur = next
	}

	return prev
}

func reverseKGroup(head *ListNode, k int) *ListNode {
	// head == nil or head.Next == nil or 
	if head == nil || k <= 1 {
		return head
	}

	// dummy
	dummy := &ListNode{Next: head}
	prev := dummy	

	for head != nil {
		tail := prev
		// reverse
		// find the tail
		for i := 0; i < k; i++ {
			tail = tail.Next
			if tail == nil {
				return dummy.Next
			}
		}
		// if fail to find the tail, return 

		// if can find the tail, continue the reverse logic

		// mark the next prev
		nextHead := tail.Next

		// reverse
		newHead := reverse(head, tail.Next)

		// reconnect the linkList
		prev.Next = newHead
		head.Next = nextHead

		// continue the next loop
		prev = head
		head = nextHead
	}
	

	// return
	return dummy.Next
}

func main() {
	
}
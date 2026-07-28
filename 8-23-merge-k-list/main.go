package main
import "container/heap"

type ListNode struct {
	Val int 
	Next *ListNode
}

type NodeHeap []*ListNode

func (h NodeHeap) Len() int {
	return len(h)
}

func (h NodeHeap) Less(i, j int) bool {
	return h[i].Val < h[j].Val 
}

func (h NodeHeap) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}

func (h *NodeHeap) Push(x interface{}) {
	*h = append(*h, x.(*ListNode))
}

func (h *NodeHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n - 1]
	*h = old[0 : n - 1]
	return x 
}

func mergeKLists(lists []*ListNode) *ListNode {
	if len(lists) == 0 {
		return nil 
	}
	h := &NodeHeap{} 
	heap.Init(h)

	for _, head := range lists {
		if head != nil {
			heap.Push(h, head)
		}
	}

	dummy := &ListNode{}
	cur := dummy 

	for h.Len() > 0 {
		node := heap.Pop(h).(*ListNode) 
		cur.Next = node
		cur = cur.Next 

		if node.Next != nil {
			heap.Push(h, node.Next)
		}
	}

	return dummy.Next 
}



/////////////////////////////////////////

/** 
分治
*/

func mergeKListsV2(lists []*ListNode) *ListNode {
	if len(lists) == 0 {
		return nil
	}
	return merge(lists, 0, len(lists) - 1)
}

func merge(lists []*ListNode, left, right int) *ListNode {
	if left == right {
		return lists[left]
	}
	mid := left + (right - left) / 2
	l1 := merge(lists, left, mid)
	l2 := merge(lists, mid + 1, right)
	return mergeTwoLists(l1, l2)
}

func mergeTwoLists(l1, l2 *ListNode) *ListNode {
	if l1 == nil && l2 == nil {
		return nil 
	}

	dummy := &ListNode{}
	tail := dummy
	
	for l1 != nil && l2 != nil {
		if l1.Val > l2.Val {
			tail.Next = l2 
			l2 = l2.Next
		} else {
			tail.Next = l1
			l1 = l1.Next
		}
		tail = tail.Next 
	}

	if l1 != nil {
		tail.Next = l1
	}
	if l2 != nil {
		tail.Next = l2
	}

	return dummy.Next
}








package swapNode

import (
	"fmt"
)

/**
* note: what is append?
* built-in function
* it can append 1 or many elements to the slice
*/



type ListNode struct {
	Val int
	Next *ListNode
}

func swapPairs(head *ListNode) *ListNode {
	dummy := &ListNode{Next: head}
	prev := dummy

	for prev.Next != nil && prev.Next.Next != nil {
		first := prev.Next
		second := prev.Next.Next

		// change
		prev.Next = second
		first.Next = second.Next
		second.Next = first

		prev = first
	}
	// return 
	return dummy.Next
}

func buildList(nums []int) *ListNode {
	return nil
}

func printList(head *ListNode) {

}

func main() {
	var n int
	fmt.Scan(&n) // read the number of the node
	// only 大写 开头函数可以被外部类访问 go 语言中

	var nums []int
	for i := 0; i < n; i++ {
		var val int
		fmt.Scan(&val)
		nums = append(nums, val)
	}
	
	head := buildList(nums)
	newHead := swapPairs(head)
	printList(newHead)
}
package reverse

type ListNode struct {
	Val int
	Next *ListNode
}
/**
func reverseList(head *ListNode) *ListNode{
	// 0. 迭代

	// 1. 定义 prev (nil)
	var prev *ListNode
	curr := head
	for curr != nil {
		// 2.0 备份
		nextTemp := curr.Next
		// 2.1 交换
		curr.Next = prev
		// 2.2 迭代
		prev = curr
		curr = nextTemp
	}

	return prev
}
*/

func reverseList(head *ListNode) *ListNode{
	// 0. 递归
	// 1. 递归边界情况
	if head == nil || head.Next == nil {
		return head
	}
	// 2. 开始递归
	newHead :=reverseList(head.Next)
	// 3. 交换
	// 3.1 实施交换
	head.Next.Next = head
	// 3.2 避免循环
	head.Next = nil
	return newHead
}



func main() {

}
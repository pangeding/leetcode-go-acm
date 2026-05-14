
```go
func reverseList(head *ListNode) *ListNode{
	// 0. 递归
	// 1. 递归边界情况
	if head == nil || head.Next == nil {
		return head // 返回的是原来的尾节点，后面变成了头结点
	}
	// 2. 开始递归
	reverseList(head.Next) // 没有保存递归后的节点
	// 3. 交换
	// 3.1 实施交换
	head.Next.Next = head
	// 3.2 避免循环
	head.Next = nil
	return head // 返回原来的头结点，变成了现在的尾结点
}
```

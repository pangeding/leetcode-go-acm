package main

import (
	"fmt";
	"strconv";
	"strings";

	"bufio";
	"os";
)
/** 
3 mistakes:

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

func buildList(nums []int) *ListNode {
	if len(nums) == 0 {
		return nil
	}
	// since the 0 doesn't necessarily 
	head := &ListNode{Val: nums[0]}
	cur := head
	for i := 1; i < len(nums); i++ {
		cur.Next = &ListNode{Val: nums[i]}
		cur = cur.Next
	}
	return head
}

/**
 * printList vs printListV2 的区别：
 * 
 * 1. 切片初始化方式不同：
 *    - printList:   var out []string          // 声明为 nil 切片
 *    - printListV2: out := make([]string, 0)  // 创建长度为 0 的空切片
 * 
 * 4. 最佳实践：
 *    - 推荐 printListV2 的写法（make([]string, 0)）
 *    - 更明确地表达"需要一个空切片"的意图
 *    - 避免 nil 切片在某些边界情况下的意外行为
 */
func printList(head *ListNode) {
	var out []string
	for head != nil {
		out = append(out, strconv.Itoa(head.Val))
		head = head.Next
	}
	fmt.Println(strings.Join(out, " "))
}

func printListV2(head *ListNode) {
	out := make([]string, 0)
	for head != nil {
		out = append(out, strconv.Itoa(head.Val))
		head = head.Next
	}
	fmt.Println(strings.Join(out, " "))
}

func usingScan() {
	var n int
	fmt.Scan(&n)

	// 先读取链表元素
	var nums []int
	for i := 0; i < n; i++ {
		var val int
		fmt.Scan(&val)
		nums = append(nums, val)
	}

	// 再读取 k 值
	var k int
	fmt.Scan(&k)

	head := buildList(nums)
	newHead := reverseKGroup(head, k)
	printList(newHead)
}

func usingBuffer() {
	// 这2行代码都是什么意思？
	// 为什么Scanwords就可以表示空格，也没说啊
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Split(bufio.ScanWords)

	// 读取节点个数
	scanner.Scan()
	// 这种什么意思？
	n, _ := strconv.Atoi(scanner.Text())

	// 读取链表元素
	nums := make([]int, n)
	for i := 0; i < n; i++ {
		scanner.Scan()
		nums[i], _ = strconv.Atoi(scanner.Text())
	}

	// 读取 k 值
	scanner.Scan()
	k, _ := strconv.Atoi(scanner.Text()) 

	head := buildList(nums)
	newHead := reverseKGroupV2(head, k)
	printList(newHead)
}

func reverseV2(head *ListNode, tail *ListNode) *ListNode {
	var prev *ListNode
	// cur
	cur := head
	// start the change
	for cur != tail {
		next := cur.Next
		cur.Next = prev
		prev = cur
		cur = next
	}
	// return
	return prev
}

func reverseKGroupV2(head *ListNode, k int) *ListNode {
	// 边界情况
	if head == nil {
		return head
	}

	// no need to define the dummy, since it is 迭代
	cur := head

	// 看还有没有k个节点
	for i := 0; i < k; i++ {
			// cur 可能先
			/**cur = cur.Next
			if cur == nil {
				return head
			}
				it is not fair. 最后一个明明符合要求还被淘汰掉了
			*/
		if cur == nil {
			return head
		}
		cur = cur.Next
	}

	// 没有k个节点就返回

	// 有 k 个节点就递归解决
	newHead := reverseV2(head, cur)

	// 进入下一个流程
	head.Next = reverseKGroupV2(cur, k)

	// return
	return newHead
}
func main() {
	// usingScan()
	usingBuffer()
}
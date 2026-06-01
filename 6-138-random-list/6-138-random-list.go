package main

import (
	"os";
	"bufio"
)


type Node struct {
	Val int
	Next *Node
	Random *Node
}

func copyRandomList(head *Node) *Node {
	if head == nil {
		return nil
	}

	// the map
	m := make(map[*Node]*Node)

	// the first copy
	cur := head 
	for cur != nil {
		m[cur] = &Node{Val: cur.Val}
		cur = cur.Next
	}

	// the second copy: copy the Next and Random
	cur = head 
	for cur != nil {
		m[cur].Next = m[cur.Next]
		m[cur].Random = m[cur.Random]
		cur = cur.Next
	}

	return m[head]
}

// why it is wrong? time out of limit
func copyRandomListV2(head *Node) *Node {
	if head == nil {
		return nil 
	}
	// put the node into the list
	cur := head 
	for cur != nil {
		copy := &Node{Val: cur.Val}
		copy.Next = cur.Next
		cur.Next = copy
		// cur = cur.Next 
		// use cur = cur.Next 就在同一个节点后面不断插入新的节点，所以不行
		cur = copy.Next
	}

	// add the Next and Random relation
	cur = head
	for cur != nil {
		if cur.Random != nil {
			cur.Next.Random = cur.Random.Next
		}
		cur = cur.Next.Next
	}

	// put out the things
	newHead := head.Next 
	cur = head
	for cur != nil {
		copy := cur.Next 
		cur.Next = copy.Next 
		if copy.Next != nil {
			copy.Next = copy.Next.Next 
		}
		cur = cur.Next 
	}

	return newHead 
}
func main() {
	// to complicated, next cycle
	scanner := bufio.NewScanner(os.Stdin)

}

/**
copy.Next = copy.Next.Next

panic: runtime error: invalid memory address or nil pointer dereference
[signal SIGSEGV: segmentation violation code=0x1 addr=0x8 pc=0x4c153a]
main.copyRandomList(0xc000012018)
solution.go, line 30
main.__helper__(...)
solution.go, line 193
main.main()
solution.go, line 229
*/
func copyRandomListV3False(head *Node) *Node {
	// 原地
	// 0. corner case

	// 1. add val
	cur := head
	for cur != nil {
		copy := &Node{Val: cur.Val}
		copy.Next = cur.Next
		cur.Next = copy
		cur = copy.Next
	}

	// 2. add the random 
	cur = head 
	for cur != nil {
		if cur.Random != nil {
			cur.Next.Random = cur.Random.Next
		}
		cur = cur.Next.Next
	}

	// 3. 抓出来 构建next
	newHead := head.Next
	cur = head 
	for cur != nil {
		// ???
		copy := cur.Next
		cur.Next = copy.Next
		// this line is wrong!!!
		copy.Next = copy.Next.Next
		cur = cur.Next
	}

	// 4. return
	return newHead
}

/**
panic: runtime error: invalid memory address or nil pointer dereference
[signal SIGSEGV: segmentation violation code=0x1 addr=0x8 pc=0x4c14fa]
main.copyRandomList(0x0)
solution.go, line 24
main.__helper__(...)
solution.go, line 194
main.main()
solution.go, line 230

null pointer exception
*/
func copyRandomListV4(head *Node) *Node {
	// 原地
	// 0. corner case

	// 1. add val
	cur := head
	for cur != nil {
		copy := &Node{Val: cur.Val}
		copy.Next = cur.Next
		cur.Next = copy
		cur = copy.Next
	}

	// 2. add the random 
	cur = head 
	for cur != nil {
		if cur.Random != nil {
			cur.Next.Random = cur.Random.Next
		}
		cur = cur.Next.Next
	}

	// 3. 抓出来 构建next
	newHead := head.Next
	cur = head 
	for cur != nil {
		copy := cur.Next
		cur.Next = copy.Next
		if copy.Next != nil {
			copy.Next = copy.Next.Next
		}
		cur = cur.Next
	}

	// 4. return
	return newHead
}



func copyRandomListV5(head *Node) *Node {
	// 原地
	// 0. corner case
	if head == nil {
		return head
	}

	// 1. add val
	cur := head
	for cur != nil {
		copy := &Node{Val: cur.Val}
		copy.Next = cur.Next
		cur.Next = copy
		cur = copy.Next
	}

	// 2. add the random 
	cur = head 
	for cur != nil {
		if cur.Random != nil {
			cur.Next.Random = cur.Random.Next
		}
		cur = cur.Next.Next
	}

	// 3. 抓出来 构建next
	newHead := head.Next
	cur = head 
	for cur != nil {
		copy := cur.Next
		cur.Next = copy.Next
		if copy.Next != nil {
			copy.Next = copy.Next.Next
		}
		cur = cur.Next
	}

	// 4. return
	return newHead
}



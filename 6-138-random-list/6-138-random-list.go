package main

import (
	"os";
	"bufio";
	"strconv";
	"strings";
	"fmt";

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

func buildList(n int, data [][2]int) *Node {
	// 0. corner case
	if n == 0 {
		return nil
	}

	// 0.5 where to store the nodes
	nodes := make([]*Node, n)

	// 1. val
	for i := 0; i < n; i++ {
		nodes[i] = &Node{Val: data[i][0]}
	}

	// 2. Next 
	for i := 0; i < n - 1; i++ {
		nodes[i].Next = nodes[i + 1]
	}

	for i := 0; i < n; i++ {
		if data[i][1] != -1 {
			nodes[i].Random = nodes[data[i][1]]
		}
	}

	// 3. Random 

	// 4. return
	return nodes[0]
}

func printList(head *Node) {
	// 0. define 
	out := []string{}
	cur := head 

	for cur != nil {
		var randomVal string
		if cur.Random == nil {
			randomVal = "null"
		}else {
			randomVal = strconv.Itoa(cur.Random.Val) 
		} 		
		out = append(out, fmt.Sprintf("[%d, %s]", cur.Val, randomVal))
		cur = cur.Next
	}

	fmt.Println(strings.Join(out, " "))
}
func main() {
	// to complicated, next cycle
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Split(bufio.ScanWords)

	// 读取节点个数
	scanner.Scan()
	n, _ := strconv.Atoi(scanner.Text())

	data := make([][2]int, n)

	for i := 0; i < n; i++ {
		scanner.Scan()
		val, _ := strconv.Atoi(scanner.Text())

		scanner.Scan()
		randomIdx, _ := strconv.Atoi(scanner.Text())

		data[i] = [2]int{val, randomIdx}
	}

	original := buildList(n, data)
	copied := copyRandomList(original)
	printList(copied)


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



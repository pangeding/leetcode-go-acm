package main


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

}


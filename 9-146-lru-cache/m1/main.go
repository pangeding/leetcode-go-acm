package main



type Node struct {
	key, value int
	prev, next *Node
}

type LRUCache struct {
	capacity int
	cache map[int]*Node
	head *Node
	tail *Node
}

func Constructor(capacity int) LRUCache {
	head := &Node{}
	tail := &Node{}
	head.next = tail
	tail.prev = head 

	return LRUCache{
		capacity: capacity,
		cache: make(map[int]*Node),
		head: head,
		tail: tail,
	}
}

func (this *LRUCache) Get(key int) int {
	if node, ok := this.cache[key]; ok {
		this.moveToHead(node)
		return node.value 
	}
	return -1
}

func (this *LRUCache) Put(key int, value int) {
	if node, ok := this.cache[key]; ok {
		node.value = value 
		this.moveToHead(node)
		return
	}


	// 容量检查 如果超出就淘汰尾部节点
	if len(this.cache) >= this.capacity {
		removed := this.removeTail() 
		delete(this.cache, removed.key) 
	}

	newNode := &Node{key: key, value: value}
	this.cache[key] = newNode 
	this.addToHead(newNode)
}



////////

// 双向链表辅助

func (this *LRUCache) addToHead(node *Node) {
	// 1. 将 这个节点 node 加到链路里

	// 1.1 先确定 node 的 next
	node.next = this.head.next

	// 1.2 然后确定 node 的 prev
	node.prev = this.head

	// 2. 处理一些节点之间的连接关系
	// 为什么要先 2.1 因为 head.next 如果先变换，head.next.prev 就丢失了

	// 2.1 头结点下一个的 prev 变成 node
	this.head.next.prev = node

	// 2.2 头结点的下一个（next）变成 node
	this.head.next = node
}

func (this *LRUCache) removeNode(node *Node) {
	node.prev.next = node.next
	node.next.prev = node.prev
}

func (this *LRUCache) moveToHead(node *Node) {
	this.removeNode(node)
	this.addToHead(node)
}

func (this *LRUCache) removeTail() *Node {
	node := this.tail.prev
	this.removeNode(node)
	return node
}




//////


// 生产环境




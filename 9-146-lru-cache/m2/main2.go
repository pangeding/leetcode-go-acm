package main

import "container/list"

type entry struct {
	key int
	value int
}

type LRUCache struct {
	capacity int
	cache map[int] *list.Element
	list *list.List 
}

func Constructor(capacity int) LRUCache {
	return LRUCache{
		capacity: capacity,
		cache: make(map[int]*list.Element),
		list: list.New(),
	}
}

func (this *LRUCache) Get(key int) int {
	if elem, ok := this.cache[key]; ok {
		this.list.MoveToFront(elem)
		return elem.Value.(*entry).value
	}
	return -1 
}

func (this *LRUCache) Put(key int, value int) {
	if elem, ok := this.cache[key]; ok {
		elem.Value.(*entry).value = value
		this.list.MoveToFront(elem)
		return 
	}

	if this.list.Len() >= this.capacity {
		oldest := this.list.Back() 
		if oldest != nil {
			delete(this.cache, oldest.Value.(*entry).key)
			this.list.Remove(oldest)
		}
	}

	newElem := this.list.PushFront(&entry{key: key, value: value})
	this.cache[key] = newElem
 }



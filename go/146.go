package leetcode

type LRUCache struct {
	size       int
	capacity   int
	cache      map[int]*DLinkedNode
	head, tail *DLinkedNode
}

type DLinkedNode struct {
	key, value int
	prev, next *DLinkedNode
}

func dln(key, value int) *DLinkedNode {
	return &DLinkedNode{
		key:   key,
		value: value,
	}
}

func Constructor(capacity int) LRUCache {
	lru := LRUCache{
		cache:    map[int]*DLinkedNode{},
		head:     dln(0, 0),
		tail:     dln(0, 0),
		capacity: capacity,
	}
	lru.head.next = lru.tail
	lru.tail.prev = lru.head
	return lru
}

func (this *LRUCache) Get(key int) int {
	if _, ok := this.cache[key]; !ok {
		return -1
	}
	node := this.cache[key]
	node.prev.next = node.next
	node.next.prev = node.prev
	node.prev = this.head
	node.next = this.head.next
	this.head.next.prev = node
	this.head.next = node
	return node.value
}

func (this *LRUCache) Put(key int, value int) {
	if _, ok := this.cache[key]; !ok {
		node := dln(key, value)
		this.cache[key] = node
		node.prev = this.head
		node.next = this.head.next
		this.head.next.prev = node
		this.head.next = node
		this.size++
		if this.size > this.capacity {
			del := this.tail.prev
			del.prev.next = del.next
			del.next.prev = del.prev
			delete(this.cache, del.key)
			this.size--
		}
	} else {
		node := this.cache[key]
		node.value = value
		node.prev.next = node.next
		node.next.prev = node.prev
		node.prev = this.head
		node.next = this.head.next
		this.head.next.prev = node
		this.head.next = node
	}
}

/**
 * Your LRUCache object will be instantiated and called as such:
 * obj := Constructor(capacity);
 * param_1 := obj.Get(key);
 * obj.Put(key,value);
 */

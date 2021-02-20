package main

import "fmt"

func main() {
	capacity := 2
	obj := Constructor(capacity)

	// fmt.Println(obj.Get(1))
	obj.Put(2, 1)
	obj.Print()
	obj.Put(2, 2)
	obj.Print()
	// obj.Print()
	fmt.Println(obj.Get(2))
	obj.Put(1, 1)
	obj.Print()
	obj.Put(4, 1)
	// fmt.Println(obj.Get(2))
	// obj.Print()
	// obj.Put(3, 3)
	// obj.Print()
	// fmt.Println(obj.Get(2))
	// obj.Print()
	// fmt.Println(obj.Get(2))
	// obj.Print()
	// obj.Put(4, 1)
	// obj.Print()
	// fmt.Println(obj.Get(2))
}

type DLinkedNode struct {
	key, val   int
	prev, next *DLinkedNode
}

type LRUCache struct {
	size     int
	capacity int
	head     *DLinkedNode
	tail     *DLinkedNode
	cache    map[int]*DLinkedNode
}

func initDLinkedNode(key, val int) *DLinkedNode {
	return &DLinkedNode{
		key: key,
		val: val,
	}
}

func Constructor(capacity int) LRUCache {
	l := LRUCache{
		size:     0,
		capacity: capacity,
		cache:    map[int]*DLinkedNode{},
		head:     initDLinkedNode(0, 0),
		tail:     initDLinkedNode(0, 0),
	}
	l.head.next = l.tail
	l.tail.prev = l.head
	return l
}

func (this *LRUCache) addToHead(node *DLinkedNode) {
	node.prev = this.head
	node.next = this.head.next
	this.head.next.prev = node
	this.head.next = node
}

func (this *LRUCache) removeNode(node *DLinkedNode) {
	node.prev.next = node.next
	node.next.prev = node.prev
}

func (this *LRUCache) moveToHead(node *DLinkedNode) {
	this.removeNode(node)
	this.addToHead(node)
}

func (this *LRUCache) removeTail() *DLinkedNode {
	node := this.tail.prev
	this.removeNode(node)
	return node
}

func (this *LRUCache) Get(key int) int {
	if _, ok := this.cache[key]; !ok {
		return -1
	}
	node := this.cache[key]

	this.moveToHead(node)
	return node.val
}

func (this *LRUCache) Print() {
	for p := this.head.next; p != nil; p = p.next {
		fmt.Printf("%d，%d->", p.key, p.val)
	}
	fmt.Println()
	fmt.Println(this.size)
}

func (this *LRUCache) Put(key int, value int) {
	newNode := initDLinkedNode(key, value)
	if item, ok := this.cache[key]; !ok {
		this.addToHead(newNode)
		this.size++
		this.cache[key] = newNode
		if this.size > this.capacity {
			removed := this.removeTail()
			delete(this.cache, removed.key)
			this.size--
		}
		this.cache[key] = newNode
	} else {
		this.cache[key] = newNode
		this.removeNode(item)
		this.addToHead(newNode)
	}
}

package main

import "fmt"

func main() {
	capacity := 2
	obj := Constructor(capacity)

	obj.Put(1, 1)
	obj.Put(2, 2)
	fmt.Printf("%d\r\n", obj.Get(1))
	obj.Put(3, 3)
	fmt.Printf("%d\r\n", obj.Get(2))
	obj.Put(4, 4)
	fmt.Printf("%d\r\n", obj.Get(1))
	fmt.Printf("%d\r\n", obj.Get(3))
	fmt.Printf("%d\r\n", obj.Get(4))
}

type LRUCache struct {
	capacity int
	cache    map[int]*DequeueNode
	head     *DequeueNode
	tail     *DequeueNode
}

type DequeueNode struct {
	Prev *DequeueNode
	Next *DequeueNode
	Key  int
	Val  int
}

func NewDequeue(key, val int) *DequeueNode {
	return &DequeueNode{nil, nil, key, val}
}

func Constructor(capacity int) LRUCache {
	l := LRUCache{
		capacity: capacity,
		cache:    map[int]*DequeueNode{},
		head:     NewDequeue(0, 0),
		tail:     NewDequeue(0, 0),
	}
	l.head.Next = l.tail
	l.tail.Prev = l.head
	return l
}

func (this *LRUCache) Get(key int) int {
	item, ok := this.cache[key]
	if !ok {
		return -1
	}

	this.removeNode(item)
	this.addToHead(item)

	return item.Val
}

func (this *LRUCache) Put(key int, value int) {
	var (
		item *DequeueNode
		ok   bool
	)
	item, ok = this.cache[key]
	if ok {
		this.removeNode(item)
	}
	item = NewDequeue(key, value)
	this.addToHead(item)
	this.cache[key] = item

	if len(this.cache) > this.capacity {
		node := this.removeTail()
		delete(this.cache, node.Key)
	}
}
func (this *LRUCache) removeNode(node *DequeueNode) {
	node.Prev.Next = node.Next
	node.Next.Prev = node.Prev
}

func (this *LRUCache) addToHead(node *DequeueNode) {
	node.Next = this.head.Next
	node.Next.Prev = node
	node.Prev = this.head
	this.head.Next = node
}

func (this *LRUCache) removeTail() *DequeueNode {
	node := this.tail.Prev
	node.Prev.Next = this.tail
	this.tail.Prev = node.Prev
	return node
}

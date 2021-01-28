package main

import (
	"fmt"
)

type Stack []int

func (s *Stack) Len() int {
	return len(*s)
}

func (s *Stack) Push(v int) {
	*s = append(*s, v)
}

func (s *Stack) Pop() (v int) {
	*s, v = (*s)[:len(*s)-1], (*s)[len(*s)-1]
	return
}

type CQueue struct {
	in  *Stack
	out *Stack
}

func main() {
	obj := Constructor()
	obj.AppendTail(3)
	obj.AppendTail(5)
	obj.AppendTail(2)
	param_2 := obj.DeleteHead()
	fmt.Println(param_2)
	param_2 = obj.DeleteHead()
	fmt.Println(param_2)
	param_2 = obj.DeleteHead()
	fmt.Println(param_2)
}

func Constructor() CQueue {
	return CQueue{
		in:  new(Stack),
		out: new(Stack),
	}
}

func (this *CQueue) AppendTail(value int) {
	this.in.Push(value)
}

func (this *CQueue) DeleteHead() int {
	if this.out.Len() == 0 {
		for this.in.Len() != 0 {
			item := this.in.Pop()
			this.out.Push(item)
		}
	}
	if this.out.Len() == 0 {
		return -1
	}
	item := this.out.Pop()
	return item
}

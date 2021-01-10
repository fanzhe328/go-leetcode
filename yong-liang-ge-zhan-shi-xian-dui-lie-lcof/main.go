package main

import (
	"container/list"
	"fmt"
)

type CQueue struct {
	In  *list.List
	Out *list.List
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
		In:  list.New(),
		Out: list.New(),
	}
}

func (this *CQueue) AppendTail(value int) {
	this.In.PushBack(value)
}

func (this *CQueue) DeleteHead() int {
	tmp := this.Out.Back()
	if tmp == nil {
		for {
			item := this.In.Back()
			if item == nil {
				break
			}
			this.Out.PushBack(item.Value.(int))
			this.In.Remove(item)
		}
	}
	tmp = this.Out.Back()
	if tmp == nil {
		return -1
	}
	this.Out.Remove(tmp)
	return tmp.Value.(int)
}

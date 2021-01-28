package main

import "fmt"

type MinStack struct {
	list       []int
	minIdxList []int
}

func main() {
	// fmt.Println(removeDuplicates([]int{0, 0, 1, 1, 1, 2, 2, 3, 3, 4}))
	obj := Constructor()
	obj.Push(-2)
	obj.Push(0)
	obj.Push(-3)
	fmt.Println(obj.Min())
	fmt.Println(obj.Top())
	obj.Pop()
	fmt.Println(obj.Min())
	obj.Pop()
	fmt.Println(obj.Min())
	obj.Pop()
	fmt.Println(obj.Min())
}

/** initialize your data structure here. */
func Constructor() MinStack {
	return MinStack{
		list:       []int{},
		minIdxList: []int{},
	}
}

func (this *MinStack) Push(x int) {
	this.list = append(this.list, x)
	if len(this.list) == 1 {
		this.minIdxList = append(this.minIdxList, 0)
		return
	}
	if len(this.minIdxList) == 0 || this.list[this.minIdxList[len(this.minIdxList)-1]] > x {
		this.minIdxList = append(this.minIdxList, len(this.list)-1)
	}
	// fmt.Println(this.minIdxList)
}

func (this *MinStack) Pop() {
	idx := len(this.list) - 1
	if this.minIdxList[len(this.minIdxList)-1] == idx {
		this.minIdxList = this.minIdxList[:len(this.minIdxList)-1]
	}
	this.list = this.list[:len(this.list)-1]
	// fmt.Println(this.minIdxList)
}

func (this *MinStack) Top() int {
	return this.list[len(this.list)-1]
}

func (this *MinStack) Min() int {
	if len(this.minIdxList) == 0 {
		return 0
	}
	return this.list[this.minIdxList[len(this.minIdxList)-1]]
}

/**
 * Your MinStack object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(x);
 * obj.Pop();
 * param_3 := obj.Top();
 * param_4 := obj.Min();
 */

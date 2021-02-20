package main

func main() {

}

type MaxQueue struct {
	q   []int
	max []int
}

func Constructor() MaxQueue {
	return MaxQueue{
		q:   []int{},
		max: []int{},
	}
}

func (this *MaxQueue) Max_value() int {
	if len(this.max) == 0 {
		return -1
	}
	return this.max[0]
}

func (this *MaxQueue) Push_back(value int) {
	this.q = append(this.q, value)
	idx := len(this.max) - 1
	for ; idx >= 0; idx-- {
		if this.max[idx] >= value {
			break
		}
	}
	this.max = append(this.max[:idx+1], value)
}

func (this *MaxQueue) Pop_front() int {
	if len(this.q) == 0 {
		return -1
	}
	var v int
	this.q, v = this.q[1:], this.q[0]
	if v == this.max[0] {
		this.max = this.max[1:]
	}
	return v
}

/**
 * Your MaxQueue object will be instantiated and called as such:
 * obj := Constructor();
 * param_1 := obj.Max_value();
 * obj.Push_back(value);
 * param_3 := obj.Pop_front();
 */

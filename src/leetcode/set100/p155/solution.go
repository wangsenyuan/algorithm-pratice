package p155

type MinStack struct {
	stack []int
	min   []int
	p     int
}

/** initialize your data structure here. */
func Constructor() MinStack {
	return MinStack{stack: make([]int, 1), min: make([]int, 10), p: 0}
}

func (this *MinStack) Push(x int) {
	if this.p == len(this.stack) {
		ext := make([]int, this.p+100)
		copy(ext, this.stack)
		this.stack = ext
		ext = make([]int, this.p+100)
		copy(ext, this.min)
		this.min = ext
	}
	this.stack[this.p] = x
	if this.p > 0 && this.min[this.p-1] < x {
		this.min[this.p] = this.min[this.p-1]
	} else {
		this.min[this.p] = x
	}
	this.p++
}

func (this *MinStack) Pop() {
	this.p--
}

func (this *MinStack) Top() int {
	return this.stack[this.p-1]
}

func (this *MinStack) GetMin() int {
	return this.min[this.p-1]
}

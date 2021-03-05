package p232

type MyQueue struct {
	stack1 []int
	stack2 []int
}

/** Initialize your data structure here. */
func Constructor() MyQueue {
	stack1 := make([]int, 0, 1)
	stack2 := make([]int, 0, 1)
	return MyQueue{stack1, stack2}
}

/** Push element x to the back of queue. */
func (this *MyQueue) Push(x int) {
	this.stack1 = append(this.stack1, x)
}

/** Removes the element from in front of queue and returns that element. */
func (this *MyQueue) Pop() int {
	n := len(this.stack1)
	for n > 1 {
		x := this.stack1[n-1]
		this.stack2 = append(this.stack2, x)
		n--
	}
	res := this.stack1[n-1]
	this.stack1 = this.stack1[:0]

	n = len(this.stack2)
	for n > 0 {
		this.stack1 = append(this.stack1, this.stack2[n-1])
		n--
	}
	this.stack2 = this.stack2[:0]

	return res
}

/** Get the front element. */
func (this *MyQueue) Peek() int {
	return this.stack1[0]
}

/** Returns whether the queue is empty. */
func (this *MyQueue) Empty() bool {
	return len(this.stack1) == 0
}

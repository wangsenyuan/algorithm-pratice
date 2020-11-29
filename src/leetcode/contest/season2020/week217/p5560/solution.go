package p5560

type Node struct {
	val  int
	prev *Node
	next *Node
}

func NewNode(val int) *Node {
	node := new(Node)
	node.val = val
	node.prev = nil
	node.next = nil
	return node
}

type Queue struct {
	head *Node
	tail *Node
	size int
}

func NewQueue() *Queue {
	que := new(Queue)
	que.head = NewNode(-1)
	que.tail = NewNode(-2)
	que.head.next = que.tail
	que.tail.prev = que.head
	que.size = 0
	return que
}

func (q *Queue) PushFront(val int) {
	node := NewNode(val)
	tmp := q.head.next

	node.next = tmp
	tmp.prev = node

	node.prev = q.head
	q.head.next = node

	q.size++
}

func (q *Queue) PopFront() *Node {
	if q.size == 0 {
		return nil
	}
	ret := q.head.next
	tmp := ret.next
	q.head.next = tmp
	tmp.prev = q.head
	q.size--
	ret.next = nil
	ret.prev = nil
	return ret
}

func (q *Queue) PushTail(val int) {
	node := NewNode(val)
	tmp := q.tail.prev

	tmp.next = node
	node.prev = tmp

	node.next = q.tail
	q.tail.prev = node

	q.size++
}

func (q *Queue) PopTail() *Node {
	if q.size == 0 {
		return nil
	}
	ret := q.tail.prev
	tmp := ret.prev

	tmp.next = q.tail
	q.tail.prev = tmp
	q.size--
	ret.prev = nil
	ret.next = nil

	return ret
}

type FrontMiddleBackQueue struct {
	first  *Queue
	second *Queue
}

func Constructor() FrontMiddleBackQueue {
	first := NewQueue()
	second := NewQueue()
	return FrontMiddleBackQueue{first, second}
}

func (this *FrontMiddleBackQueue) PushFront(val int) {
	this.first.PushFront(val)
	this.adjust()
}

func (this *FrontMiddleBackQueue) adjust() {
	if this.second.size > 0 {
		this.first.PushTail(this.second.PopFront().val)
	}

	n := this.first.size + this.second.size
	for this.first.size > this.second.size+(n&1) {
		tmp := this.first.PopTail()
		this.second.PushFront(tmp.val)
	}
}

func (this *FrontMiddleBackQueue) PushMiddle(val int) {
	if this.first.size > this.second.size {
		this.second.PushFront(this.first.PopTail().val)
	}

	this.first.PushTail(val)
	this.adjust()
}

func (this *FrontMiddleBackQueue) PushBack(val int) {
	this.second.PushTail(val)

	this.adjust()
}

func (this *FrontMiddleBackQueue) PopFront() int {
	ret := this.first.PopFront()
	if ret == nil {
		return -1
	}
	ans := ret.val

	this.adjust()

	return ans
}

func (this *FrontMiddleBackQueue) PopMiddle() int {
	ret := this.first.PopTail()
	if ret == nil {
		return -1
	}
	ans := ret.val

	this.adjust()

	return ans
}

func (this *FrontMiddleBackQueue) PopBack() int {
	if this.second.size == 0 && this.first.size == 0 {
		return -1
	}
	var ans int
	if this.second.size > 0 {
		ans = this.second.PopTail().val
	} else {
		ans = this.first.PopTail().val
	}

	this.adjust()

	return ans
}

/**
 * Your FrontMiddleBackQueue object will be instantiated and called as such:
 * obj := Constructor();
 * obj.PushFront(val);
 * obj.PushMiddle(val);
 * obj.PushBack(val);
 * param_4 := obj.PopFront();
 * param_5 := obj.PopMiddle();
 * param_6 := obj.PopBack();
 */

package p855

type Node struct {
	index      int
	prev, next *Node
}

type ExamRoom struct {
	seats map[int]*Node
	head  *Node
	N     int
}

func Constructor(N int) ExamRoom {
	room := new(ExamRoom)
	room.seats = make(map[int]*Node)
	room.N = N
	return *room
}

func (this *ExamRoom) Seat() int {
	if this.head == nil {
		node := new(Node)
		node.index = 0
		this.head = node
		this.seats[0] = node
		return 0
	}

	node := new(Node)
	node.index = 0
	dist := this.head.index

	var prev, tail *Node

	for curr := this.head; curr != nil; curr = curr.next {
		if curr.prev != nil {
			d := (curr.index - curr.prev.index) / 2
			if d > dist {
				dist = d
				node.index = curr.prev.index + d
				prev = curr.prev
			}
		}
		tail = curr
	}

	if this.N-1-tail.index > dist {
		node.index = this.N - 1
		prev = tail
	}

	if prev == nil {
		// a new head
		node.next = this.head
		this.head.prev = node
		this.head = node
		this.seats[0] = node
		return 0
	}

	next := prev.next
	this.seats[node.index] = node
	prev.next, node.next, node.prev = node, next, prev
	if next != nil {
		next.prev = node
	}
	return node.index
}

func (this *ExamRoom) Leave(p int) {
	node := this.seats[p]
	delete(this.seats, p)

	prev, next := node.prev, node.next
	if prev == nil {
		//head deleted
		this.head = next
		if next != nil {
			next.prev = nil
		}
	} else {
		prev.next = next
		if next != nil {
			next.prev = prev
		}
	}
	node.prev = nil
	node.next = nil
}

/**
 * Your ExamRoom object will be instantiated and called as such:
 * obj := Constructor(N);
 * param_1 := obj.Seat();
 * obj.Leave(p);
 */

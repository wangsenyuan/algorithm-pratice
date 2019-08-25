package p1172

type DinnerPlates struct {
	capacity        int
	stacks          []*Stack
	unFullStackHead *Stack
	unFullStackTail *Stack
}

type Stack struct {
	pos   int
	nums  []int
	next  *Stack
	prev  *Stack
	index int
}

func Constructor(capacity int) DinnerPlates {
	var stacks []*Stack
	head := new(Stack)
	tail := new(Stack)
	head.next = tail
	tail.prev = head
	return DinnerPlates{capacity, stacks, head, tail}
}

func putNodeAt(before *Stack, next *Stack, node *Stack) {
	before.next = node
	node.prev = before
	node.next = next
	next.prev = node
}

func removeNode(node *Stack) {
	before := node.prev
	next := node.next
	before.next = next
	next.prev = before
	node.prev = nil
	node.next = nil
}

func (this *DinnerPlates) Push(val int) {
	if this.unFullStackHead.next == this.unFullStackTail {
		// empty
		node := new(Stack)
		node.nums = make([]int, this.capacity)
		node.nums[node.pos] = val
		node.pos++

		this.stacks = append(this.stacks, node)
		node.index = len(this.stacks) - 1

		if node.pos < this.capacity {
			putNodeAt(this.unFullStackHead, this.unFullStackTail, node)
		}
	} else {
		node := this.unFullStackHead.next
		node.nums[node.pos] = val
		node.pos++

		if node.pos == this.capacity {
			// need to remove it from the list
			removeNode(node)
		}
	}
}

func (this *DinnerPlates) Pop() int {
	// need to pop it from right
	n := len(this.stacks)

	if n == 0 {
		return -1
	}

	node := this.stacks[n-1]
	res := node.nums[node.pos-1]
	node.pos--

	if node.pos == 0 {
		// we need to shrink it
		this.stacks = this.stacks[:n-1]
		if this.unFullStackTail.prev == node {
			// normally it should be
			removeNode(node)
		}
	} else if node.pos+1 == this.capacity {
		// previsouely it is full, now put it at the end of tail
		putNodeAt(this.unFullStackTail.prev, this.unFullStackTail, node)
	}

	return res
}

func (this *DinnerPlates) PopAtStack(index int) int {
	n := len(this.stacks)
	if index >= n {
		return -1
	}
	node := this.stacks[index]
	if node.pos == 0 {
		return -1
	}
	res := node.nums[node.pos-1]
	node.pos--

	if node.pos < this.capacity {
		// need to put it in the un-full list
		i := index - 1
		for i >= 0 && this.stacks[i].prev == nil {
			i--
		}
		var before *Stack
		if i < 0 {
			before = this.unFullStackHead
		} else {
			before = this.stacks[i]
		}

		j := index + 1
		for j < n && this.stacks[j].next == nil {
			j++
		}

		var after *Stack
		if j == n {
			after = this.unFullStackTail
		} else {
			after = this.stacks[j]
		}

		putNodeAt(before, after, node)
	}

	return res
}

/**
 * Your DinnerPlates object will be instantiated and called as such:
 * obj := Constructor(capacity);
 * obj.Push(val);
 * param_2 := obj.Pop();
 * param_3 := obj.PopAtStack(index);
 */

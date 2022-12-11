package p2500

type Node struct {
	start int
	end   int
	next  *Node
}

func NewNode(start, end int) *Node {
	node := new(Node)
	node.start = start
	node.end = end
	return node
}

func (node *Node) Size() int {
	return node.end - node.start
}

type Allocator struct {
	free    *Node
	used    *Node
	records map[int][]int
}

func Constructor(n int) Allocator {
	free := NewNode(-1, -1)
	used := NewNode(-1, -1)
	free.next = NewNode(0, n)
	return Allocator{free, used, make(map[int][]int)}
}

func (this *Allocator) Allocate(size int, mID int) int {
	node := this.free.next
	prev := this.free
	for node != nil {
		if node.Size() >= size {
			a := NewNode(node.start, node.start+size)
			node.start += size
			if node.Size() == 0 {
				// need to free it
				prev.next = node.next
			}

			u := this.used

			for u.next != nil && u.next.end <= a.start {
				u = u.next
			}

			a.next = u.next
			u.next = a

			if this.records[mID] == nil {
				this.records[mID] = make([]int, 0, 1)
			}

			this.records[mID] = append(this.records[mID], a.start)

			return a.start
		}
		prev = node
		node = node.next
	}

	return -1
}

func (this *Allocator) Free(mID int) int {
	var res int
	for _, start := range this.records[mID] {
		p := this.used
		u := this.used.next
		for u.start != start {
			p = u
			u = u.next
		}
		p.next = u.next
		u.next = nil
		// u.start = start
		p = this.free
		n := this.free.next

		for n != nil && n.end < start {
			p = n
			n = n.next
		}

		if n == nil {
			p.next = u
		} else {
			// n.end >= start
			if n.end == start {
				if n.next != nil && u.end == n.next.start {
					n.end = n.next.end
					n.next = n.next.next
				} else {
					n.end = u.end
				}
			} else {
				// n.end > start
				if u.end == n.start {
					n.start = start
				} else {
					// u.end < n.start
					u.next = n
					p.next = u
				}
			}
		}

		res += u.Size()
	}

	delete(this.records, mID)

	return res
}

/**
 * Your Allocator object will be instantiated and called as such:
 * obj := Constructor(n);
 * param_1 := obj.Allocate(size,mID);
 * param_2 := obj.Free(mID);
 */

/**
 * Your Allocator object will be instantiated and called as such:
 * obj := Constructor(n);
 * param_1 := obj.Allocate(size,mID);
 * param_2 := obj.Free(mID);
 */

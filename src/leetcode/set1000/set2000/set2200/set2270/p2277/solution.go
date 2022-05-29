package p2277

type Node struct {
	start, end  int
	left, right *Node
	total_free  int64
	max_free    int64
}

func (node *Node) merge() {
	node.total_free = 0
	node.max_free = 0

	if node.left != nil {
		if node.left.max_free == 0 {
			node.left = nil
		} else {
			node.total_free += node.left.total_free
			node.max_free = node.left.max_free
		}
	}
	if node.right != nil {
		if node.right.max_free == 0 {
			node.right = nil
		} else {
			node.total_free += node.right.total_free
			if node.max_free < node.right.max_free {
				node.max_free = node.right.max_free
			}
		}
	}
}

func TakeLeftMostMaxFreeBefore(root *Node, k int64, maxRow int) []int {

	var loop func(node *Node) []int

	loop = func(node *Node) []int {

		if node.start == node.end {
			res := []int{node.start, int(node.max_free)}
			node.max_free -= k
			node.total_free -= k
			return res
		}
		var res []int
		if node.left != nil && node.left.max_free >= k && node.left.start <= maxRow {
			res = loop(node.left)
		} else if node.right != nil && node.right.max_free >= k && node.right.start <= maxRow {
			res = loop(node.right)
		}

		node.merge()

		return res
	}

	return loop(root)
}

func ScatterFreeBefore(root *Node, k int64, maxRow int) bool {
	var sum func(node *Node) int64

	sum = func(node *Node) int64 {
		if node == nil {
			return 0
		}
		if node.start > maxRow {
			return 0
		}
		if node.end <= maxRow {
			return node.total_free
		}
		// start < maxRow & end > maxRow
		return sum(node.left) + sum(node.right)
	}

	total := sum(root)

	if total < k {
		return false
	}

	var loop func(node *Node, num int64)

	loop = func(node *Node, num int64) {
		if node.start == node.end {
			node.total_free -= num
			node.max_free -= num
			return
		}
		if node.left != nil {
			if node.left.total_free <= num {
				num -= node.left.total_free
				node.left.max_free = 0
			} else {
				// num < node.left.total_free
				loop(node.left, num)
				num = 0
			}
		}

		if num > 0 {
			loop(node.right, num)
		}
		// node.total_free >= num
		node.merge()
	}

	loop(root, k)

	return true
}

func min(a, b int64) int64 {
	if a <= b {
		return a
	}
	return b
}

type BookMyShow struct {
	root *Node
	n    int
	m    int
}

func Constructor(n int, m int) BookMyShow {

	var build func(start, end int) *Node

	build = func(start, end int) *Node {
		node := new(Node)
		node.start = start
		node.end = end

		if start == end {
			node.total_free = int64(m)
			node.max_free = int64(m)
			return node
		}
		mid := (start + end) / 2
		node.left = build(start, mid)
		node.right = build(mid+1, end)
		node.merge()
		return node
	}
	root := build(0, n-1)
	return BookMyShow{root, n, m}
}

func (this *BookMyShow) Gather(k int, maxRow int) []int {
	if this.root.max_free < int64(k) {
		return nil
	}
	res := TakeLeftMostMaxFreeBefore(this.root, int64(k), maxRow)
	if len(res) == 0 {
		return nil
	}
	res[1] = this.m - res[1]
	return res
}

func (this *BookMyShow) Scatter(k int, maxRow int) bool {
	return ScatterFreeBefore(this.root, int64(k), maxRow)
}

/**
 * Your BookMyShow object will be instantiated and called as such:
 * obj := Constructor(n, m);
 * param_1 := obj.Gather(k,maxRow);
 * param_2 := obj.Scatter(k,maxRow);
 */

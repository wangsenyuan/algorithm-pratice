package main

func main() {

}

func solve(N, Q, L, R int, queries [][]int) []int {
	return nil
}

type SegTree struct {
	start, end  int   // the range covered
	total       int64 // total number of sub array in [start:end] staisfy the condition
	A, B        int   // A left-most value out of [L:R] or start, B right-most value out of [L:R] or end
	hasHole     bool
	left, right *SegTree
}

func createNode(start, end int) *SegTree {
	node := new(SegTree)
	node.start = start
	node.end = end
	return node
}

func (root *SegTree) Query(l, r int, L, R int) int64 {
	var loop func(node *SegTree) int64
	loop = func(node *SegTree) int64 {
		if node == nil {
			return 0
		}
		if node.start < l || node.end > r {
			return 0
		}
		if l <= node.start && node.end <= r {
			return node.total
		}

		res := loop(node.left) + loop(node.right)

		if node.left == nil || node.right == nil {
			return res
		}
		var A, B int
		if node.left.hasHole {
			A = max(node.left.B, l-1)
		} else {
			// start need to be included
			A = max(node.start-1, l-1)
		}

		if node.right.hasHole {
			B = min(node.right.A, r+1)
		} else {
			B = min(node.end+1, r+1)
		}

		tmp := int64(node.left.end-A) * int64(B-node.right.start)
		return res + tmp
	}
	return loop(root)
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func (root *SegTree) Update(pos int, val, L, R int) {

	var loop func(node *SegTree, start, end int) *SegTree
	loop = func(node *SegTree, start, end int) *SegTree {
		if node == nil {
			node = createNode(start, end)
		}

		if start == end {
			if val >= L && val <= R {
				node.total = 1
			} else {
				node.hasHole = true
				node.A = start
				node.B = end
			}
			return node
		}

		mid := (start + end) / 2
		if pos <= mid {
			node.left = loop(node.left, start, mid)
		} else {
			node.right = loop(node.right, mid+1, end)
		}

		if node.right == nil {
			node.total = node.left.total
			node.hasHole = node.left.hasHole
			node.A = node.left.A
			node.B = node.left.B
		} else if node.left == nil {
			node.total = node.right.total
			node.hasHole = node.right.hasHole
			node.A = node.right.A
			node.B = node.right.B
		} else {
			node.hasHole = node.left.hasHole || node.right.hasHole
			node.total = node.left.total + node.right.total
			if node.left.hasHole && !node.right.hasHole {
				node.A = node.left.A
				node.B = node.left.B
				node.total += int64(mid-node.left.B) * int64(end-mid)
			} else if !node.left.hasHole && node.right.hasHole {
				node.A = node.right.A
				node.B = node.right.B
				node.total += int64(mid-start+1) * int64(node.right.B-mid-1)
			} else {
				node.A = node.left.A
				node.B = node.right.B
				node.total += int64(mid-node.left.B) * int64(node.right.B-mid-1)
			}
		}
		return node
	}

	loop(root, root.start, root.end)
}

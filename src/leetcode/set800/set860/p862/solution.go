package p862

import "sort"

func shortestSubarray(A []int, K int) int {
	n := len(A)
	var sum int64
	sums := make([]int64, n+1)
	sums[0] = 0
	stack := make([]int, n+1)
	var p int
	stack[p] = 0
	p++

	best := -1
	for i := 1; i <= n; i++ {
		sum += int64(A[i-1])
		sums[i] = sum
		prev := sum - int64(K)
		j := sort.Search(p, func(j int) bool {
			return sums[stack[j]] > prev
		})
		j--
		if j >= 0 {
			// found a value less than or equal to prev
			// stack[j] is the position
			if best < 0 || i-stack[j] < best {
				best = i - stack[j]
			}
		}
		for p > 0 && sums[stack[p-1]] >= sum {
			p--
		}
		stack[p] = i
		p++
	}
	return best
}

func shortestSubarray1(A []int, K int) int {
	n := len(A)
	var sum int64

	var root *Treap
	root = root.Insert(0, -1)

	best := -1

	for i := 0; i < n; i++ {
		sum += int64(A[i])
		j, found := root.Find(sum - int64(K))
		if found {
			if best < 0 || i-j < best {
				best = i - j
			}
		}
		root = root.Insert(sum, i)
	}

	return best
}

type Treap struct {
	val         int64
	priority    int
	left, right *Treap
}

func (root *Treap) Insert(val int64, priority int) *Treap {
	// priority is the index
	// val is the sum
	var loop func(node *Treap) *Treap
	loop = func(node *Treap) *Treap {
		if node == nil {
			node = new(Treap)
			node.val = val
			node.priority = priority
			return node
		}
		if val <= node.val {
			node.left = loop(node.left)
		} else {
			node.right = loop(node.right)
		}
		return balance(node)
	}
	return loop(root)
}

func (root *Treap) Find(val int64) (int, bool) {
	var loop func(node *Treap) *Treap

	loop = func(node *Treap) *Treap {
		if node == nil {
			return nil
		}
		if node.val <= val {
			return node
		}
		return loop(node.left)
	}
	res := loop(root)

	if res == nil {
		return -1, false
	}

	// find the right-most index that is less than val
	return res.priority, true
}

func balance(node *Treap) *Treap {
	if node.left != nil && node.left.priority > node.priority {
		node = rotateRight(node)
	}

	if node.right != nil && node.right.priority > node.priority {
		node = rotateLeft(node)
	}

	return node
}

func rotateRight(node *Treap) *Treap {
	tmp := node.left
	node.left = tmp.right
	tmp.right = node
	return tmp
}

func rotateLeft(node *Treap) *Treap {
	tmp := node.right
	node.right = tmp.left
	tmp.left = node
	return tmp
}

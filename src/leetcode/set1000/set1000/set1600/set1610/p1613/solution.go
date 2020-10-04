package p1613

func busiestServers(k int, arrival []int, load []int) []int {
	tree := BuildTree(k)

	cnt := make([]int, k)

	for i := 0; i < len(arrival); i++ {
		x := arrival[i]
		y := load[i]
		j := i % k
		tmp := tree.Find(j, k, x)
		if tmp < 0 {
			// not found
			tmp = tree.Find(0, j, x)
		}
		if tmp < 0 {
			continue
		}
		tree.Update(tmp, x+y)
		cnt[tmp]++
	}
	var best int

	for i := 1; i < k; i++ {
		if cnt[i] > cnt[best] {
			best = i
		}
	}

	res := make([]int, 0, k)

	for i := 0; i < k; i++ {
		if cnt[i] == cnt[best] {
			res = append(res, i)
		}
	}

	return res
}

type Node struct {
	left, right *Node
	start, end  int
	minValue    int
}

func (node Node) GetMid() int {
	return (node.start + node.end) / 2
}

type Tree struct {
	nodes []*Node
	n     int
	root  *Node
}

func BuildTree(n int) Tree {
	nodes := make([]*Node, n)
	var loop func(start, end int) *Node
	loop = func(start, end int) *Node {
		if start+1 == end {
			node := new(Node)
			node.start = start
			node.end = end
			nodes[start] = node
			return node
		}
		mid := (start + end) / 2
		left := loop(start, mid)
		right := loop(mid, end)
		node := new(Node)
		node.start = start
		node.end = end
		node.left = left
		node.right = right
		return node
	}
	root := loop(0, n)

	return Tree{nodes, n, root}
}

func (tree Tree) Find(l, r int, x int) int {
	// this loop is to find the root node, that cover [l...r)
	node := tree.root.Find(l, r, x)
	if node == nil {
		return -1
	}
	return node.start
}

func (node *Node) Find(l, r int, x int) *Node {
	if node == nil || node.start >= r || node.end <= l {
		return nil
	}
	if node.minValue > x {
		return nil
	}

	if node.left == nil && node.right == nil {
		return node
	}

	res := node.left.Find(l, min(r, node.left.end), x)
	if res == nil {
		res = node.right.Find(max(l, node.right.start), r, x)
	}
	return res
}

func (tree *Tree) Update(p int, x int) {
	tree.root.Update(p, x)
}

func (node *Node) Update(p int, x int) {
	if node.start == node.end-1 {
		node.minValue = x
		return
	}
	if p < node.GetMid() {
		node.left.Update(p, x)
	} else {
		node.right.Update(p, x)
	}
	node.minValue = min(node.left.minValue, node.right.minValue)
}

func min(a, b int) int {
	if a <= b {
		return a
	}
	return b
}

func max(a, b int) int {
	if a >= b {
		return a
	}
	return b
}

package p2002

func smallestMissingValueSubtree(parents []int, nums []int) []int {
	n := len(parents)
	// n <= 100000

	g := NewGraph(n, n-1)

	for i := 1; i < n; i++ {
		g.AddEdge(parents[i], i)
	}

	var root *Node

	pick := make([]int, n)

	// dfs is used to find which child has 1
	var dfs func(u int) bool

	dfs = func(u int) bool {
		pick[u] = -1

		if nums[u] == 1 {
			return true
		}
		var ok bool
		for i := g.nodes[u]; i > 0; i = g.next[i] {
			if dfs(g.to[i]) {
				pick[u] = g.to[i]
				ok = true
				break
			}
		}
		return ok
	}

	ans := make([]int, n)

	for i := 0; i < n; i++ {
		ans[i] = 1
	}

	if !dfs(0) {
		// no 1 find
		return ans
	}

	var dfs2 func(u int)

	dfs2 = func(u int) {
		if nums[u] <= n {
			root = Set(root, nums[u]-1, 0, n-1)
		}
		for i := g.nodes[u]; i > 0; i = g.next[i] {
			dfs2(g.to[i])
		}
	}

	var dfs1 func(u int)

	dfs1 = func(u int) {
		if pick[u] < 0 {
			// u has nums[u] == 1
			dfs2(u)
		} else {
			// first process pick[u]
			dfs1(pick[u])

			for i := g.nodes[u]; i > 0; i = g.next[i] {
				v := g.to[i]
				if v != pick[u] {
					dfs2(v)
				}
			}

			if nums[u] <= n {
				root = Set(root, nums[u]-1, 0, n-1)
			}
		}

		if root != nil && root.size == n {
			ans[u] = n + 1
		} else {
			ans[u] = Find(root, 0, n-1) + 1
		}
	}

	dfs1(0)

	return ans
}

type Node struct {
	left, right *Node
	start, end  int
	size        int
}

func Set(node *Node, v int, start, end int) *Node {
	if node == nil {
		node = new(Node)
		node.start = start
		node.end = end
	}
	node.size++
	if start == end {
		return node
	}
	mid := (start + end) / 2
	if v <= mid {
		node.left = Set(node.left, v, start, mid)
	} else {
		node.right = Set(node.right, v, mid+1, end)
	}
	return node
}

func Find(node *Node, start, end int) int {
	if node == nil {
		return start
	}
	if start == end {
		return start
	}

	mid := (start + end) / 2
	if node.left == nil || node.left.size < mid-start+1 {
		return Find(node.left, start, mid)
	}

	return Find(node.right, mid+1, end)
}

type Graph struct {
	nodes []int
	next  []int
	to    []int
	cur   int
}

func NewGraph(n int, e int) *Graph {
	nodes := make([]int, n)
	next := make([]int, e+2)
	to := make([]int, e+2)
	return &Graph{nodes, next, to, 0}
}

func (g *Graph) AddEdge(u, v int) {
	g.cur++
	g.next[g.cur] = g.nodes[u]
	g.nodes[u] = g.cur
	g.to[g.cur] = v
}

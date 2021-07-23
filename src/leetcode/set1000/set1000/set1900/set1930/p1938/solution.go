package p1938

func maxGeneticDifference(parents []int, queries [][]int) []int {
	n := len(parents)
	g := NewGraph(n, n-1)

	var root int

	for i := 0; i < n; i++ {
		p := parents[i]
		if p >= 0 {
			g.AddEdge(p, i)
		} else {
			root = i
		}
	}

	nodes := make([]*Node, n)
	nodes[root] = AddValue(nil, root, H-1)
	var dfs func(u int)

	dfs = func(u int) {
		for i := g.nodes[u]; i > 0; i = g.next[i] {
			v := g.to[i]
			nodes[v] = AddValue(nodes[u], v, H-1)
			dfs(v)
		}
	}

	dfs(root)

	res := make([]int, len(queries))

	for i, cur := range queries {
		x, v := cur[0], cur[1]
		res[i] = FindMaxXor(nodes[x], v)
	}
	return res
}

const H = 18

type Node struct {
	children [2]*Node
}

func copyNode(src *Node) *Node {
	node := new(Node)
	if src != nil {
		node.children[0] = src.children[0]
		node.children[1] = src.children[1]
	}
	return node
}

func AddValue(node *Node, v int, p int) *Node {
	if p < 0 {
		return new(Node)
	}
	res := copyNode(node)
	x := (v >> p) & 1
	res.children[x] = AddValue(res.children[x], v, p-1)
	return res
}

func FindMaxXor(node *Node, v int) int {
	tmp := node
	var res int
	for i := H - 1; i >= 0; i-- {
		x := (v >> i) & 1
		if tmp.children[1-x] != nil {
			res |= 1 << i
			tmp = tmp.children[1-x]
		} else {
			tmp = tmp.children[x]
		}
	}
	return res
}

type Graph struct {
	nodes []int
	next  []int
	to    []int
	cur   int
}

func NewGraph(n int, e int) *Graph {
	e += 10
	g := new(Graph)
	g.nodes = make([]int, n)
	g.next = make([]int, e)
	g.to = make([]int, e)
	g.cur = 0
	return g
}

func (g *Graph) AddEdge(u, v int) {
	g.cur++
	g.next[g.cur] = g.nodes[u]
	g.nodes[u] = g.cur
	g.to[g.cur] = v
}

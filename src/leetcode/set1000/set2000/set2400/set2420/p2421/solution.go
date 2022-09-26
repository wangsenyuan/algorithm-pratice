package p2421

import "sort"

func numberOfGoodPaths(vals []int, edges [][]int) int {
	// 观察，如果节点u的值比它的所有子孙节点的值要大，那么就没有必要再保留它们
	n := len(vals)

	g := NewGraph(n, 2*(n-1))

	for _, edge := range edges {
		u, v := edge[0], edge[1]
		g.AddEdge(u, v)
		g.AddEdge(v, u)
	}

	nodes := make([]Node, n)

	for i := 0; i < n; i++ {
		nodes[i] = Node{i, vals[i]}
	}

	sort.Slice(nodes, func(i, j int) bool {
		return nodes[i].val < nodes[j].val || nodes[i].val == nodes[j].val && nodes[i].id < nodes[j].id
	})

	arr := make([]int, n)
	cnt := make([]int, n)
	sz := make([]int, n)

	for i := 0; i < n; i++ {
		arr[i] = i
		cnt[i] = 1
		sz[i] = 1
	}

	var find func(x int) int

	find = func(x int) int {
		if arr[x] != x {
			arr[x] = find(arr[x])
		}
		return arr[x]
	}

	union := func(x int, y int) {
		px := find(x)
		py := find(y)
		if px == py {
			return
		}
		if cnt[px] < cnt[py] {
			px, py = py, px
		}
		cnt[px] += cnt[py]
		arr[py] = px
		sz[px] += sz[py]
	}

	var res int

	for j, k := 0, 0; j < n; j++ {
		node := nodes[j]
		u := node.id

		for k < j && nodes[k].val < node.val {
			pv := find(nodes[k].id)
			sz[pv]--
			k++
		}

		for i := g.node[u]; i > 0; i = g.next[i] {
			v := g.to[i]
			if node.val > vals[v] || node.val == vals[v] && u > v {
				// v is processed already
				union(u, v)
			}
		}

		pu := find(u)

		res += sz[pu]
	}

	return res
}

func numberOfGoodPaths1(vals []int, edges [][]int) int {
	// 观察，如果节点u的值比它的所有子孙节点的值要大，那么就没有必要再保留它们
	n := len(vals)

	g := NewGraph(n, 2*(n-1))

	for _, edge := range edges {
		u, v := edge[0], edge[1]
		g.AddEdge(u, v)
		g.AddEdge(v, u)
	}

	nodes := make([]Node, n)

	for i := 0; i < n; i++ {
		nodes[i] = Node{i, vals[i]}
	}

	sort.Slice(nodes, func(i, j int) bool {
		return nodes[i].val < nodes[j].val || nodes[i].val == nodes[j].val && nodes[i].id < nodes[j].id
	})

	arr := make([]int, n)
	cnt := make([]int, n)
	sz := make([]map[int]int, n)

	for i := 0; i < n; i++ {
		arr[i] = i
		cnt[i] = 1
		sz[i] = make(map[int]int)
		sz[i][vals[i]]++
	}

	var find func(x int) int

	find = func(x int) int {
		if arr[x] != x {
			arr[x] = find(arr[x])
		}
		return arr[x]
	}

	union := func(x int, y int) {
		px := find(x)
		py := find(y)
		if px == py {
			return
		}
		if cnt[px] < cnt[py] {
			px, py = py, px
		}
		cnt[px] += cnt[py]
		arr[py] = px
		for k, v := range sz[py] {
			sz[px][k] += v
		}
	}

	var res int

	for _, node := range nodes {
		u := node.id
		for i := g.node[u]; i > 0; i = g.next[i] {
			v := g.to[i]
			if node.val > vals[v] || node.val == vals[v] && u > v {
				// v is processed already
				union(u, v)
			}
		}
		pu := find(u)
		res += sz[pu][node.val]
	}

	return res
}

type Node struct {
	id  int
	val int
}

type Graph struct {
	node []int
	next []int
	to   []int
	cur  int
}

func NewGraph(n int, e int) *Graph {
	g := new(Graph)
	g.node = make([]int, n)
	g.next = make([]int, e+1)
	g.to = make([]int, e+1)
	return g
}

func (g *Graph) AddEdge(u, v int) {
	g.cur++
	g.next[g.cur] = g.node[u]
	g.node[u] = g.cur
	g.to[g.cur] = v
}

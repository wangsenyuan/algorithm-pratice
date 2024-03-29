package main

import (
	"bufio"
	"bytes"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	n, m := readTwoNums(reader)
	edges := make([][]int, n-1)
	for i := 0; i < n-1; i++ {
		edges[i] = readNNums(reader, 2)
	}
	queries := make([][]int, m)
	for i := 0; i < m; i++ {
		queries[i] = readNNums(reader, 3)
	}

	res := solve(n, edges, queries)
	var buf bytes.Buffer
	for _, x := range res {
		buf.WriteString(fmt.Sprintf("%d\n", x))
	}

	fmt.Print(buf.String())
}

func readString(reader *bufio.Reader) string {
	s, _ := reader.ReadString('\n')
	for i := 0; i < len(s); i++ {
		if s[i] == '\n' || s[i] == '\r' {
			return s[:i]
		}
	}
	return s
}

func normalize(s string) string {

	for i := len(s); i > 0; i-- {
		if s[i-1] >= 'a' && s[i-1] <= 'z' {
			return s[:i]
		}
	}
	return ""
}

func readInt(bytes []byte, from int, val *int) int {
	i := from
	sign := 1
	if bytes[i] == '-' {
		sign = -1
		i++
	}
	tmp := 0
	for i < len(bytes) && bytes[i] >= '0' && bytes[i] <= '9' {
		tmp = tmp*10 + int(bytes[i]-'0')
		i++
	}
	*val = tmp * sign
	return i
}

func readNum(reader *bufio.Reader) (a int) {
	bs, _ := reader.ReadBytes('\n')
	readInt(bs, 0, &a)
	return
}

func readTwoNums(reader *bufio.Reader) (a int, b int) {
	res := readNNums(reader, 2)
	a, b = res[0], res[1]
	return
}

func readThreeNums(reader *bufio.Reader) (a int, b int, c int) {
	res := readNNums(reader, 3)
	a, b, c = res[0], res[1], res[2]
	return
}

func readNNums(reader *bufio.Reader, n int) []int {
	res := make([]int, n)
	x := 0
	bs, _ := reader.ReadBytes('\n')
	for i := 0; i < n; i++ {
		for x < len(bs) && (bs[x] < '0' || bs[x] > '9') && bs[x] != '-' {
			x++
		}
		x = readInt(bs, x, &res[i])
	}
	return res
}

const H = 20

const inf = 1 << 60

func solve(n int, edges [][]int, queries [][]int) []int {
	g := NewGraph(n, n)

	for i := 1; i < n; i++ {
		edge := edges[i-1]
		j, w := edge[0]-1, edge[1]

		g.AddEdge(j, i, w)
	}

	arr := make([]int, n)

	sz := make([]int, n)
	var dfs func(u int, dist int)
	dfs = func(u int, dist int) {
		arr[u] = inf
		sz[u] = 1
		for i := g.nodes[u]; i > 0; i = g.next[i] {
			v := g.to[i]
			dfs(v, dist+g.val[i])
			sz[u] += sz[v]
		}
		if g.nodes[u] == 0 {
			arr[u] = dist
		}
	}

	dfs(0, 0)

	tr := Build(arr)

	qs := make([][]int, n)
	for i, cur := range queries {
		v := cur[0] - 1
		if len(qs[v]) == 0 {
			qs[v] = make([]int, 0, 1)
		}
		qs[v] = append(qs[v], i)
	}

	ans := make([]int, len(queries))

	var dfs2 func(u int)

	dfs2 = func(u int) {
		for _, i := range qs[u] {
			cur := queries[i]
			l, r := cur[1]-1, cur[2]-1
			ans[i] = tr.Get(l, r, n)
		}

		for i := g.nodes[u]; i > 0; i = g.next[i] {
			v := g.to[i]
			w := g.val[i]
			// 所有的加w
			tr.Update(0, n-1, w, n)
			// 但是v的子树中的-w
			tr.Update(v, v+sz[v]-1, -2*w, n)

			dfs2(v)
			// restore
			tr.Update(v, v+sz[v]-1, 2*w, n)
			tr.Update(0, n-1, -w, n)
		}
	}

	dfs2(0)

	return ans
}

func min(a, b int) int {
	if a <= b {
		return a
	}
	return b
}

type Node struct {
	left, right *Node
	lazy        int
	val         int
}

func Build(arr []int) *Node {
	n := len(arr)
	var loop func(l int, r int) *Node
	loop = func(l int, r int) *Node {
		node := new(Node)
		if l == r {
			node.val = arr[l]
			return node
		}
		node.val = inf
		mid := (l + r) / 2
		node.left = loop(l, mid)
		node.right = loop(mid+1, r)
		node.val = min(node.left.val, node.right.val)
		return node
	}
	return loop(0, n-1)
}

func (node *Node) update(v int) {
	node.val += v
	node.lazy += v
}

func (node *Node) push() {
	if node.lazy != 0 && node.left != nil {
		node.left.update(node.lazy)
		node.right.update(node.lazy)
		node.lazy = 0
	}
}

func (node *Node) Update(L int, R int, v int, n int) {
	var loop func(node *Node, l int, r int)

	loop = func(node *Node, l int, r int) {
		if r < L || R < l {
			return
		}
		node.push()
		if L <= l && r <= R {
			node.update(v)
			return
		}
		mid := (l + r) / 2
		loop(node.left, l, mid)
		loop(node.right, mid+1, r)
		node.val = min(node.left.val, node.right.val)
	}

	loop(node, 0, n-1)
}

func (node *Node) Get(L int, R int, n int) int {
	var loop func(node *Node, l int, r int) int
	loop = func(node *Node, l int, r int) int {
		if r < L || R < l {
			return inf
		}
		node.push()
		if L <= l && r <= R {
			return node.val
		}
		mid := (l + r) / 2
		a := loop(node.left, l, mid)
		b := loop(node.right, mid+1, r)
		return min(a, b)
	}
	return loop(node, 0, n-1)
}

type Graph struct {
	nodes []int
	next  []int
	to    []int
	val   []int
	cur   int
}

func NewGraph(n int, e int) *Graph {
	nodes := make([]int, n)
	e++
	next := make([]int, e)
	to := make([]int, e)
	val := make([]int, e)
	return &Graph{nodes, next, to, val, 0}
}

func (g *Graph) AddEdge(u, v, w int) {
	g.cur++
	g.next[g.cur] = g.nodes[u]
	g.nodes[u] = g.cur
	g.to[g.cur] = v
	g.val[g.cur] = w
}

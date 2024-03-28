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
	c := readNNums(reader, n)
	edges := make([][]int, n-1)
	for i := 0; i < n-1; i++ {
		edges[i] = readNNums(reader, 2)
	}
	queries := make([][]int, m)
	for i := 0; i < m; i++ {
		s, _ := reader.ReadBytes('\n')
		var t int
		pos := readInt(s, 0, &t) + 1
		if t == 1 {
			queries[i] = make([]int, 3)
		} else {
			queries[i] = make([]int, 2)
		}
		queries[i][0] = t
		for j := 1; j < len(queries[i]); j++ {
			pos = readInt(s, pos, &queries[i][j]) + 1
		}
	}

	res := solve(n, c, edges, queries)

	var buf bytes.Buffer
	for _, x := range res {
		buf.WriteString(fmt.Sprintf("%d\n", x))
	}
	fmt.Print(buf.String())
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

func solve(n int, c []int, edges [][]int, queries [][]int) []int {
	g := NewGraph(n, n*2)

	for _, e := range edges {
		u, v := e[0]-1, e[1]-1
		g.AddEdge(u, v)
		g.AddEdge(v, u)
	}

	in := make([]int, n)
	out := make([]int, n)
	var arr []int

	var dfs func(p int, u int)

	dfs = func(p int, u int) {
		in[u] = len(arr)
		arr = append(arr, c[u])

		for i := g.nodes[u]; i > 0; i = g.next[i] {
			v := g.to[i]
			if p != v {
				dfs(u, v)
			}
		}

		out[u] = len(arr)
	}

	dfs(-1, 0)

	tr := Build(arr)

	var ans []int

	for _, cur := range queries {
		if cur[0] == 1 {
			v, x := cur[1]-1, cur[2]
			tr.Set(0, n-1, in[v], out[v]-1, x)
		} else {
			v := cur[1] - 1
			tmp := tr.Get(0, n-1, in[v], out[v]-1)
			ans = append(ans, digitCount(tmp))
		}
	}

	return ans
}

func digitCount(num int) int {
	var res int
	for num > 0 {
		res++
		num -= num & -num
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
	nodes := make([]int, n)
	next := make([]int, e)
	to := make([]int, e)
	return &Graph{nodes, next, to, 0}
}

func (g *Graph) AddEdge(u, v int) {
	g.cur++
	g.next[g.cur] = g.nodes[u]
	g.nodes[u] = g.cur
	g.to[g.cur] = v
}

type Node struct {
	left, right *Node
	lazy        int
	val         int
}

func Build(arr []int) *Node {
	var loop func(l int, r int) *Node
	loop = func(l int, r int) *Node {
		node := new(Node)

		if l == r {
			node.lazy = 0
			node.val = 1 << (arr[l] - 1)
		} else {
			mid := (l + r) / 2
			node.left = loop(l, mid)
			node.right = loop(mid+1, r)
			node.val = node.left.val | node.right.val
		}
		return node
	}

	return loop(0, len(arr)-1)
}

func (node *Node) set(v int) {
	node.lazy = v
	node.val = 1 << (v - 1)
}

func (node *Node) push() {
	if node.lazy != 0 && node.left != nil {
		node.left.set(node.lazy)
		node.right.set(node.lazy)
		node.lazy = 0
	}
}

func (node *Node) Set(l int, r int, L int, R int, v int) {
	if r < L || R < l {
		// no overlap
		return
	}
	node.push()
	if L <= l && r <= R {
		node.set(v)
		return
	}
	mid := (l + r) / 2
	node.left.Set(l, mid, L, R, v)
	node.right.Set(mid+1, r, L, R, v)
	node.val = node.left.val | node.right.val
}

func (node *Node) Get(l int, r int, L int, R int) int {
	if r < L || R < l {
		return 0
	}
	node.push()
	if L <= l && r <= R {
		return node.val
	}
	mid := (l + r) / 2
	a := node.left.Get(l, mid, L, R)
	b := node.right.Get(mid+1, r, L, R)
	return a | b
}

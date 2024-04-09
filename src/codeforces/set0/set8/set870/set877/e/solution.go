package main

import (
	"bufio"
	"bytes"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	n := readNum(reader)
	P := readNNums(reader, n-1)
	a := readNNums(reader, n)
	m := readNum(reader)
	queries := make([]string, m)
	for i := 0; i < m; i++ {
		queries[i] = readString(reader)
	}

	res := solve(n, P, a, queries)

	var buf bytes.Buffer
	for i := 0; i < len(res); i++ {
		buf.WriteString(fmt.Sprintf("%d\n", res[i]))
	}

	fmt.Print(buf.String())
}

func readString(r *bufio.Reader) string {
	s, _ := r.ReadString('\n')
	for i := 0; i < len(s); i++ {
		if s[i] == '\n' || s[i] == '\r' {
			return s[:i]
		}
	}
	return s
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

// about 1000000
const H = 20

func solve(n int, p []int, lights []int, queries []string) []int {
	g := NewGraph(n, n)

	for i := 0; i < n-1; i++ {
		g.AddEdge(p[i]-1, i+1)
	}
	ord := make([]int, n)
	var arr []int
	sz := make([]int, n)
	var dfs func(u int)
	dfs = func(u int) {
		ord[u] = len(arr)
		arr = append(arr, lights[u])
		sz[u] = 1
		for i := g.nodes[u]; i > 0; i = g.next[i] {
			v := g.to[i]
			dfs(v)
			sz[u] += sz[v]
		}
	}
	dfs(0)

	root := Build(arr)

	var ans []int

	for _, cur := range queries {
		var v int
		readInt([]byte(cur), 4, &v)
		v--
		if cur[:3] == "pow" {
			root.Update(ord[v], ord[v]+sz[v]-1, n)
		} else {
			tmp := root.Query(ord[v], ord[v]+sz[v]-1, n)
			ans = append(ans, tmp)
		}
	}

	return ans
}

type Node struct {
	left, right *Node
	lazy        int
	cnt         [2]int
}

func (node *Node) pull() {
	for i := 0; i < 2; i++ {
		node.cnt[i] = node.left.cnt[i] + node.right.cnt[i]
	}
}

func Build(arr []int) *Node {
	var loop func(l int, r int) *Node

	loop = func(l int, r int) *Node {
		node := new(Node)
		if l == r {
			node.cnt[arr[l]]++
			return node
		}
		mid := (l + r) / 2
		node.left = loop(l, mid)
		node.right = loop(mid+1, r)
		node.pull()
		return node
	}

	return loop(0, len(arr)-1)
}

func (node *Node) update(v int) {
	node.lazy += v
	// node.lazy = 目前为止的翻转次数
	if v%2 == 1 {
		node.cnt[0], node.cnt[1] = node.cnt[1], node.cnt[0]
	}
}

func (node *Node) push() {
	if node.lazy != 0 && node.left != nil {
		node.left.update(node.lazy)
		node.right.update(node.lazy)
		node.lazy = 0
	}
}

func (root *Node) Update(L int, R int, n int) {
	var loop func(node *Node, l int, r int)

	loop = func(node *Node, l int, r int) {
		if R < l || r < L {
			return
		}
		node.push()
		if L <= l && r <= R {
			node.update(1)
			return
		}
		mid := (l + r) / 2
		loop(node.left, l, mid)
		loop(node.right, mid+1, r)
		node.pull()
	}

	loop(root, 0, n-1)
}

func (root *Node) Query(L int, R int, n int) int {
	var loop func(node *Node, l int, r int) int

	loop = func(node *Node, l int, r int) int {
		if R < l || r < L {
			return 0
		}
		node.push()
		if L <= l && r <= R {
			return node.cnt[1]
		}
		mid := (l + r) / 2
		a := loop(node.left, l, mid)
		b := loop(node.right, mid+1, r)
		return a + b
	}

	return loop(root, 0, n-1)
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

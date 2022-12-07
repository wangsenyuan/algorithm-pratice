package main

import (
	"bufio"
	"bytes"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	var buf bytes.Buffer

	tc := readNum(reader)

	for tc > 0 {
		tc--
		n := readNum(reader)
		A := readNNums(reader, n)
		edges := make([][]int, n-1)
		for i := 0; i < n-1; i++ {
			edges[i] = readNNums(reader, 2)
		}
		q := readNum(reader)
		queries := make([][]int, q)
		for i := 0; i < q; i++ {
			queries[i] = readAtMostNNums(reader, 3)
		}

		res := solve(n, A, edges, queries)

		for _, ans := range res {
			buf.WriteString(fmt.Sprintf("%d\n", ans))
		}
	}

	fmt.Print(buf.String())
}

func readAtMostNNums(reader *bufio.Reader, n int) []int {
	res := make([]int, 0, n)
	x := 0
	bs, _ := reader.ReadBytes('\n')
	for i := 0; i < n; i++ {
		for x < len(bs) && (bs[x] < '0' || bs[x] > '9') {
			x++
		}
		if x == len(bs) {
			break
		}
		var y int

		x = readInt(bs, x, &y)

		res = append(res, y)
	}

	return res
}

func readUint64(bytes []byte, from int, val *uint64) int {
	i := from

	var tmp uint64
	for i < len(bytes) && bytes[i] >= '0' && bytes[i] <= '9' {
		tmp = tmp*10 + uint64(bytes[i]-'0')
		i++
	}
	*val = tmp

	return i
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

func readString(reader *bufio.Reader) string {
	s, _ := reader.ReadString('\n')
	for i := 0; i < len(s); i++ {
		if s[i] == '\n' || s[i] == '\r' {
			return s[:i]
		}
	}
	return s
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

func solve(n int, A []int, edges [][]int, queries [][]int) []int {
	//	假设 1 u, c, 直接按照原来的树结构处理，会出现o(n)的情况
	// 如果转化成欧拉tour，可以在o(lgn)时间内处理， （用 segment lazy update)
	// 但是如何处理转换root呢？
	// 考虑现在的root 是 x != 1
	//	执行 1 u, c (u != x)
	// 如果 u （在原图中）是 x子树部分, 可以按照原图处理
	// 如果 u （在原图中）不是x的祖先节点， 也可以按原图处理
	// 如果 u 是 x的祖先，则需要将全部更新成c，且将u的子树恢复成原有的颜色
	// 如果 u == x, 更新全图
	// query也是一样的，就是query(1, c) - query(children, c)
	g := NewGraph(n, 2*n)

	for _, edge := range edges {
		u, v := edge[0], edge[1]
		u--
		v--
		g.AddEdge(u, v)
		g.AddEdge(v, u)
	}

	in := make([]int, n)
	out := make([]int, n)

	anc := make([][]int, n)

	var time int

	var dfs func(p int, u int)
	dfs = func(p int, u int) {
		anc[u] = make([]int, H)
		anc[u][0] = p
		for i := 1; i < H; i++ {
			anc[u][i] = anc[anc[u][i-1]][i-1]
		}
		in[u] = time
		time++

		for i := g.nodes[u]; i > 0; i = g.next[i] {
			v := g.to[i]
			if p != v {
				dfs(u, v)
			}
		}

		out[u] = time
	}

	dfs(0, 0)

	isAnc := func(u int, v int) bool {
		return in[u] <= in[v] && out[v] <= out[u]
	}

	getChild := func(u int, v int) int {
		for i := H - 1; i >= 0; i-- {
			if !isAnc(anc[u][i], v) {
				u = anc[u][i]
			}
		}
		return u
	}

	tree := NewNode(0, n)

	for i := 0; i < n; i++ {
		tree.Set(in[i], A[i])
	}

	var res []int

	var root int

	for _, qr := range queries {
		if qr[0] == 1 {
			u := qr[1] - 1
			c := qr[2]
			if u == root {
				tree.RangeSet(0, n, c, true)
			} else if isAnc(u, root) {
				v := getChild(root, u)
				tree.RangeSet(out[v], n, c, true)
				tree.RangeSet(0, in[v], c, true)
			} else {
				tree.RangeSet(in[u], out[u], c, true)
			}
		} else if qr[0] == 2 {
			root = qr[1] - 1
		} else {
			// query
			u := qr[1] - 1
			c := qr[2]
			var ans int
			if u == root {
				ans = tree.Query(0, n, c)
			} else if isAnc(u, root) {
				v := getChild(root, u)
				ans = tree.Query(out[v], n, c) + tree.Query(0, in[v], c)
			} else {
				ans = tree.Query(in[u], out[u], c)
			}
			res = append(res, ans)
		}
	}

	return res
}

type Node struct {
	freq   map[int]int
	l, r   *Node
	par    *Node
	lo, hi int
	lazy   bool
	val    int
}

func NewNode(lo int, hi int) *Node {
	node := new(Node)
	node.freq = make(map[int]int)
	node.lo = lo
	node.hi = hi
	node.lazy = false
	node.val = 0
	return node
}

func (node *Node) push() {
	if node.l == nil {
		mid := (node.lo + node.hi) / 2
		node.l = NewNode(node.lo, mid)
		node.r = NewNode(mid, node.hi)
		node.l.par = node
		node.r.par = node
	}
	if node.lazy {
		node.l.RangeSet(node.lo, node.hi, node.val, false)
		node.r.RangeSet(node.lo, node.hi, node.val, false)
		node.lazy = false
	}
}

func (node *Node) Set(p int, x int) {
	node.freq[x]++
	if node.lo+1 == node.hi {
		return
	}
	node.push()
	mid := (node.lo + node.hi) / 2
	if p < mid {
		node.l.Set(p, x)
	} else {
		node.r.Set(p, x)
	}
}

func (node *Node) RangeSet(L int, R int, x int, up bool) {
	if node.hi <= L || R <= node.lo {
		return
	}
	if L <= node.lo && node.hi <= R {
		node.lazy = true
		node.val = x
		if up {
			cur := node.par

			for cur != nil {
				for k, v := range node.freq {
					cur.freq[k] -= v
					if cur.freq[k] == 0 {
						delete(cur.freq, k)
					}
				}
				cur.freq[x] += node.hi - node.lo
				cur = cur.par
			}
		}
		node.freq = make(map[int]int)
		node.freq[x] = node.hi - node.lo
		return
	}
	node.push()
	node.l.RangeSet(L, R, x, up)
	node.r.RangeSet(L, R, x, up)
}

func (node *Node) Query(L int, R int, c int) int {
	if R <= node.lo || node.hi <= L {
		return 0
	}
	if L <= node.lo && node.hi <= R {
		return node.freq[c]
	}
	node.push()
	return node.l.Query(L, R, c) + node.r.Query(L, R, c)
}

type Graph struct {
	nodes []int
	next  []int
	to    []int
	cur   int
}

func NewGraph(n int, e int) *Graph {
	g := new(Graph)
	g.nodes = make([]int, n)
	g.next = make([]int, e+3)
	g.to = make([]int, e+3)
	g.cur = 0
	return g
}

func (g *Graph) AddEdge(u, v int) {
	g.cur++
	g.next[g.cur] = g.nodes[u]
	g.nodes[u] = g.cur
	g.to[g.cur] = v
}

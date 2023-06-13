package main

import (
	"bufio"
	"bytes"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	tc := readNum(reader)

	var buf bytes.Buffer
	for tc > 0 {
		tc--
		n := readNum(reader)
		P1 := readNNums(reader, n)
		P2 := readNNums(reader, n)
		res := solve(n, P1, P2)
		buf.WriteString(fmt.Sprintf("%d\n", res))
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

func readNInt64s(reader *bufio.Reader, n int) []int64 {
	res := make([]int64, n)
	s, _ := reader.ReadBytes('\n')

	var pos int

	for i := 0; i < n; i++ {
		pos = readInt64(s, pos, &res[i]) + 1
	}

	return res
}

func readInt64(bytes []byte, from int, val *int64) int {
	i := from
	var sign int64 = 1
	if bytes[i] == '-' {
		sign = -1
		i++
	}
	var tmp int64
	for i < len(bytes) && bytes[i] >= '0' && bytes[i] <= '9' {
		tmp = tmp*10 + int64(bytes[i]-'0')
		i++
	}
	*val = tmp * sign
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

	for len(bs) == 0 || bs[0] == '\n' || bs[0] == '\r' {
		bs, _ = reader.ReadBytes('\n')
	}

	for i := 0; i < n; i++ {
		for x < len(bs) && (bs[x] < '0' || bs[x] > '9') && bs[x] != '-' {
			x++
		}
		x = readInt(bs, x, &res[i])
	}
	return res
}

const H = 20

func solve(n int, P1 []int, P2 []int) int {
	// 有几个观察，如果w是满足的公共祖先
	//   1 存在 (u1, v2) w1是u1的祖先，w2是v2的祖先
	//   2. 但是 w也是 (w1, w2)的祖先
	//   3. 所以，可以找到深度的节点x, 它在两棵树中的深度是一致的
	//   4. 这样的节点，组成的数的，最大的深度，就是答案
	g1 := CreateGraph(n, P1)
	g2 := CreateGraph(n, P2)

	//arr := make([]int, n)

	in := make([]int, n)
	out := make([]int, n)
	D := make([]int, n)
	var time int
	var dfs func(u int)
	dfs = func(u int) {
		in[u] = time
		//arr[time] = u
		time++
		for i := g1.nodes[u]; i > 0; i = g1.next[i] {
			v := g1.to[i]
			D[v] = D[u] + 1
			dfs(v)
		}
		out[u] = time
	}

	dfs(g1.root)

	tr := NewSegTree(n)

	var res int
	var dfs2 func(u int, depth int)
	dfs2 = func(u int, depth int) {
		if depth == D[u] {
			tr.Update(in[u], out[u]-1, 1)
			res = max(res, tr.Get(in[u], out[u]-1))
		}

		for i := g2.nodes[u]; i > 0; i = g2.next[i] {
			v := g2.to[i]
			dfs2(v, depth+1)
		}

		if depth == D[u] {
			tr.Update(in[u], out[u]-1, -1)
		}
	}

	dfs2(g2.root, 0)

	return res
}

type SegTree struct {
	nodes []int
	lazy  []int
	n     int
}

func NewSegTree(n int) *SegTree {
	arr := make([]int, 4*n)
	lazy := make([]int, 4*n)
	return &SegTree{arr, lazy, n}
}

func (t *SegTree) pushLazy(i int, v int) {
	t.nodes[i] += v
	t.lazy[i] += v
}

func (t *SegTree) push(i int, l int, r int) {
	if t.lazy[i] != 0 {
		if l < r {
			t.pushLazy(2*i, t.lazy[i])
			t.pushLazy(2*i+1, t.lazy[i])
		}
		t.lazy[i] = 0
	}
}

func (t *SegTree) Update(L int, R int, v int) {
	var dfs func(i int, l int, r int)
	dfs = func(i int, l int, r int) {
		if R < l || r < L {
			return
		}
		t.push(i, l, r)
		if L <= l && r <= R {
			t.pushLazy(i, v)
			return
		}
		mid := (l + r) / 2
		dfs(2*i, l, mid)
		dfs(2*i+1, mid+1, r)
		t.nodes[i] = max(t.nodes[i*2], t.nodes[i*2+1])
	}

	dfs(1, 0, t.n-1)
}

func (t *SegTree) Get(L int, R int) int {
	var dfs func(i int, l int, r int) int

	dfs = func(i int, l int, r int) int {
		if R < l || r < L {
			return 0
		}
		t.push(i, l, r)
		if L <= l && r <= R {
			return t.nodes[i]
		}
		mid := (l + r) / 2
		a := dfs(2*i, l, mid)
		b := dfs(2*i+1, mid+1, r)
		return max(a, b)
	}

	return dfs(1, 0, t.n-1)
}

func max(a, b int) int {
	if a >= b {
		return a
	}
	return b
}

func CreateGraph(n int, P []int) *Graph {
	g := NewGraph(n, n)
	for i := 0; i < n; i++ {
		j := P[i]
		j--
		if j >= 0 {
			g.AddEdge(j, i)
		} else {
			g.root = i
		}
	}
	return g
}

type Graph struct {
	nodes []int
	next  []int
	to    []int
	cur   int
	root  int
}

func NewGraph(n int, e int) *Graph {
	nodes := make([]int, n)
	next := make([]int, e)
	to := make([]int, e)
	return &Graph{nodes, next, to, 0, -1}
}

func (g *Graph) AddEdge(u, v int) {
	g.cur++
	g.next[g.cur] = g.nodes[u]
	g.nodes[u] = g.cur
	g.to[g.cur] = v
}

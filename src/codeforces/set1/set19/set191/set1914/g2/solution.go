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
		bulbs := readNNums(reader, 2*n)
		sz, ways := solve(n, bulbs)
		buf.WriteString(fmt.Sprintf("%d %d\n", sz, ways))
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

const mod = 998244353

func add(a, b int) int {
	a += b
	if a >= mod {
		a -= mod
	}
	return a
}

func mul(a, b int) int {
	return a * b % mod
}

func pow(a, b int) int {
	r := 1
	for b > 0 {
		if b&1 == 1 {
			r = mul(r, a)
		}
		a = mul(a, a)
		b >>= 1
	}
	return r
}

func solve(n int, bulbs []int) (int, int) {
	// 每个color you两个位置
	tr := NewSegTree(n * 2)

	for i := 0; i < n*2; i++ {
		tr.Update(i, inf)
	}
	// pos[i] 表示第二个位置
	pos := make([]int, n)

	for i, c := range bulbs {
		pos[c-1] = i
	}

	g := NewGraph(n, 2*n)
	r := NewGraph(n, 2*n)

	for i, c := range bulbs {
		j := tr.Get(i+1, n*2)
		if j < inf {
			g.AddEdge(c-1, bulbs[j]-1)
			r.AddEdge(bulbs[j]-1, c-1)
		}
		j = pos[c-1]
		if j > i {
			tr.Update(j, j)
		} else {
			tr.Update(j, inf)
		}
	}

	vis := make([]bool, n)

	var order []int

	var dfs func(u int)

	dfs = func(u int) {
		vis[u] = true
		for i := g.nodes[u]; i > 0; i = g.next[i] {
			v := g.to[i]
			if !vis[v] {
				dfs(v)
			}
		}
		order = append(order, u)
	}

	for i := 0; i < n; i++ {
		if !vis[i] {
			dfs(i)
		}
	}

	belong := make([]int, n)
	for i := 0; i < n; i++ {
		belong[i] = -1
	}

	comp := make([][]int, n)

	var dfs2 func(u int, c int)

	dfs2 = func(u int, c int) {
		comp[c] = append(comp[c], u)
		belong[u] = c
		for i := r.nodes[u]; i > 0; i = r.next[i] {
			v := r.to[i]
			if belong[v] < 0 {
				dfs2(v, c)
			}
		}
	}
	var cid int
	for i := n - 1; i >= 0; i-- {
		u := order[i]
		if belong[u] < 0 {
			dfs2(u, cid)
			cid++
		}
	}

	deg := make([]int, cid)
	for u := 0; u < n; u++ {
		uc := belong[u]
		for i := g.nodes[u]; i > 0; i = g.next[i] {
			v := g.to[i]
			vc := belong[v]
			if uc != vc {
				deg[uc]++
			}
		}
	}

	var set int
	ways := 1

	for i := 0; i < cid; i++ {
		if deg[i] == 0 {
			set++
			ways = mul(ways, mul(2, len(comp[i])))
		}
	}

	return set, ways
}

type Graph struct {
	nodes []int
	next  []int
	to    []int
	cur   int
}

func NewGraph(n int, e int) *Graph {
	nodes := make([]int, n)
	next := make([]int, e+10)
	to := make([]int, e+10)
	return &Graph{nodes, next, to, 0}
}

func (g *Graph) AddEdge(u, v int) {
	g.cur++
	g.next[g.cur] = g.nodes[u]
	g.nodes[u] = g.cur
	g.to[g.cur] = v
}

const inf = 1 << 30

type SegTree struct {
	arr []int
	sz  int
}

func NewSegTree(n int) *SegTree {
	arr := make([]int, 2*n)
	return &SegTree{arr, n}
}

func (t *SegTree) Update(p int, v int) {
	p += t.sz
	t.arr[p] = v
	for p > 1 {
		t.arr[p>>1] = min(t.arr[p], t.arr[p^1])
		p >>= 1
	}
}

func (t *SegTree) Get(l int, r int) int {
	l += t.sz
	r += t.sz
	var res int = inf
	for l < r {
		if l&1 == 1 {
			res = min(res, t.arr[l])
			l++
		}
		if r&1 == 1 {
			r--
			res = min(res, t.arr[r])
		}
		l >>= 1
		r >>= 1
	}
	return res
}

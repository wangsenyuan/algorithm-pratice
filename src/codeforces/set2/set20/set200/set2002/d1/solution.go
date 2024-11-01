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
		n, m := readTwoNums(reader)
		a := readNNums(reader, n-1)
		p := readNNums(reader, n)
		queries := make([][]int, m)
		for i := range m {
			queries[i] = readNNums(reader, 2)
		}
		res := solve(n, a, p, queries)
		for _, x := range res {
			if x {
				buf.WriteString("YES\n")
			} else {
				buf.WriteString("NO\n")
			}
		}
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
	for i := 0; i < n; i++ {
		for x < len(bs) && (bs[x] < '0' || bs[x] > '9') && bs[x] != '-' {
			x++
		}
		x = readInt(bs, x, &res[i])
	}
	return res
}

func solve(n int, a []int, p []int, queries [][]int) []bool {
	g := NewGraph(n, n)

	fa := make([]int, n)
	for i := 0; i < n-1; i++ {
		fa[i+1] = a[i] - 1
		g.AddEdge(fa[i+1], i+1)
	}
	sz := make([]int, n)
	pos := make([]int, n)
	var clock int
	var dfs func(u int)
	dfs = func(u int) {
		pos[u] = clock
		clock++
		sz[u]++
		for i := g.nodes[u]; i > 0; i = g.next[i] {
			v := g.to[i]
			dfs(v)
			sz[u] += sz[v]
		}
	}

	fa[0] = -1
	dfs(0)
	tr1 := NewSegTree(n, -1, func(a, b int) int {
		return max(a, b)
	})

	tr2 := NewSegTree(n, n, func(a, b int) int {
		return min(a, b)
	})

	at := make([]int, n)

	for i := 0; i < n; i++ {
		p[i]--
		at[p[i]] = i
		tr1.Update(i, pos[p[i]])
		tr2.Update(i, pos[p[i]])
	}

	checkBad := func(u int) bool {
		// i是u现在所在的位置
		i := at[u]
		if i+sz[u] > n {
			return true
		}
		// i + sz[u] <= n
		v := tr1.Get(i, i+sz[u], -1)
		w := tr2.Get(i, i+sz[u], n)
		// j是里面最大的位置
		// 这个区间内都有值，但不一定连续
		return v-w+1 != sz[u] || v >= pos[u]+sz[u] || w < pos[u]
	}

	bad := make(map[int]bool)

	for i := 0; i < n; i++ {
		if checkBad(i) {
			bad[i] = true
		}
	}

	update := func(x int) {
		for x >= 0 {
			tmp := checkBad(x)
			if tmp {
				bad[x] = true
			} else {
				delete(bad, x)
			}
			x = fa[x]
		}
	}

	swap := func(x int, y int) {
		// x, y 是p的下标
		u, v := p[x], p[y]
		// px 的值要变成py
		tr1.Update(at[u], pos[v])
		tr1.Update(at[v], pos[u])

		tr2.Update(at[u], pos[v])
		tr2.Update(at[v], pos[u])

		p[x], p[y] = v, u
		at[u], at[v] = y, x

		update(u)
		update(v)
	}

	ans := make([]bool, len(queries))

	for i, cur := range queries {
		x, y := cur[0], cur[1]
		swap(x-1, y-1)
		ans[i] = len(bad) == 0
	}

	return ans
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

type SegTree struct {
	arr []int
	sz  int
	op  func(int, int) int
}

func NewSegTree(n int, iv int, op func(int, int) int) *SegTree {
	res := make([]int, 2*n)

	for i := 2*n - 1; i > 0; i-- {
		res[i] = iv
	}

	return &SegTree{res, n, op}
}

func (t *SegTree) Update(p int, v int) {
	p += t.sz
	t.arr[p] = v
	for p > 1 {
		t.arr[p>>1] = t.op(t.arr[p], t.arr[p^1])
		p >>= 1
	}
}

func (t *SegTree) Get(l int, r int, res int) int {
	l += t.sz
	r += t.sz

	for l < r {
		if l&1 == 1 {
			res = t.op(res, t.arr[l])
			l++
		}
		if r&1 == 1 {
			r--
			res = t.op(res, t.arr[r])
		}
		l >>= 1
		r >>= 1
	}
	return res
}

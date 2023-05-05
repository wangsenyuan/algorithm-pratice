package main

import (
	"bufio"
	"bytes"
	"fmt"
	"os"
	"sort"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	n := readNum(reader)
	E := make([][]int, n-1)
	for i := 0; i < n-1; i++ {
		E[i] = readNNums(reader, 2)
	}
	solve(n, E, reader)
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

func readString(reader *bufio.Reader) string {
	s, _ := reader.ReadString('\n')
	for i := 0; i < len(s); i++ {
		if s[i] == '\n' {
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

func solve(n int, E [][]int, reader *bufio.Reader) {
	g := NewGraph(n, 2*n)

	for i, e := range E {
		u, v := e[0], e[1]
		u--
		v--
		g.AddEdge(u, v, i+1)
		g.AddEdge(v, u, i+1)
	}

	edge := make([]int, n)
	edge[0] = -1
	sz := make([]int, n)
	D := make([]int, n)
	P := make([]int, n)
	big := make([]int, n)
	var dfs1 func(p int, u int)

	dfs1 = func(p int, u int) {
		P[u] = p
		sz[u]++
		big[u] = -1
		for i := g.nodes[u]; i > 0; i = g.next[i] {
			v := g.to[i]
			if p != v {
				edge[v] = g.val[i]
				D[v] = D[u] + 1
				dfs1(u, v)
				sz[u] += sz[v]
				if big[u] < 0 || sz[big[u]] < sz[v] {
					big[u] = v
				}
			}
		}
	}

	dfs1(-1, 0)

	in := make([]int, n)
	next := make([]int, n)

	var timer int

	var dfs2 func(p int, u int, c int)

	dfs2 = func(p int, u int, c int) {
		in[u] = timer
		timer++
		next[u] = c
		if big[u] >= 0 {
			dfs2(u, big[u], c)
		}
		for i := g.nodes[u]; i > 0; i = g.next[i] {
			v := g.to[i]
			if v != p && v != big[u] {
				dfs2(u, v, v)
			}
		}
	}
	dfs2(-1, 0, 0)

	t := NewTree(n)

	flip := func(l int, r int) {
		t.Flip(0, 0, n, l, r+1)
	}

	get := func(l int, r int) Val {
		return t.GetVal(0, 0, n, l, r+1)
	}

	flipPath := func(v int, u int) {
		for ; next[v] != next[u]; u = P[next[u]] {
			if D[next[v]] > D[next[u]] {
				u, v = v, u
			}
			flip(in[next[u]], in[u])
		}
		if D[v] > D[u] {
			v, u = u, v
		}
		// D[v] <= D[u]
		flip(in[v], in[u])
	}

	getPath := func(v int, u int) Val {
		res := Val{0, 0}
		for ; next[v] != next[u]; u = P[next[u]] {
			if D[next[v]] > D[next[u]] {
				u, v = v, u
			}
			res = res.Add(get(in[next[u]], in[u]))
		}
		if D[v] > D[u] {
			u, v = v, u
		}
		res = res.Add(get(in[v], in[u]))
		return res
	}

	var sum int64
	var cnt int

	active := make([]bool, n)
	var activeCnt int

	activate := func(x int) {
		var id int
		if P[x] >= 0 {
			id = edge[x]
			v := getPath(0, P[x])
			sum -= v.x
			cnt -= v.y
			flipPath(0, P[x])
			v = getPath(0, P[x])
			sum += v.x
			cnt += v.y
		}
		cnt++
		sum += int64(id)
		t.SetVal(0, 0, n, in[x], Val{int64(id), 1})
		active[x] = true
		activeCnt++
	}

	query1 := func(x int) {
		activate(x)
		var ans int64
		if cnt*2 == activeCnt {
			ans = sum
		}
		fmt.Printf("%d\n", ans)
	}

	var dfs3 func(x int, edges *[]int) int

	currentSize := make([]int, n)

	dfs3 = func(x int, edges *[]int) int {
		if !active[x] {
			return 0
		}
		currentSize[x] = 1
		for i := g.nodes[x]; i > 0; i = g.next[i] {
			y := g.to[i]
			if y != P[x] {
				currentSize[x] += dfs3(y, edges)
			}
		}
		if currentSize[x]%2 == 1 {
			*edges = append(*edges, edge[x])
		}
		return currentSize[x]
	}

	query2 := func() {
		var edges []int
		if cnt*2 == activeCnt {
			dfs3(0, &edges)
			sort.Ints(edges)
		}
		var buf bytes.Buffer
		buf.WriteString(fmt.Sprintf("%d", len(edges)))
		for _, x := range edges {
			buf.WriteString(fmt.Sprintf(" %d", x))
		}
		fmt.Println(buf.String())
	}

	activate(0)

	for {
		s, _ := reader.ReadBytes('\n')
		if s[0] == '1' {
			var x int
			readInt(s, 2, &x)
			x--
			query1(x)
		} else if s[0] == '2' {
			query2()
		} else {
			break
		}
	}
}

type Val struct {
	x int64
	y int
}

func (this Val) Add(that Val) Val {
	return Val{this.x + that.x, this.y + that.y}
}

func (this Val) Sub(that Val) Val {
	return Val{this.x - that.x, this.y - that.y}
}

type Tree struct {
	arr  []Val
	full []Val
	f    []int
	n    int
}

func NewTree(n int) *Tree {
	arr := make([]Val, 4*n)
	full := make([]Val, 4*n)
	f := make([]int, 4*n)
	return &Tree{arr, full, f, n}
}

func (t *Tree) getVal(v int) Val {
	if t.f[v] == 1 {
		return t.full[v].Sub(t.arr[v])
	}
	return t.arr[v]
}

func (t *Tree) push(v int) {
	if t.f[v] == 1 {
		t.arr[v] = t.full[v].Sub(t.arr[v])
		t.f[2*v+1] ^= 1
		t.f[2*v+2] ^= 1
		t.f[v] = 0
	}
}

func (t *Tree) update(v int) {
	t.full[v] = t.full[2*v+1].Add(t.full[2*v+2])
	t.arr[v] = t.getVal(2*v + 1).Add(t.getVal(2*v + 2))
}

func (t *Tree) SetVal(v int, l int, r int, pos int, cur Val) {
	if r-l == 1 {
		t.f[v] = 0
		t.full[v] = cur
		t.arr[v] = cur
		return
	}
	t.push(v)
	mid := (l + r) / 2
	if pos < mid {
		t.SetVal(2*v+1, l, mid, pos, cur)
	} else {
		t.SetVal(2*v+2, mid, r, pos, cur)
	}
	t.update(v)
}

func (t *Tree) Flip(v int, l int, r int, L int, R int) {
	if L >= R {
		return
	}
	if L == l && r == R {
		t.f[v] ^= 1
		return
	}
	t.push(v)
	mid := (l + r) / 2
	t.Flip(2*v+1, l, mid, L, min(R, mid))
	t.Flip(2*v+2, mid, r, max(L, mid), R)
	t.update(v)
}

func (t *Tree) GetVal(v int, l int, r int, L int, R int) Val {
	if L >= R {
		return Val{0, 0}
	}
	if L == l && r == R {
		return t.getVal(v)
	}
	mid := (l + r) / 2
	ans := Val{0, 0}
	t.push(v)
	ans = ans.Add(t.GetVal(2*v+1, l, mid, L, min(R, mid)))
	ans = ans.Add(t.GetVal(2*v+2, mid, r, max(L, mid), R))
	t.update(v)
	return ans
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
	next := make([]int, e+3)
	to := make([]int, e+3)
	val := make([]int, e+3)
	var cur int
	return &Graph{nodes, next, to, val, cur}
}

func (g *Graph) AddEdge(u, v, w int) {
	g.cur++
	g.next[g.cur] = g.nodes[u]
	g.nodes[u] = g.cur
	g.to[g.cur] = v
	g.val[g.cur] = w
}

func min(a, b int) int {
	if a <= b {
		return a
	}
	return b
}

func max(a, b int) int {
	return a + b - min(a, b)
}

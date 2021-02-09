package main

import (
	"bufio"
	"bytes"
	"fmt"
	"math"
	"os"
	"sort"
)

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
		for x < len(bs) && (bs[x] < '0' || bs[x] > '9') {
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

func main() {
	reader := bufio.NewReader(os.Stdin)

	n, q := readTwoNums(reader)
	C := readNNums(reader, n)

	E := make([][]int, n-1)
	for i := 0; i < n-1; i++ {
		E[i] = readNNums(reader, 2)
	}
	Q := make([][]int, q)
	for i := 0; i < q; i++ {
		Q[i] = readNNums(reader, 3)
	}
	res := solve(n, q, C, E, Q)

	var buf bytes.Buffer
	for i := 0; i < len(res); i++ {
		buf.WriteString(fmt.Sprintf("%d\n", res[i]))
	}
	fmt.Print(buf.String())
}

const H = 20

func solve(n int, q int, C []int, E [][]int, Q [][]int) []int {
	var counter = 1
	compress := make(map[int]int)
	for i := 0; i < n; i++ {
		if _, found := compress[C[i]]; !found {
			compress[C[i]] = counter
			counter++
		}
		C[i] = compress[C[i]]
	}

	g := NewGraph(n, len(E)*2+2)

	for _, e := range E {
		u, v := e[0], e[1]
		u--
		v--
		g.AddEdge(u, v)
		g.AddEdge(v, u)
	}

	dep := make([]int, n)
	fa := make([][]int, H)
	for i := 0; i < H; i++ {
		fa[i] = make([]int, n)
	}
	dfn := make([]int, n)
	dfsSeq := make([]int, 2*n+10)
	var dfsTot int

	var dfs func(p, u int)

	dfs = func(p, u int) {
		fa[0][u] = p
		dfn[u] = dfsTot
		dfsSeq[dfsTot] = u
		dfsTot++
		for i := g.nodes[u]; i > 0; i = g.next[i] {
			v := g.to[i]
			if v != p {
				dep[v] = dep[u] + 1
				dfs(u, v)
			}
		}
		dfsSeq[dfsTot] = u
		dfsTot++
	}

	dfs(-1, 0)

	blk := int(math.Sqrt(float64(dfsTot)))

	belong := make([]int, dfsTot+10)

	for i := 0; i < n; i++ {
		belong[i] = i / blk
	}

	for i := 1; i < H; i++ {
		for j := 0; j < n; j++ {
			if fa[i-1][j] < 0 {
				fa[i][j] = -1
			} else {
				fa[i][j] = fa[i-1][fa[i-1][j]]
			}
		}
	}

	lca := func(x, y int) int {
		if dep[x] > dep[y] {
			x, y = y, x
		}
		for i := 0; dep[x] != dep[y]; i++ {
			if ((dep[y]-dep[x])>>uint(i))&1 == 1 {
				y = fa[i][y]
			}
		}

		s := H - 1

		for x != y {
			for s > 0 && fa[s][x] == fa[s][y] {
				s--
			}
			x = fa[s][x]
			y = fa[s][y]
		}

		return x
	}

	id := make([]int, q)
	for i := 0; i < q; i++ {
		id[i] = -1
	}

	to := make([]int, q)
	queries := make([]Query, 0, q)

	var totQ, totM int
	for i := 0; i < q; i++ {
		ty, x, y := Q[i][0], Q[i][1], Q[i][2]
		x--
		y--
		if ty == 2 {
			totM++
			id[totM] = x
			y++
			if _, ok := compress[y]; !ok {
				compress[y] = counter
				counter++
			}
			to[totM] = compress[y]
		} else {
			if dfn[x] > dfn[y] {
				x, y = y, x
			}
			queries = append(queries, Query{dfn[x] + 1, dfn[y] + 1, totM, totQ})
			totQ++
		}
	}

	qq := Queries{queries, belong}

	sort.Sort(qq)

	cnt := make([]int, len(compress)+10)

	in := make([]int, dfsTot)

	var now int

	inc := func(x int) {
		if cnt[x] == 0 {
			now++
		}
		cnt[x]++
	}

	dec := func(x int) {
		cnt[x]--
		if cnt[x] == 0 {
			now--
		}
	}

	change := func(x int) {
		x = dfsSeq[x]
		if in[x] > 0 {
			in[x] = 0
			dec(C[x])
		} else {
			in[x]++
			inc(C[x])
		}
	}

	refresh := func(t int) {
		if id[t] != -1 && C[id[t]] != to[t] {
			if in[id[t]] > 0 {
				dec(C[id[t]])
			}
			C[id[t]], to[t] = to[t], C[id[t]]
			if in[id[t]] > 0 {
				inc(C[id[t]])
			}
		}
	}

	ans := make([]int, totQ)
	var l, r, t int
	for i := 0; i < totQ; i++ {
		q := qq.items[i]
		for l > q.x {
			l--
			change(l)
		}
		for r < q.y {
			change(r)
			r++
		}
		for l < q.x {
			change(l)
			l++
		}
		for r > q.y {
			r--
			change(r)
		}
		for t > q.z {
			refresh(t)
			t--
		}
		for t < q.z {
			t++
			refresh(t)
		}
		mid := lca(dfsSeq[q.x-1], dfsSeq[q.y-1])

		if in[mid] == 0 {
			inc(C[mid])
		}
		ans[q.id] = now
		if in[mid] == 0 {
			dec(C[mid])
		}
	}
	return ans
}

type Query struct {
	x, y, z, id int
}

type Queries struct {
	items  []Query
	belong []int
}

func (this Queries) Len() int {
	return len(this.items)
}

func (this Queries) Less(i, j int) bool {
	items := this.items
	belong := this.belong
	a, b := items[i], items[j]
	if belong[a.x] == belong[b.x] {
		return a.x < b.x
	}
	if belong[a.y] != belong[b.y] {
		return a.y < b.y
	}

	return a.z < b.z
}

func (this Queries) Swap(i, j int) {
	this.items[i], this.items[j] = this.items[j], this.items[i]
}

type Graph struct {
	nodes []int
	next  []int
	to    []int
	cur   int
}

func NewGraph(n int, sz int) *Graph {
	nodes := make([]int, n)
	next := make([]int, sz)
	to := make([]int, sz)
	return &Graph{nodes, next, to, 1}
}

func (g *Graph) AddEdge(u, v int) {
	g.cur++
	g.next[g.cur] = g.nodes[u]
	g.nodes[u] = g.cur
	g.to[g.cur] = v
}

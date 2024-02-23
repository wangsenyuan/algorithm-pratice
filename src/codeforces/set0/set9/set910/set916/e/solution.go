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
	a := readNNums(reader, n)
	edges := make([][]int, n-1)
	for i := 0; i < n-1; i++ {
		edges[i] = readNNums(reader, 2)
	}
	ops := make([][]int, m)
	for i := 0; i < m; i++ {
		s, _ := reader.ReadBytes('\n')
		if s[0] == '1' {
			// change root
			var r int
			readInt(s, 2, &r)
			ops[i] = []int{1, r}
		} else if s[0] == '2' {
			var u, v, c int
			pos := readInt(s, 2, &u)
			pos = readInt(s, pos+1, &v)
			readInt(s, pos+1, &c)
			ops[i] = []int{2, u, v, c}
		} else {
			var u int
			readInt(s, 2, &u)
			ops[i] = []int{3, u}
		}
	}
	res := solve(a, edges, ops)
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

const H = 18

func solve(a []int, edges [][]int, ops [][]int) []int {
	n := len(a)
	g := NewGraph(n, 2*n)

	for _, edge := range edges {
		u, v := edge[0]-1, edge[1]-1
		g.AddEdge(u, v)
		g.AddEdge(v, u)
	}

	dep := make([]int, n)
	ord := make([]int, n)
	eul := make([]int, n)
	sz := make([]int, n)

	ancs := make([][]int, n)

	var time int

	var dfs func(p int, u int)
	dfs = func(p int, u int) {
		ord[u] = time
		eul[time] = u
		time++
		sz[u] = 1

		ancs[u] = make([]int, H)
		ancs[u][0] = p
		for i := 1; i < H; i++ {
			ancs[u][i] = ancs[ancs[u][i-1]][i-1]
		}
		for i := g.nodes[u]; i > 0; i = g.next[i] {
			v := g.to[i]
			if p != v {
				dep[v] = dep[u] + 1
				dfs(u, v)
				sz[u] += sz[v]
			}
		}
	}

	dfs(0, 0)

	kthAnc := func(u int, k int) int {
		for i := H - 1; i >= 0; i-- {
			if (k>>i)&1 == 1 {
				u = ancs[u][i]
			}
		}
		return u
	}

	lca := func(u int, v int) int {
		if dep[u] < dep[v] {
			u, v = v, u
		}
		u = kthAnc(u, dep[u]-dep[v])
		if u == v {
			return u
		}
		for i := H - 1; i >= 0; i-- {
			if ancs[u][i] != ancs[v][i] {
				u = ancs[u][i]
				v = ancs[v][i]
			}
		}
		return ancs[u][0]
	}

	arr := make([]int, n)
	for i := 0; i < n; i++ {
		arr[i] = a[eul[i]]
	}

	st := NewSegTree(arr)

	root := 0
	rl, rr := 0, n-1

	var ans []int

	for _, cur := range ops {
		if cur[0] == 1 {
			root = cur[1] - 1
			rl = ord[root]
			rr = ord[root] + sz[root] - 1
		} else if cur[0] == 2 {
			u, v, c := cur[1]-1, cur[2]-1, cur[3]
			var cnt int
			if rl <= ord[u] && ord[u] <= rr {
				cnt++
			}
			if rl <= ord[v] && ord[v] <= rr {
				cnt++
			}
			var origin int
			if cnt == 2 {
				origin = lca(u, v)
			} else if cnt == 1 {
				origin = root
			} else {
				// cnt == 0
				x := lca(u, root)
				y := lca(v, root)
				z := lca(u, v)
				origin = x
				if dep[y] > dep[origin] {
					origin = y
				}
				if dep[z] > dep[origin] {
					origin = z
				}
			}
			if origin == root {
				st.Update(0, n-1, c)
			} else if rl < ord[origin] && ord[origin] <= rr {
				// origin 是 root的子树
				st.Update(ord[origin], ord[origin]+sz[origin]-1, c)
			} else if ord[origin] < ord[root] && ord[root] <= ord[origin]+sz[origin]-1 {
				// root是origin的子树
				st.Update(0, n-1, c)
				undo := kthAnc(root, dep[root]-dep[origin]-1)
				st.Update(ord[undo], ord[undo]+sz[undo]-1, -c)
			} else {
				// root和origin不在一条path上
				st.Update(ord[origin], ord[origin]+sz[origin]-1, c)
			}
		} else {
			origin := cur[1] - 1
			var tmp int
			if origin == root {
				tmp = st.Query(0, n-1)
			} else if rl < ord[origin] && ord[origin] <= rr {
				tmp = st.Query(ord[origin], ord[origin]+sz[origin]-1)
			} else if ord[origin] < ord[root] && ord[root] <= ord[origin]+sz[origin]-1 {
				tmp = st.Query(0, n-1)
				undo := kthAnc(root, dep[root]-dep[origin]-1)
				tmp -= st.Query(ord[undo], ord[undo]+sz[undo]-1)
			} else {
				tmp = st.Query(ord[origin], ord[origin]+sz[origin]-1)
			}
			ans = append(ans, tmp)
		}
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

type SegTree struct {
	sz   int
	val  []int
	lazy []int
}

func NewSegTree(arr []int) *SegTree {
	n := len(arr)
	val := make([]int, n*4)
	lazy := make([]int, n*4)

	var build func(i int, l int, r int)
	build = func(i int, l int, r int) {
		if l == r {
			val[i] = arr[l]
			return
		}
		mid := (l + r) >> 1
		build(2*i, l, mid)
		build(2*i+1, mid+1, r)
		val[i] = val[2*i] + val[2*i+1]
	}

	build(1, 0, n-1)
	return &SegTree{n, val, lazy}
}

func (t *SegTree) pushLazy(i int, l int, r int, v int) {
	t.val[i] += (r - l + 1) * v
	t.lazy[i] += v
}

func (t *SegTree) push(i int, l int, r int) {
	mid := (l + r) / 2
	if t.lazy[i] != 0 {
		t.pushLazy(2*i, l, mid, t.lazy[i])
		t.pushLazy(2*i+1, mid+1, r, t.lazy[i])
		t.lazy[i] = 0
	}
}

func (t *SegTree) Update(L int, R int, v int) {
	var dfs func(i int, l int, r int)

	dfs = func(i int, l int, r int) {
		if R < l || r < L {
			return
		}

		if L <= l && r <= R {
			t.pushLazy(i, l, r, v)
			return
		}

		t.push(i, l, r)

		mid := (l + r) / 2
		dfs(2*i, l, mid)
		dfs(2*i+1, mid+1, r)
		t.val[i] = t.val[2*i] + t.val[2*i+1]
	}

	dfs(1, 0, t.sz-1)
}

func (t *SegTree) Query(L int, R int) int {
	var dfs func(i int, l int, r int) int

	dfs = func(i int, l int, r int) int {
		if R < l || r < L {
			return 0
		}
		if L <= l && r <= R {
			return t.val[i]
		}
		t.push(i, l, r)
		mid := (l + r) / 2
		return dfs(2*i, l, mid) + dfs(2*i+1, mid+1, r)
	}

	return dfs(1, 0, t.sz-1)
}

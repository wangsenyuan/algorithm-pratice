package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	n := readNum(reader)
	have := make([]int, n-1)
	limit := make([]int, n-1)
	for i := 0; i < n-1; i++ {
		have[i], limit[i] = readTwoNums(reader)
	}
	P := make([]int, n-1)
	for i := 0; i < n-1; i++ {
		P[i] = readNum(reader)
	}
	fmt.Println(solve(n, P, have, limit))
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

func prepend(v int, arr []int) []int {
	res := make([]int, len(arr)+1)
	res[0] = v
	copy(res[1:], arr)
	return res
}

func solve(n int, P []int, have []int, limit []int) int64 {
	have = prepend(0, have)
	limit = prepend(0, limit)
	g := NewGraph(n, n)
	for i := 1; i <= len(P); i++ {
		p := P[i-1] - 1
		g.AddEdge(p, i)
	}

	sz := make([]int, n)
	cnt := make([]int, n)
	depth := make([]int, n)
	var dfs1 func(u int) int

	dfs1 = func(u int) int {
		sz[u]++
		cnt[depth[u]]++
		var h int
		for i := g.nodes[u]; i > 0; i = g.next[i] {
			v := g.to[i]
			depth[v] = depth[u] + 1
			vh := dfs1(v)
			sz[u] += sz[v]
			if vh > h {
				h = vh
			}
		}
		return h + 1
	}
	// m is the tree depth
	m := dfs1(0)
	present := make([][]int, m)
	for i := 0; i < m; i++ {
		present[i] = make([]int, 0, cnt[i])
	}

	big := make([]bool, n)
	off := make([]bool, n)

	tree := NewTree(m)

	var add func(u int)

	add = func(u int) {
		tree.Add(depth[u], int64(have[u]))
		present[depth[u]] = append(present[depth[u]], u)
		for i := g.nodes[u]; i > 0; i = g.next[i] {
			v := g.to[i]
			if !big[v] && !off[v] {
				add(v)
			}
		}
	}

	var dfs2 func(u int, keep bool)

	dfs2 = func(u int, keep bool) {
		bc := -1
		for i := g.nodes[u]; i > 0; i = g.next[i] {
			v := g.to[i]
			if bc < 0 || sz[bc] < sz[v] {
				bc = v
			}
		}
		for i := g.nodes[u]; i > 0; i = g.next[i] {
			v := g.to[i]
			if v != bc {
				// not keep values form this child
				dfs2(v, false)
			}
		}
		if bc >= 0 {
			// keep value from bc
			dfs2(bc, true)
			big[bc] = true
		}

		add(u)
		// because it is children first, so, parent value will not affect this
		x := tree.GetFirstPositionExceedValue(int64(limit[u]))
		if x < m && u != 0 {
			// values won't be counted
			for i := x; i < m; i++ {
				if len(present[i]) == 0 {
					break
				}
				for _, v := range present[i] {
					tree.Add(depth[v], -int64(have[v]))
					off[v] = true
				}
				present[i] = present[i][:0]
			}
		}
		if bc >= 0 {
			big[bc] = false
		}
		if !keep {
			for i := depth[u]; i < m; i++ {
				if len(present[i]) == 0 {
					break
				}
				for _, v := range present[i] {
					tree.Add(i, -int64(have[v]))
				}
				present[i] = present[i][:0]
			}
		}
	}

	dfs2(0, true)

	var ans int64
	for i := 0; i < m; i++ {
		for _, v := range present[i] {
			ans += int64(have[v])
		}
	}

	return ans
}

type Tree struct {
	val []int64
	sz  int
}

func NewTree(n int) *Tree {
	t := new(Tree)
	t.val = make([]int64, 4*n)
	for i := 0; i < len(t.val); i++ {
		t.val[i] = 0
	}
	t.sz = n
	return t
}

func (t *Tree) Add(p int, v int64) {

	var loop func(i int, l, r int)

	loop = func(i int, l, r int) {
		if l == r {
			t.val[i] += v
			return
		}
		mid := (l + r) / 2
		if p <= mid {
			loop(2*i, l, mid)
		} else {
			loop(2*i+1, mid+1, r)
		}
		t.val[i] = max(t.val[2*i], t.val[i*2+1])
	}
	loop(1, 0, t.sz-1)
}

func (t *Tree) GetFirstPositionExceedValue(v int64) int {
	var loop func(i int, l, r int) int

	loop = func(i int, l, r int) int {
		if t.val[i] <= v {
			return t.sz
		}
		if l == r {
			return l
		}
		mid := (l + r) / 2
		if t.val[2*i] > v {
			return loop(2*i, l, mid)
		}
		return loop(2*i+1, mid+1, r)
	}
	return loop(1, 0, t.sz-1)
}

func max(a, b int64) int64 {
	if a >= b {
		return a
	}
	return b
}

type Graph struct {
	nodes []int
	next  []int
	to    []int
	cur   int
}

func NewGraph(n, e int) *Graph {
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

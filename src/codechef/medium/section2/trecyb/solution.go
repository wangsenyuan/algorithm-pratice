package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	n := readNum(reader)

	E := make([][]int, n-1)
	for i := 0; i < n-1; i++ {
		E[i] = readNNums(reader, 2)
	}

	fmt.Println(solve(n, E))
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

func solve(n int, E [][]int) int64 {
	g := NewGraph(n, len(E)*2)
	for _, e := range E {
		u, v := e[0], e[1]
		u--
		v--
		g.AddEdge(u, v)
		g.AddEdge(v, u)
	}

	depth := make([]int, n)

	cnt := make([]int, n)
	que := make([]int, n)
	var front, end int
	que[end] = 0
	depth[0] = 1
	cnt[0]++
	end++
	var mx int
	for front < end {
		u := que[front]
		front++
		mx = max(mx, depth[u])
		for i := g.nodes[u]; i > 0; i = g.next[i] {
			v := g.to[i]
			if depth[v] > 0 {
				continue
			}
			depth[v] = depth[u] + 1
			cnt[depth[v]-1] ++
			que[end] = v
			end++
		}
	}

	layer := make([][]int, mx)
	for i := 0; i < mx; i++ {
		layer[i] = make([]int, 0, cnt[i])
	}

	for i := 0; i < n; i++ {
		depth[i]--
		layer[depth[i]] = append(layer[depth[i]], i)
	}

	tree := NewSegTree(n)

	var beauty int64

	for l := 0; l < mx; l++ {
		for _, x := range layer[l] {
			beauty += int64(tree.Query(0, x))
		}
		for _, x := range layer[l] {
			tree.Update(x, 1)
		}
	}

	for i := 0; i < n; i++ {
		tree.Update(i, 0)
	}

	for l := mx - 1; l >= 0; l-- {
		for _, x := range layer[l] {
			beauty -= int64(tree.Query(0, x))
		}
		for _, x := range layer[l] {
			tree.Update(x, 1)
		}
	}
	return beauty
}

type SegTree []int

func NewSegTree(n int) SegTree {
	arr := make([]int, 2*n)
	return SegTree(arr)
}

func (tree SegTree) Update(p int, v int) {
	n := len(tree)
	n /= 2
	p += n
	tree[p] = v
	for p > 0 {
		tree[p>>1] = tree[p] + tree[p^1]
		p >>= 1
	}
}

func (tree SegTree) Query(l, r int) int {
	n := len(tree)
	n /= 2
	l += n
	r += n
	var res int
	for l < r {
		if l&1 == 1 {
			res += tree[l]
			l++
		}
		if r&1 == 1 {
			r--
			res += tree[r]
		}
		l >>= 1
		r >>= 1
	}
	return res
}

func max(a, b int) int {
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

func NewGraph(n int, e int) *Graph {
	g := new(Graph)
	g.nodes = make([]int, n)
	g.next = make([]int, e+10)
	g.to = make([]int, e+10)
	g.cur = 0
	return g
}

func (g *Graph) AddEdge(u, v int) {
	g.cur++
	g.next[g.cur] = g.nodes[u]
	g.nodes[u] = g.cur
	g.to[g.cur] = v
}

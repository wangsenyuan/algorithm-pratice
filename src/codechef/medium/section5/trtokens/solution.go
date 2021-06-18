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
		S, _ := reader.ReadString('\n')
		P := readNNums(reader, n-1)
		res := solve(n, S, P)
		buf.WriteString(fmt.Sprintf("%d\n", res))
	}
	fmt.Print(buf.String())
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

func solve(n int, S string, P []int) int64 {
	g := NewGraph(n, n-1)

	for i := 1; i < n; i++ {
		p := P[i-1] - 1
		g.AddEdge(p, i)
	}

	level := make([]int, n)
	st := make([]int, n)
	eu := make([]int, n)
	var time int
	var eulerTour func(u int, l int)

	eulerTour = func(u int, l int) {
		st[u] = time
		eu[time] = u
		time++
		level[u] = l
		for i := g.nodes[u]; i > 0; i = g.next[i] {
			v := g.to[i]
			eulerTour(v, l+1)
		}
	}

	eulerTour(0, 0)

	arr := make([]int, n)
	for i := 0; i < time; i++ {
		if S[eu[i]] == '1' {
			arr[i] = level[eu[i]]
		}
	}

	tree := NewSegTree(n, arr)

	sz := make([]int, n)

	var dfs func(u int)

	var res int64

	dfs = func(u int) {
		sz[u]++
		for i := g.nodes[u]; i > 0; i = g.next[i] {
			dfs(g.to[i])
			sz[u] += sz[g.to[i]]
		}

		if S[u] == '0' {
			l := st[u]
			r := l + sz[u]
			tmp := tree.Query(l, r)
			if tmp.first != 0 {
				res += int64(tmp.first - level[u])
				tree.Update(tmp.second, 0)
				tree.Update(st[u], level[u])
			}
		}
	}

	dfs(0)

	return res
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
	return g
}

func (g *Graph) AddEdge(u, v int) {
	g.cur++
	g.next[g.cur] = g.nodes[u]
	g.nodes[u] = g.cur
	g.to[g.cur] = v
}

type Pair struct {
	first  int
	second int
}

func (p Pair) Clone() Pair {
	return Pair{p.first, p.second}
}

func (p Pair) UpdateFirst(first int) Pair {
	return Pair{first, p.second}
}

type SegTree struct {
	arr []Pair
	n   int
}

func NewSegTree(n int, in []int) *SegTree {
	arr := make([]Pair, 2*n)

	for i := n; i < 2*n; i++ {
		arr[i] = Pair{in[i-n], i - n}
	}

	for i := n - 1; i > 0; i-- {
		l := i << 1
		r := l + 1
		if arr[l].first >= arr[r].first {
			arr[i] = arr[l].Clone()
		} else {
			arr[i] = arr[r].Clone()
		}
	}

	return &SegTree{arr, n}
}

func (t *SegTree) Query(l, r int) Pair {
	l += t.n
	r += t.n
	ans := Pair{0, -1}

	for l < r {
		if l&1 == 1 {
			if t.arr[l].first > ans.first {
				ans = t.arr[l]
			}
			l++
		}
		if r&1 == 1 {
			r--
			if t.arr[r].first > ans.first {
				ans = t.arr[r]
			}
		}
		l >>= 1
		r >>= 1
	}
	return ans
}

func (t *SegTree) Update(p int, v int) {
	p += t.n
	t.arr[p] = t.arr[p].UpdateFirst(v)

	for p > 0 {
		q := p ^ 1
		if t.arr[p].first >= t.arr[q].first {
			t.arr[p>>1] = t.arr[p].Clone()
		} else {
			t.arr[p>>1] = t.arr[q].Clone()
		}
		p >>= 1
	}
}

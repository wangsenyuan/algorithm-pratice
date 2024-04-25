package main

import (
	"bufio"
	"fmt"
	"os"
	"slices"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	n, m := readTwoNums(reader)
	people := make([][]int, n)

	for i := 0; i < n; i++ {
		people[i] = readNNums(reader, 2)
	}

	pairs := make([][]int, m)

	for i := 0; i < m; i++ {
		pairs[i] = readNNums(reader, 2)
	}

	res := solve(people, pairs)

	s := fmt.Sprintf("%v", res)

	fmt.Println(s[1 : len(s)-1])
}

func readString(r *bufio.Reader) string {
	s, _ := r.ReadBytes('\n')
	for i := 0; i < len(s); i++ {
		if s[i] == '\n' || s[i] == '\r' {
			return string(s[:i])
		}
	}
	return string(s)
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

func solve(people [][]int, avoids [][]int) []int {
	n := len(people)

	g := NewGraph(n, len(avoids)*2)

	for _, cur := range avoids {
		x, y := cur[0]-1, cur[1]-1
		g.AddEdge(x, y)
		g.AddEdge(y, x)
	}

	ps := make([]Pair, n)

	for i := 0; i < n; i++ {
		cur := people[i]
		ps[i] = Pair{cur[0] - cur[1], i}
	}

	// 按照 x[i] - y[i] 降序排
	slices.SortFunc(ps, func(a, b Pair) int {
		return b.first - a.first
	})

	pos := make([]int, n)
	for i := 0; i < n; i++ {
		pos[ps[i].second] = i
	}

	cnt := NewSegTree(n)
	sum := NewSegTree(n)

	for i := 0; i < n; i++ {
		j := ps[i].second
		cnt.Update(i, 1)
		// 先使用x[j]
		sum.Update(i, people[j][0])
	}

	res := make([]int, n)

	for pi, p := range ps {
		u := p.second

		for i := g.nodes[u]; i > 0; i = g.next[i] {
			v := g.to[i]
			cnt.Update(pos[v], 0)
			// set to 0
			sum.Update(pos[v], 0)
		}
		// 减去自己的贡献
		res[u] = cnt.Get(0, pi)*people[u][0] + cnt.Get(pi+1, n)*people[u][1] + sum.Get(0, n) - people[u][0]

		for i := g.nodes[u]; i > 0; i = g.next[i] {
			v := g.to[i]
			cnt.Update(pos[v], 1)
			if pos[v] < pi {
				// 前面使用的是y[i]
				sum.Update(pos[v], people[v][1])
			} else {
				// 后面使用的x[i]
				sum.Update(pos[v], people[v][0])
			}
		}
		// set to y[i]
		sum.Update(pi, people[u][1])
	}

	return res
}

type Pair struct {
	first  int
	second int
}

type SegTree struct {
	size int
	arr  []int
}

func NewSegTree(n int) *SegTree {
	arr := make([]int, 2*n)

	return &SegTree{n, arr}
}

func (seg *SegTree) Update(p int, v int) {
	p += seg.size
	seg.arr[p] = v
	for p > 1 {
		seg.arr[p>>1] = seg.arr[p] + seg.arr[p^1]
		p >>= 1
	}
}

func (seg *SegTree) Get(l, r int) int {
	var res int
	l += seg.size
	r += seg.size
	for l < r {
		if l&1 == 1 {
			res += seg.arr[l]
			l++
		}
		if r&1 == 1 {
			r--
			res += seg.arr[r]
		}
		l >>= 1
		r >>= 1
	}
	return res
}

type Graph struct {
	nodes []int
	next  []int
	to    []int
	cur   int
}

func NewGraph(n int, e int) *Graph {
	nodes := make([]int, n)
	e++
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

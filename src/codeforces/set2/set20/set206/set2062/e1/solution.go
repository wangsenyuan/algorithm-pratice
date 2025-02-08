package main

import (
	"bufio"
	"bytes"
	"fmt"
	"os"
	"slices"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	var buf bytes.Buffer
	tc := readNum(reader)
	for range tc {
		_, _, res := process(reader)
		buf.WriteString(fmt.Sprintf("%d\n", res))
	}
	buf.WriteTo(os.Stdout)
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

func process(reader *bufio.Reader) (w []int, edges [][]int, res int) {
	n := readNum(reader)
	w = readNNums(reader, n)
	edges = make([][]int, n-1)
	for i := range n - 1 {
		edges[i] = readNNums(reader, 2)
	}
	res = solve(w, edges)
	return
}

func solve(w []int, edges [][]int) int {
	n := len(w)

	g := NewGraph(n, 2*n)

	for _, e := range edges {
		u, v := e[0]-1, e[1]-1
		g.AddEdge(u, v)
		g.AddEdge(v, u)
	}

	sz := make([]int, n)
	in := make([]int, n)
	var pos int
	var dfs func(p int, u int)
	dfs = func(p int, u int) {
		sz[u]++
		in[u] = pos
		pos++
		for i := g.nodes[u]; i > 0; i = g.next[i] {
			v := g.to[i]
			if p != v {
				dfs(u, v)
				sz[u] += sz[v]
			}
		}
	}

	dfs(-1, 0)

	type pair struct {
		first  int
		second int
	}

	arr := make([]pair, n)
	for i := range n {
		arr[i] = pair{w[i], i}
	}

	slices.SortFunc(arr, func(a, b pair) int {
		return b.first - a.first
	})

	set := make(BIT, n+10)

	for i := 0; i < n; {
		j := i
		for i < n && arr[i].first == arr[j].first {
			u := arr[i].second
			cnt := set.getRange(in[u], in[u]+sz[u]-1)
			if cnt < j {
				// 有一个节点的weight比u大，且在它的子树范围外
				return u + 1
			}
			i++
		}
		for j < i {
			set.update(in[arr[j].second], 1)
			j++
		}
	}

	return 0
}

type BIT []int

func (set BIT) update(p int, v int) {
	p++
	for p < len(set) {
		set[p] += v
		p += p & -p
	}
}

func (set BIT) get(p int) int {
	p++
	var res int
	for p > 0 {
		res += set[p]
		p -= p & -p
	}
	return res
}

func (set BIT) getRange(l int, r int) int {
	res := set.get(r)
	if l > 0 {
		res -= set.get(l - 1)
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
	next := make([]int, e+3)
	to := make([]int, e+3)
	var cur int
	return &Graph{nodes, next, to, cur}
}

func (g *Graph) AddEdge(u, v int) {
	g.cur++
	g.next[g.cur] = g.nodes[u]
	g.nodes[u] = g.cur
	g.to[g.cur] = v
}

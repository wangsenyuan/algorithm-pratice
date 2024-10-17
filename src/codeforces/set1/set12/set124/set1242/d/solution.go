package main

import (
	"bufio"
	"fmt"
	"os"
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
		for x < len(bs) && (bs[x] < '0' || bs[x] > '9') && bs[x] != '-' {
			x++
		}
		x = readInt(bs, x, &res[i])
	}
	return res
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

func main() {
	reader := bufio.NewReader(os.Stdin)

	n, m := readTwoNums(reader)
	edges := make([][]int, m)
	for i := 0; i < m; i++ {
		edges[i] = readNNums(reader, 2)
	}
	res := solve(n, edges)

	fmt.Println(res)
}

func solve(n int, edges [][]int) int {
	g := NewGraph(n, len(edges))
	for _, e := range edges {
		u, v := e[0]-1, e[1]-1
		u, v = max(u, v), min(u, v)
		g.AddEdge(u, v)
	}

	set := NewDSU(n)
	active := make([]int, n)
	var top int
	cnt := make([]int, n)

	for u := 0; u < n; u++ {
		for i := g.nodes[u]; i > 0; i = g.next[i] {
			v := g.to[i]
			v = set.Find(v)
			cnt[v]++
		}
		var to_merge []int
		for i := 0; i < top; i++ {
			v := active[i]
			v = set.Find(v)
			if cnt[v] < set.cnt[v] {
				to_merge = append(to_merge, v)
			}
		}
		// reset cnt
		for i := g.nodes[u]; i > 0; i = g.next[i] {
			v := g.to[i]
			v = set.Find(v)
			cnt[v]--
		}
		if len(to_merge) == 0 {
			active[top] = u
			top++
			continue
		}
		// 这里要重新合并出active出来
		for _, v := range to_merge {
			set.Union(u, v)
		}
		var i int
		for j := 0; j < top; j++ {
			v := active[j]
			if set.Find(u) != set.Find(v) {
				active[i] = v
				i++
			}
		}
		active[i] = set.Find(u)
		top = i + 1
	}
	return top - 1
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
	e++
	g.next = make([]int, e)
	g.to = make([]int, e)
	g.cur = 0
	return g
}

func (g *Graph) AddEdge(u, v int) {
	g.cur++
	g.next[g.cur] = g.nodes[u]
	g.nodes[u] = g.cur
	g.to[g.cur] = v
}

type DSU struct {
	arr []int
	cnt []int
}

func NewDSU(n int) *DSU {
	arr := make([]int, n)
	cnt := make([]int, n)
	for i := 0; i < n; i++ {
		arr[i] = i
		cnt[i] = 1
	}
	return &DSU{arr, cnt}
}

func (this *DSU) Find(x int) int {
	if this.arr[x] != x {
		this.arr[x] = this.Find(this.arr[x])
	}
	return this.arr[x]
}

func (this *DSU) Union(x int, y int) bool {
	px := this.Find(x)
	py := this.Find(y)

	if px == py {
		return false
	}
	if this.cnt[px] < this.cnt[py] {
		px, py = py, px
	}
	this.cnt[px] += this.cnt[py]
	this.arr[py] = px
	return true
}

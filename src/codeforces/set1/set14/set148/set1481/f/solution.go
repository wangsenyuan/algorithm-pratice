package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	n, x := readTwoNums(reader)

	P := readNNums(reader, n-1)

	cnt, res := solve(n, x, P)

	fmt.Println(cnt)
	fmt.Println(res)
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

func solve(n int, x int, P []int) (int, string) {
	a := Pair{x, 'a'}
	b := Pair{n - x, 'b'}
	if a.first > b.first {
		a, b = b, a
	}

	g := NewGraph(n, n-1+10)

	for i := 0; i < len(P); i++ {
		g.AddEdge(P[i]-1, i+1)
	}
	sz := make([]int, n)
	dep := make([]int, n)
	var mx int
	var dfs func(u int, d int) int
	dfs = func(u int, d int) int {
		mx = max(mx, d)
		dep[u] = d
		sz[u]++
		for i := g.nodes[u]; i > 0; i = g.next[i] {
			sz[u] += dfs(g.to[i], d+1)
		}
		return sz[u]
	}

	dfs(0, 0)

	v := make([]*Pair, mx+1)
	for i := 0; i <= mx; i++ {
		v[i] = new(Pair)
		v[i].second = i
	}

	for i := 0; i < n; i++ {
		v[dep[i]].first++
	}

	cur := make([][]int, mx+1)
	for i := 0; i <= mx; i++ {
		cur[i] = make([]int, 0, v[i].first)
	}
	for i := 0; i < n; i++ {
		cur[dep[i]] = append(cur[dep[i]], i)
	}

	sort.Sort(Pairs(v))

	v2 := make([]*Pair, 0, len(v))

	for i := 0; i < len(v); i++ {
		if i == 0 || v[i].first != v[i-1].first {
			v2 = append(v2, &Pair{v[i].first, 1})
		} else {
			v2[len(v2)-1].second++
		}
	}

	dp := make([][]bool, len(v2)+1)
	for i := 0; i <= len(v2); i++ {
		dp[i] = make([]bool, n+1)
	}

	dp[len(v2)][0] = true
	last := make([]int, n+1)
	for i := len(v2) - 1; i >= 0; i-- {
		val := v2[i].first
		frq := v2[i].second
		for j := 0; j < val; j++ {
			last[j] = -1
		}
		for j := 0; j <= a.first; j++ {
			if dp[i+1][j] {
				last[j%val] = j
			}
			if last[j%val] < 0 || (j-last[j%val])/val > frq {
				dp[i][j] = false
			} else {
				dp[i][j] = true
			}
		}
	}

	F := make([]int, n+10)

	var buildPath func(i, j int)

	buildPath = func(i, j int) {
		if i == len(v2) {
			return
		}
		for !dp[i+1][j] {
			F[v2[i].first]++
			j -= v2[i].first
		}
		buildPath(i+1, j)
	}

	C := make([]byte, n)
	if dp[0][a.first] {
		buildPath(0, a.first)
		for i := 0; i < n; i++ {
			C[i] = byte(b.second)
		}
		for i := 0; i <= mx; i++ {
			if F[len(cur[i])] == 0 {
				continue
			}
			F[len(cur[i])]--
			for j := 0; j < len(cur[i]); j++ {
				C[cur[i][j]] = byte(a.second)
			}
		}
		mx++
	} else {
		for i := 0; i <= mx; i++ {
			sort.Sort(SortByTreeSize{cur[i], sz})
			if a.first < b.first {
				a, b = b, a
			}
			l := len(cur[i])
			for l > 0 && a.first > 0 {
				C[cur[i][l-1]] = byte(a.second)
				a.first--
				l--
			}
			for l > 0 && b.first > 0 {
				C[cur[i][l-1]] = byte(b.second)
				b.first--
				l--
			}
		}

		mx += 2
	}

	return mx, string(C)
}

type SortByTreeSize struct {
	nodes []int
	size  []int
}

func (this SortByTreeSize) Len() int {
	return len(this.nodes)
}

func (this SortByTreeSize) Less(i, j int) bool {
	a, b := this.nodes[i], this.nodes[j]
	return this.size[a] < this.size[b]
}

func (this SortByTreeSize) Swap(i, j int) {
	this.nodes[i], this.nodes[j] = this.nodes[j], this.nodes[i]
}

func max(a, b int) int {
	if a >= b {
		return a
	}
	return b
}

type Pair struct {
	first, second int
}

type Pairs []*Pair

func (this Pairs) Len() int {
	return len(this)
}

func (this Pairs) Less(i, j int) bool {
	return this[i].first < this[j].first
}

func (this Pairs) Swap(i, j int) {
	this[i], this[j] = this[j], this[i]
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
	return &Graph{nodes, next, to, 0}
}

func (g *Graph) AddEdge(u, v int) {
	g.cur++
	g.next[g.cur] = g.nodes[u]
	g.nodes[u] = g.cur
	g.to[g.cur] = v
}

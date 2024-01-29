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

	tc := readNum(reader)

	var buf bytes.Buffer

	for tc > 0 {
		tc--
		n, m := readTwoNums(reader)
		H := readNNums(reader, n)
		roads := make([][]int, m)
		for i := 0; i < m; i++ {
			roads[i] = readNNums(reader, 2)
		}
		k := readNum(reader)
		queries := make([][]int, k)
		for i := 0; i < k; i++ {
			queries[i] = readNNums(reader, 3)
		}
		res := solve(H, roads, queries)

		for _, b := range res {
			if b {
				buf.WriteString("YES\n")
			} else {
				buf.WriteString("NO\n")
			}
		}
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

func solve(H []int, roads [][]int, queries [][]int) []bool {
	n := len(H)
	g := NewGraph(n, len(roads)*2)
	for _, road := range roads {
		u, v := road[0]-1, road[1]-1
		if H[v] <= H[u] {
			g.AddEdge(u, v)
		} else {
			g.AddEdge(v, u)
		}
	}

	ids := make([]int, n)
	for i := 0; i < n; i++ {
		ids[i] = i
	}
	sort.Slice(ids, func(i, j int) bool {
		return H[ids[i]] < H[ids[j]]
	})

	qids := make([]int, len(queries))
	for i := 0; i < len(queries); i++ {
		qids[i] = i
	}

	sort.Slice(qids, func(i, j int) bool {
		a := queries[qids[i]][2] + H[queries[qids[i]][0]-1]
		b := queries[qids[j]][2] + H[queries[qids[j]][0]-1]
		return a <= b
	})

	dsu := NewDSU(n)

	ans := make([]bool, len(qids))
	for i, j := 0, 0; i < len(qids); i++ {
		qi := qids[i]
		cur := queries[qi]

		a, b, e := cur[0]-1, cur[1]-1, cur[2]

		for j < n && H[ids[j]] <= H[a]+e {
			u := ids[j]
			for it := g.nodes[u]; it > 0; it = g.next[it] {
				v := g.to[it]
				dsu.Union(u, v)
			}

			j++
		}

		if dsu.Find(a) == dsu.Find(b) {
			ans[qi] = true
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

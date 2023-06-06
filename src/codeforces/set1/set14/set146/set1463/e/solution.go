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

	tc := 1

	var buf bytes.Buffer

	for tc > 0 {
		tc--
		n, m := readTwoNums(reader)
		P := readNNums(reader, n)
		E := make([][]int, m)
		for i := 0; i < m; i++ {
			E[i] = readNNums(reader, 2)
		}
		res := solve(n, P, E)
		if len(res) == 0 {
			buf.WriteString("0\n")
			continue
		}
		for _, ans := range res {
			buf.WriteString(fmt.Sprintf("%d ", ans))
		}
		buf.WriteByte('\n')
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

func solve(n int, P []int, Q [][]int) []int {
	rt := -1
	g := NewGraph(n, n)
	for i := 0; i < n; i++ {
		j := P[i] - 1
		if j < 0 {
			rt = i
		} else {
			g.AddEdge(j, i)
		}
	}
	h := NewGraph(n, n)
	set := NewSet(n)
	for _, cur := range Q {
		x, y := cur[0]-1, cur[1]-1
		h.AddEdge(x, y)
		set.Union(x, y)
	}
	ord := h.TopoSort(n)
	if len(ord) == 0 {
		return nil
	}
	pos := make([]int, n)
	for i := 0; i < n; i++ {
		pos[ord[i]] = i
	}

	ng := NewGraph(n, n)

	for u := 0; u < n; u++ {
		for i := g.nodes[u]; i > 0; i = g.next[i] {
			v := g.to[i]
			pu := set.Find(u)
			pv := set.Find(v)
			if pu != pv {
				ng.AddEdge(pu, pv)
			}
		}
	}

	nord := ng.TopoSort(n)

	if len(nord) == 0 {
		return nil
	}

	comp := make([][]int, n)
	for i := 0; i < n; i++ {
		j := set.Find(i)
		if len(comp[j]) == 0 {
			comp[j] = make([]int, 0, 1)
		}
		comp[j] = append(comp[j], i)
	}

	var fin []int

	for _, p := range nord {
		sort.Slice(comp[p], func(i, j int) bool {
			return pos[comp[p][i]] < pos[comp[p][j]]
		})
		for _, v := range comp[p] {
			fin = append(fin, v)
		}
	}

	for i := 0; i < n; i++ {
		pos[fin[i]] = i
	}

	var dfs func(u int) bool

	dfs = func(u int) bool {
		for i := g.nodes[u]; i > 0; i = g.next[i] {
			v := g.to[i]
			if pos[v] < pos[u] {
				return false
			}
			if !dfs(v) {
				return false
			}
		}
		return true
	}

	if !dfs(rt) {
		return nil
	}
	for i := 0; i < n; i++ {
		fin[i]++
	}
	return fin
}

func reverse(arr []int) {
	for i, j := 0, len(arr)-1; i < j; i, j = i+1, j-1 {
		arr[i], arr[j] = arr[j], arr[i]
	}
}

type Graph struct {
	nodes []int
	next  []int
	to    []int
	cur   int
}

func NewGraph(n int, e int) *Graph {
	nodes := make([]int, n)
	next := make([]int, e+1)
	to := make([]int, e+1)
	return &Graph{nodes, next, to, 0}
}

func (g *Graph) AddEdge(u, v int) {
	g.cur++
	g.next[g.cur] = g.nodes[u]
	g.nodes[u] = g.cur
	g.to[g.cur] = v
}

func (g *Graph) TopoSort(n int) []int {
	vis := make([]int, n)

	ord := make([]int, 0, n)

	var dfs func(u int) bool
	dfs = func(u int) bool {
		if vis[u] == 2 {
			return true
		}
		vis[u]++
		for i := g.nodes[u]; i > 0; i = g.next[i] {
			v := g.to[i]
			if vis[v] == 1 || !dfs(v) {
				return false
			}
		}
		vis[u]++
		ord = append(ord, u)
		return true
	}

	for i := 0; i < n; i++ {
		if vis[i] == 0 && !dfs(i) {
			return nil
		}
	}

	reverse(ord)

	return ord
}

type Set struct {
	arr []int
	cnt []int
}

func NewSet(n int) *Set {
	arr := make([]int, n)
	cnt := make([]int, n)
	for i := 0; i < n; i++ {
		arr[i] = i
		cnt[i]++
	}
	return &Set{arr, cnt}
}

func (s *Set) Find(u int) int {
	if s.arr[u] != u {
		s.arr[u] = s.Find(s.arr[u])
	}
	return s.arr[u]
}

func (s *Set) Union(a, b int) bool {
	a = s.Find(a)
	b = s.Find(b)
	if a == b {
		return false
	}
	if s.cnt[a] < s.cnt[b] {
		a, b = b, a
	}
	s.cnt[a] += s.cnt[b]
	s.arr[b] = a
	return true
}

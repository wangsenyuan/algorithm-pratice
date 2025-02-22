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
	n, m := readTwoNums(reader)
	E := make([][]int, m)
	for i := 0; i < m; i++ {
		E[i] = readNNums(reader, 3)
	}
	res := solve(n, E)
	var buf bytes.Buffer
	for _, ans := range res {
		buf.WriteString(ans)
		buf.WriteByte('\n')
	}
	fmt.Print(buf.String())
}

func readFloat64(bs []byte, pos int, r *float64) int {
	var first float64
	var second float64
	exp := 1.0
	var dot bool
	for pos < len(bs) && (bs[pos] == '.' || isDigit(bs[pos])) {
		if bs[pos] == '.' {
			dot = true
		} else if !dot {
			first = first*10 + float64(bs[pos]-'0')
		} else {
			second = second*10 + float64(bs[pos]-'0')
			exp *= 10
		}
		pos++
	}
	*r = first + second/exp
	//fmt.Fprintf(os.Stderr, "%.6f ", *r)
	return pos
}

func isDigit(b byte) bool {
	return b >= '0' && b <= '9'
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

func readString(reader *bufio.Reader) string {
	s, _ := reader.ReadString('\n')
	for i := 0; i < len(s); i++ {
		if s[i] == '\n' || s[i] == '\r' {
			return s[:i]
		}
	}
	return s
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

func solve(n int, E [][]int) []string {
	// 这条边至少是和选中的某条边的长度相同
	m := len(E)
	edges := make([]Edge, m)

	for i := 0; i < m; i++ {
		edges[i] = Edge{E[i][0] - 1, E[i][1] - 1, E[i][2], i}
	}

	sort.Slice(edges, func(i, j int) bool {
		return edges[i].w < edges[j].w
	})

	labels := make([]string, m)

	for i := 0; i < m; i++ {
		labels[i] = "none"
	}

	low := make([]int, n)
	dfn := make([]int, n)

	type Pair struct {
		first  int
		second int
	}

	g := make([][]Pair, n)
	for i := 0; i < n; i++ {
		g[i] = make([]Pair, 0, 1)
	}

	var tarjan func(id int, u int, clock *int)

	tarjan = func(id int, u int, clock *int) {
		*clock++
		dfn[u] = *clock
		low[u] = dfn[u]
		for _, cur := range g[u] {
			v := cur.first
			if cur.second == id {
				continue
			}
			if dfn[v] == 0 {
				tarjan(cur.second, v, clock)
				low[u] = min(low[u], low[v])
				if low[v] > dfn[u] {
					labels[cur.second] = "any"
				}
			} else {
				low[u] = min(low[u], dfn[v])
			}
		}
	}

	addEdge := func(u, v, id int) {
		g[u] = append(g[u], Pair{v, id})
		g[v] = append(g[v], Pair{u, id})
	}

	set := NewSet(n)

	for i := 0; i < m; {
		j := i
		for i < m && edges[i].w == edges[j].w {
			i++
		}
		for k := j; k < i; k++ {
			u := set.Find(edges[k].u)
			v := set.Find(edges[k].v)
			if u != v {
				labels[edges[k].id] = "at least one"
				addEdge(u, v, edges[k].id)
				dfn[u] = 0
				dfn[v] = 0
			}
		}
		for k := j; k < i; k++ {
			u := set.Find(edges[k].u)
			v := set.Find(edges[k].v)
			if u != v && dfn[u] == 0 {
				tarjan(-1, u, new(int))
			}
		}
		for k := j; k < i; k++ {
			u := set.Find(edges[k].u)
			v := set.Find(edges[k].v)
			if u != v {
				set.Union(u, v)
			}
			g[u] = g[u][:0]
			g[v] = g[v][:0]
		}
	}

	return labels
}

func min(a, b int) int {
	if a <= b {
		return a
	}
	return b
}

type Edge struct {
	u  int
	v  int
	w  int
	id int
}

type Set struct {
	arr  []int
	cnt  []int
	size int
}

func NewSet(n int) *Set {
	arr := make([]int, n)
	cnt := make([]int, n)
	for i := 0; i < n; i++ {
		arr[i] = i
		cnt[i] = 1
	}
	return &Set{arr, cnt, n}
}

func (s *Set) Find(x int) int {
	if s.arr[x] != x {
		s.arr[x] = s.Find(s.arr[x])
	}
	return s.arr[x]
}

func (s *Set) Union(a, b int) bool {
	a = s.Find(a)
	b = s.Find(b)
	if a != b {
		if s.cnt[a] < s.cnt[b] {
			a, b = b, a
		}
		s.cnt[a] += s.cnt[b]
		s.arr[b] = a
		s.size--
		return true
	} else {
		return false
	}
}

func (s *Set) Size() int {
	return s.size
}

func (s *Set) Reset() {
	for i := 0; i < len(s.arr); i++ {
		s.arr[i] = i
		s.cnt[i] = 1
	}
	s.size = len(s.arr)
}

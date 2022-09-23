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
		n := readNum(reader)
		m := readNum(reader)
		conns := make([][]int, m)
		for i := 0; i < m; i++ {
			conns[i] = readNNums(reader, 2)
		}
		k, deps, heads := solve(n, m, conns)
		buf.WriteString(fmt.Sprintf("%d\n", k))
		for i := 0; i < n; i++ {
			buf.WriteString(fmt.Sprintf("%d ", deps[i]))
		}
		buf.WriteByte('\n')
		for i := 0; i < k; i++ {
			buf.WriteString(fmt.Sprintf("%d ", heads[i]))
		}
		buf.WriteByte('\n')
	}

	fmt.Print(buf.String())
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

func solve(n int, m int, conns [][]int) (int, []int, []int) {

	if m == n*(n-1)/2 {
		// in one department
		dept := make([]int, n)
		for i := 0; i < n; i++ {
			dept[i] = 1
		}
		return 1, dept, []int{1}
	}

	can := make([][]bool, n)

	for i := 0; i < n; i++ {
		can[i] = make([]bool, n)
	}
	deg := make([]int, n)
	for _, conn := range conns {
		u, v := conn[0], conn[1]
		u--
		v--
		can[u][v] = true
		can[v][u] = true
		deg[u]++
		deg[v]++
	}
	var most int
	for i := 0; i < n; i++ {
		if deg[i] > deg[most] {
			most = i
		}
	}

	var arr1 []int
	var arr2 []int

	for i := 0; i < n; i++ {
		if !can[most][i] {
			continue
		}
		if len(arr1) == 0 || can[i][arr1[0]] {
			arr1 = append(arr1, i)
		} else {
			arr2 = append(arr2, i)
		}
	}

	arr1 = append(arr1, most)
	sort.Ints(arr1)
	arr2 = append(arr2, most)
	sort.Ints(arr2)
	mark := make([]int, n)

	check := func(heads []int) bool {
		for i := 0; i < n; i++ {
			mark[i] = 0
		}
		id := 1
		for _, x := range heads {
			mark[x] = id
			id++
		}

		for i := 0; i < n; i++ {
			if mark[i] > 0 {
				continue
			}
			var found int
			for _, x := range heads {
				if can[i][x] {
					mark[i] = mark[x]
					found++
				}
			}
			if found != 1 {
				return false
			}
		}
		return true
	}
	heads := arr1

	if !check(arr1) {
		heads = arr2
		check(arr2)
	}
	for i := 0; i < len(heads); i++ {
		heads[i]++
	}

	return len(heads), mark, heads
}

type Set struct {
	arr []int
	cnt []int
}

func NewSet(n int) *Set {
	set := new(Set)
	set.arr = make([]int, n)
	set.cnt = make([]int, n)
	for i := 0; i < n; i++ {
		set.arr[i] = i
		set.cnt[i]++
	}
	return set
}

func (set *Set) Find(x int) int {
	if set.arr[x] != x {
		set.arr[x] = set.Find(set.arr[x])
	}
	return set.arr[x]
}

func (set *Set) Union(a, b int) bool {
	pa := set.Find(a)
	pb := set.Find(b)
	if pa == pb {
		return false
	}
	if set.cnt[pa] < set.cnt[pb] {
		pa, pb = pb, pa
	}
	set.cnt[pa] += set.cnt[pb]
	set.arr[pb] = pa
	return true
}

func (set *Set) Reset() {
	for i := 0; i < len(set.arr); i++ {
		set.arr[i] = i
		set.cnt[i] = 1
	}
}

type Graph struct {
	node []int
	next []int
	to   []int
	cur  int
}

func NewGraph(n, e int) *Graph {
	g := new(Graph)
	g.node = make([]int, n)
	g.to = make([]int, e+1)
	g.next = make([]int, e+1)
	g.cur = 0
	return g
}

func (g *Graph) AddEdge(u, v int) {
	g.cur++
	g.next[g.cur] = g.node[u]
	g.node[u] = g.cur
	g.to[g.cur] = v
}

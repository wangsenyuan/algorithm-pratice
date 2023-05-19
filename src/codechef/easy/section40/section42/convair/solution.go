package main

import (
	"bufio"
	"bytes"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	tc := readNum(reader)
	var buf bytes.Buffer

	for tc > 0 {
		tc--
		n, m := readTwoNums(reader)
		E := make([][]int, m)
		for i := 0; i < m; i++ {
			E[i] = readNNums(reader, 2)
		}
		cost, res := solve(n, E)
		//fmt.Printf("%d %d\n", cost, len(res))
		buf.WriteString(fmt.Sprintf("%d %d\n", cost, len(res)))
		for _, e := range res {
			buf.WriteString(fmt.Sprintf("%d %d\n", e[0], e[1]))
			//fmt.Printf("%d %d\n", e[0], e[1])
		}
	}

	fmt.Print(buf.String())
}

func readString(reader *bufio.Reader) string {
	s, _ := reader.ReadString('\n')
	for i := 0; i < len(s); i++ {
		if s[i] == '\n' {
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

func solve(n int, E [][]int) (int, [][]int) {
	set := NewDSU(n)
	var free [][]int
	var use [][]int

	deg := make([]int, n)

	for _, e := range E {
		e[0]--
		e[1]--
		u, v := e[0], e[1]

		deg[u]++
		deg[v]++

		if !set.Union(u, v) {
			free = append(free, e)
		} else {
			use = append(use, e)
		}
	}

	if len(free) > 1 {
		var ext [][]int
		first := free[0]
		for i := 1; i < len(free); i++ {
			a, b := first[0], first[1]
			c, d := free[i][0], free[i][1]
			if set.Find(a) != set.Find(c) {
				first[1] = c
				set.Union(b, d)
				use = append(use, []int{b, d})
			} else {
				ext = append(ext, free[i])
			}
		}
		free = append(ext, first)
	}

	in := make([]int, n)
	for i := 0; i < n; i++ {
		in[i] = -1
	}

	edge_count := make([]int, n)
	nodes := make([]map[int]bool, n)
	for i := 0; i < n; i++ {
		nodes[i] = make(map[int]bool)
	}
	for i, e := range use {
		u, v := e[0], e[1]
		p := set.Find(e[0])
		edge_count[p]++
		nodes[p][u] = true
		nodes[p][v] = true
		in[u] = i
		in[v] = i
	}

	for _, e := range free {
		u, v := e[0], e[1]
		p := set.Find(e[0])
		edge_count[p]++
		nodes[p][u] = true
		nodes[p][v] = true
	}

	var trees []int
	var singles []int

	for i := 0; i < n; i++ {
		j := set.Find(i)
		if edge_count[j]+1 == len(nodes[j]) {
			trees = append(trees, in[i])
			edge_count[j]++
		}
		if i == j && len(nodes[j]) == 0 {
			singles = append(singles, i)
		}
	}

	rem := make([]bool, len(use))
	var id int
	for ; id < len(free) && id < len(trees); id++ {
		a, b := free[id][0], free[id][1]
		j := trees[id]
		c, d := use[j][0], use[j][1]
		rem[j] = true
		use = append(use, []int{a, c})
		use = append(use, []int{b, d})
		set.Union(a, c)
	}

	for i := 0; id < len(free) && i+1 < len(singles); i += 2 {
		a, b := free[id][0], free[id][1]
		c := singles[i]
		d := singles[i+1]
		use = append(use, []int{a, c})
		use = append(use, []int{b, d})
		set.Union(a, c)
		set.Union(b, d)
		id++
	}

	for id < len(free) {
		use = append(use, free[id])
		id++
	}

	var roots []int

	for i := 0; i < n; i++ {
		p := set.Find(i)
		if p == i {
			roots = append(roots, p)
		}
	}

	for i := 1; i < len(roots); i++ {
		use = append(use, []int{roots[i-1], roots[i]})
		set.Union(roots[i-1], roots[i])
	}

	var m int
	for i := 0; i < len(use); i++ {
		if i >= len(rem) || !rem[i] {
			use[m] = use[i]
			m++
		}
	}

	use = use[:m]

	deg2 := make([]int, n)

	for _, e := range use {
		u, v := e[0], e[1]
		deg2[u]++
		deg2[v]++
		e[0]++
		e[1]++
	}

	var cost int

	for i := 0; i < n; i++ {
		cost += abs(deg[i] - deg2[i])
	}

	return cost, use
}

func abs(num int) int {
	if num < 0 {
		return -num
	}
	return num
}

type DSU struct {
	arr  []int
	cnt  []int
	size int
}

func NewDSU(n int) *DSU {
	set := new(DSU)
	set.arr = make([]int, n)
	set.cnt = make([]int, n)
	for i := 0; i < n; i++ {
		set.arr[i] = i
		set.cnt[i] = 1
	}
	set.size = n
	return set
}

func (set *DSU) Find(u int) int {
	if set.arr[u] != u {
		set.arr[u] = set.Find(set.arr[u])
	}
	return set.arr[u]
}

func (set *DSU) Union(a, b int) bool {
	a = set.Find(a)
	b = set.Find(b)
	if a == b {
		return false
	}
	if set.cnt[a] < set.cnt[b] {
		a, b = b, a
	}
	set.cnt[a] += set.cnt[b]
	set.arr[b] = a
	set.size--
	return true
}

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
		mat := make([]string, n)
		for i := 0; i < n; i++ {
			mat[i] = readString(reader)
		}
		res := solve(n, mat)
		buf.WriteString(fmt.Sprintf("%d\n", len(res)))
		for j := 0; j < len(res); j++ {
			buf.WriteString(fmt.Sprintf("%d", res[j]))
			if j == len(res)-1 {
				buf.WriteByte('\n')
			} else {
				buf.WriteByte(' ')
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

func normalize(s string) string {

	for i := len(s); i > 0; i-- {
		if s[i-1] >= 'a' && s[i-1] <= 'z' {
			return s[:i]
		}
	}
	return ""
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

func solve(n int, mat []string) []int {
	g := NewGraph(n, n*n)

	deg := make([]int, n)
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			if mat[i][j] == '1' {
				g.AddEdge(i, j)
				deg[i]++
			}
		}
	}

	// 如果已经在一个分组了，那么就是0
	// 如果存在多个分组，且这个分组的sz <= 2，那么其中最小的分组
	// 如果sz = 3, 且这个分组是一个圈，3也是一个答案
	// 如果存在sz >= 4的分组，且不是个完全图，那么答案都是1
	gid := make([]int, n)

	for i := 0; i < n; i++ {
		gid[i] = -1
	}

	var comps [][]int

	var dfs func(u int, id int)

	dfs = func(u int, id int) {
		gid[u] = id
		comps[id] = append(comps[id], u)
		for i := g.nodes[u]; i > 0; i = g.next[i] {
			v := g.to[i]
			if gid[v] < 0 {
				dfs(v, id)
			}
		}
	}

	for u := 0; u < n; u++ {
		if gid[u] < 0 {
			comps = append(comps, make([]int, 0, 1))
			dfs(u, len(comps)-1)
		}
	}
	m := len(comps)
	if m == 1 {
		return nil
	}

	check := func(comp []int) int {
		if len(comp) == 1 {
			return comp[0]
		}

		if len(comp) == 2 {
			return -1
		}

		// 在一个强联通分量中
		x := comp[0]
		for _, u := range comp {
			if deg[u] < deg[x] {
				x = u
			}
		}

		if deg[x] < len(comp)-1 {
			return x
		}

		return -1
	}

	// 先找答案1
	for _, comp := range comps {
		u := check(comp)
		if u >= 0 {
			return []int{u + 1}
		}
	}
	// find answer 2
	for _, comp := range comps {
		if len(comp) == 2 {
			return []int{comp[0] + 1, comp[1] + 1}
		}
	}
	if len(comps) >= 3 {
		a := comps[0][0]
		b := comps[1][0]
		return []int{a + 1, b + 1}
	}

	// find the smallest one
	sort.Slice(comps, func(i, j int) bool {
		return len(comps[i]) < len(comps[j])
	})

	ans := make([]int, len(comps[0]))
	for i := 0; i < len(comps[0]); i++ {
		ans[i] = comps[0][i] + 1
	}

	return ans
}

type Pair struct {
	first  int
	second int
}

func min(a, b int) int {
	if a <= b {
		return a
	}
	return b
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

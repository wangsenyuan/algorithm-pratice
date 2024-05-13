package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	n := readNum(reader)
	p := readNNums(reader, n-1)
	s := readNNums(reader, n)

	res := solve(n, p, s)

	fmt.Println(res)
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

const oo = 1 << 30

func solve(n int, p []int, s []int) int {
	g := NewGraph(n, n)

	for i := 0; i < n-1; i++ {
		g.AddEdge(p[i]-1, i+1)
	}

	ans := make([]int, n)

	var dfs func(u int, h int) bool

	dfs = func(u int, h int) bool {
		if h%2 == 0 {
			if g.nodes[u] == 0 {
				ans[u] = 0
				return true
			}
			sp := s[p[u-1]-1]
			ans[u] = oo
			// it has leaf
			for i := g.nodes[u]; i > 0; i = g.next[i] {
				// 子节点肯定在奇数层
				v := g.to[i]
				// 保证 a[v] >= 0
				ans[u] = min(ans[u], s[v]-sp)
			}
			if ans[u] < 0 {
				return false
			}
			s[u] = sp + ans[u]
		}

		for i := g.nodes[u]; i > 0; i = g.next[i] {
			v := g.to[i]
			if h%2 == 0 {
				ans[v] = s[v] - s[u]
			}
			if !dfs(v, h+1) {
				return false
			}
		}

		return true
	}

	ans[0] = s[0]

	ok := dfs(0, 1)

	if !ok {
		return -1
	}

	var sum int

	for i := 0; i < n; i++ {
		sum += ans[i]
	}
	return sum
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

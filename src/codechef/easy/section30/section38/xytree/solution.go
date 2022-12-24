package main

import (
	"bufio"
	"bytes"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	var buf bytes.Buffer

	tc := readNum(reader)

	for tc > 0 {
		tc--
		n := readNum(reader)
		E := make([][]int, n-1)
		for i := 0; i < n-1; i++ {
			E[i] = readNNums(reader, 2)
		}
		m := readNum(reader)
		Q := make([][]int, m)
		for i := 0; i < m; i++ {
			s, _ := reader.ReadBytes('\n')
			var x int
			pos := readInt(s, 0, &x)
			if x == 1 {
				var u int
				readInt(s, pos+1, &u)
				Q[i] = []int{x, u}
			} else {
				Q[i] = []int{2}
			}
		}
		res := solve(n, E, Q)
		for _, ans := range res {
			buf.WriteString(fmt.Sprintf("%d\n", ans))
		}
	}

	fmt.Print(buf.String())
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
func solve(n int, E [][]int, Q [][]int) []int {
	g := NewGraph(n, 2*n)

	for _, e := range E {
		u, v := e[0], e[1]
		u--
		v--
		g.AddEdge(u, v)
		g.AddEdge(v, u)
	}

	deg := make([]int, n)
	P := make([]int, n)

	var dfs func(p int, u int)
	dfs = func(p int, u int) {
		P[u] = p
		for i := g.nodes[u]; i > 0; i = g.next[i] {
			v := g.to[i]
			if p != v {
				deg[u]++
				dfs(u, v)
			}
		}
	}

	dfs(-1, 0)

	var ans []int
	cur := n - 1

	val := make([]bool, n)

	for _, q := range Q {
		if q[0] == 1 {
			u := q[1]
			u--
			if deg[u] == 0 && !val[u] {
				val[u] = true
				cur--
				if P[u] >= 0 {
					// 连接 P[u] 和 u 的边变成了bad
					deg[P[u]]--
				}
			}
		} else {
			if !val[0] {
				ans = append(ans, cur)
			} else {
				ans = append(ans, n-1)
			}
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

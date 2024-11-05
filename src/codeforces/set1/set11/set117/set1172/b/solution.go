package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	n := readNum(reader)
	lines := make([][]int, n-1)
	for i := 0; i < n-1; i++ {
		lines[i] = readNNums(reader, 2)
	}
	res := solve(n, lines)

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

const mod = 998244353

func add(a, b int) int {
	a += b
	if a >= mod {
		a -= mod
	}
	return a
}

func mul(a, b int) int {
	return a * b % mod
}

func solve(n int, lines [][]int) int {
	f := make([]int, n+1)
	f[0] = 1
	for i := 1; i <= n; i++ {
		f[i] = mul(i, f[i-1])
	}

	g := NewGraph(n, 2*n)

	for _, line := range lines {
		u, v := line[0]-1, line[1]-1
		g.AddEdge(u, v)
		g.AddEdge(v, u)
	}

	var dfs func(p int, u int) int

	dfs = func(p int, u int) int {
		var cnt int
		res := 1
		for i := g.nodes[u]; i > 0; i = g.next[i] {
			v := g.to[i]
			if v != p {
				res = mul(res, dfs(u, v))
				cnt++
			}
		}
		if u != 0 {
			// u不能在前后，可以在子节点的中间位置
			cnt++
		}
		return mul(res, f[cnt])
	}

	ans := dfs(-1, 0)

	return mul(ans, n)
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

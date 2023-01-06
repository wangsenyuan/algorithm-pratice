package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	n := readNum(reader)
	E := make([][]int, n-1)
	for i := 0; i < n-1; i++ {
		E[i] = readNNums(reader, 2)
	}
	res := solve(n, E)
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

func solve(n int, E [][]int) int {

	g := NewGraph(n, 2*n)

	for _, e := range E {
		u, v := e[0]-1, e[1]-1
		g.AddEdge(u, v)
		g.AddEdge(v, u)
	}

	dp := make([]int, n)
	hp := make([]int, n)

	marked := make([]bool, g.cur+1)

	var dfs2 func(p int, u int) Pair

	dfs2 = func(p int, u int) Pair {
		res := Pair{u, 0}
		for i := g.node[u]; i > 0; i = g.next[i] {
			if !marked[i] && g.to[i] != p {
				res = res.Max(dfs2(u, g.to[i]))
			}
		}
		return Pair{res.first, res.second + 1}
	}

	var ans int

	var dfs1 func(p int, u int)

	dfs1 = func(p int, u int) {
		big := -1
		for i := g.node[u]; i > 0; i = g.next[i] {
			v := g.to[i]
			if p != v {
				dfs1(u, v)
				dp[u] = max(dp[u], dp[v])
				hp[u] = max(hp[u], hp[v])
				if big < 0 || hp[big] < hp[v] {
					big = v
				}
				marked[i] = true
				a := dfs2(-1, u)

				if a.second > 1 {
					b := dfs2(-1, a.first)
					ans = max(ans, (dp[v]-1)*(b.second-1))
				}

				marked[i] = false
			}
		}

		hp[u]++
		dp[u] = max(dp[u], hp[u])

		if big > 0 {
			for i := g.node[u]; i > 0; i = g.next[i] {
				v := g.to[i]
				if p != v && v != big {
					dp[u] = max(dp[u], hp[big]+1+hp[v])
				}
			}
		}
	}

	dfs1(-1, 0)

	return ans
}

type Pair struct {
	first  int
	second int
}

func (this Pair) Max(that Pair) Pair {
	if this.second > that.second {
		return this
	}
	return that
}

func max(a, b int) int {
	if a >= b {
		return a
	}
	return b
}

type Graph struct {
	node []int
	next []int
	to   []int
	cur  int
}

func NewGraph(n int, e int) *Graph {
	g := new(Graph)
	g.node = make([]int, n)
	g.next = make([]int, e+2)
	g.to = make([]int, e+2)
	g.cur = 1
	return g
}

func (g *Graph) AddEdge(u, v int) {
	g.cur++
	g.next[g.cur] = g.node[u]
	g.node[u] = g.cur
	g.to[g.cur] = v
}

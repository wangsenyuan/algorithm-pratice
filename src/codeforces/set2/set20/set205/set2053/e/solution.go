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
		res := process(reader)
		buf.WriteString(fmt.Sprintf("%d\n", res))
	}

	fmt.Print(buf.String())
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

func process(reader *bufio.Reader) int {
	n := readNum(reader)
	edges := make([][]int, n-1)
	for i := range n - 1 {
		edges[i] = readNNums(reader, 2)
	}
	return solve(n, edges)
}

const inf = 1 << 60

func solve(n int, edges [][]int) int {
	if n <= 2 {
		return 0
	}
	g := NewGraph(n, 2*n)
	deg := make([]int, n)
	for _, e := range edges {
		u, v := e[0]-1, e[1]-1
		g.AddEdge(u, v)
		g.AddEdge(v, u)
		deg[u]++
		deg[v]++
	}

	var c1 int
	// dp[u] 表示它的不是叶子节点的领居的个数
	dp := make([]int, n)
	for u := 0; u < n; u++ {
		if deg[u] == 1 {
			c1++
		} else {
			// 内部节点
			for i := g.nodes[u]; i > 0; i = g.next[i] {
				v := g.to[i]
				if deg[v] > 1 {
					dp[u]++
				}
			}
		}
	}

	ans := c1 * (n - c1)

	var c2 int
	for u := 0; u < n; u++ {
		if deg[u] > 1 && deg[u] == dp[u] {
			// u在一次移动后，无法到到叶子节点
			c2++
		}
	}

	for u := 0; u < n; u++ {
		if deg[u] > 1 && deg[u] != dp[u] {
			// 假设经过一次操作后，q到到了u点，它有叶子节点的领居
			// 那么它就可以移动过去
			// dp[u] - 1 是排除掉p到u的路径上的点，其他的邻居（非叶子节点）都会往前移动到u
			ans += c2 * (dp[u] - 1)
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
	g := new(Graph)
	g.nodes = make([]int, n)
	e++
	g.next = make([]int, e)
	g.to = make([]int, e)
	g.cur = 0
	return g
}

func (g *Graph) AddEdge(u, v int) {
	g.cur++
	g.next[g.cur] = g.nodes[u]
	g.nodes[u] = g.cur
	g.to[g.cur] = v
}

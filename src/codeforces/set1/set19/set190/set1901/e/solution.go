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
		n := readNum(reader)

		s, _ := reader.ReadBytes('\n')
		var pos int
		for i := 0; i < n; i++ {
			pos = readInt(s, pos, &a[i]) + 1
		}
		for i := 0; i < n-1; i++ {
			s, _ = reader.ReadBytes('\n')
			pos = readInt(s, 0, &edges[i][0])
			readInt(s, pos+1, &edges[i][1])
		}

		res := solve(a[:n], edges[:n-1])
		buf.WriteString(fmt.Sprintf("%d\n", res))
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

const oo = 1 << 60
const N = 500010

var g *Graph
var dp [N]int
var a [N]int
var edges [N][]int

func init() {
	g = NewGraph(N, 2*N)
	for i := 0; i < N; i++ {
		edges[i] = make([]int, 2)
	}
}

func solve(a []int, edges [][]int) int {
	n := len(a)

	if n == 2 {
		return max(0, a[0], a[1], a[0]+a[1])
	}

	for _, e := range edges {
		u, v := e[0]-1, e[1]-1
		g.AddEdge(u, v)
		g.AddEdge(v, u)
	}
	var ans int

	var dfs func(p int, u int)

	dfs = func(p int, u int) {
		sum := []int{-oo, -oo, -oo, -oo}
		sum[0] = 0
		for i := g.nodes[u]; i > 0; i = g.next[i] {
			v := g.to[i]
			if p != v {
				dfs(u, v)
				for j := 3; j >= 0; j-- {
					sum[min(j+1, 3)] = max(sum[min(j+1, 3)], sum[j]+dp[v])
				}
			}
		}

		dp[u] = -oo

		for j := 0; j < 4; j++ {
			dp[u] = max(dp[u], sum[j]+nvl(j == 1, 0, a[u]))
			ans = max(ans, sum[j]+nvl(j == 2, 0, a[u]))
		}
	}

	dfs(-1, 0)

	g.Reset(n)

	return ans
}

func nvl(b bool, first int, second int) int {
	if b {
		return first
	}
	return second
}

func min(a, b int) int {
	if a <= b {
		return a
	}
	return b
}

func max(arr ...int) int {
	res := arr[0]
	for i := 1; i < len(arr); i++ {
		if res < arr[i] {
			res = arr[i]
		}
	}
	return res
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

func (g *Graph) Reset(n int) {
	for i := 0; i < n; i++ {
		g.nodes[i] = 0
	}
	g.cur = 0
}

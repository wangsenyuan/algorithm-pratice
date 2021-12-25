package main

import (
	"bufio"
	"bytes"
	"fmt"
	"os"
)

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

func main() {
	reader := bufio.NewReader(os.Stdin)
	var buf bytes.Buffer

	tc := readNum(reader)

	for tc > 0 {
		tc--
		n, k := readTwoNums(reader)
		P := readNNums(reader, n-1)
		res := solve(n, k, P)
		buf.WriteString(fmt.Sprintf("%d\n", res))
	}

	fmt.Print(buf.String())
}

const INF = 1 << 30

func solve(n, k int, P []int) int {
	g := NewGraph(n, len(P))

	for i := 1; i < n; i++ {
		g.AddEdge(P[i-1]-1, i)
	}

	dp := make([][]int, n)
	for i := 0; i < n; i++ {
		dp[i] = make([]int, k+1)
		for j := 0; j <= k; j++ {
			dp[i][j] = INF
		}
	}

	var dfs func(u int) int

	dfs = func(u int) int {
		dp[u][0] = 0
		sz := 1
		var tmp int

		for i := g.nodes[u]; i > 0; i = g.next[i] {
			v := g.to[i]
			zs := dfs(v)
			sz += zs

			for a := tmp; a >= 0; a-- {
				if dp[u][a] == INF {
					continue
				}
				for b := 1; b <= zs; b++ {
					if a+b > k {
						break
					}
					if dp[v][b] == INF {
						continue
					}

					dp[u][a+b] = min(dp[u][a+b], dp[u][a]+dp[v][b])
					tmp = max(tmp, a+b)
				}
			}
		}
		if sz <= k {
			dp[u][sz] = 1
		}
		return sz
	}

	dfs(0)

	fp := make([]int, k+1)
	fp[0] = 0
	for x := 1; x <= k; x++ {
		fp[x] = INF
		for i := 1; i <= min(n, x); i++ {
			fp[x] = min(fp[x], 1+dp[0][i]+fp[x-i])
		}
	}

	return fp[k] - 1
}

func max(a, b int) int {
	if a >= b {
		return a
	}
	return b
}

func min(a, b int) int {
	if a <= b {
		return a
	}
	return b
}

func search(n int, fn func(int) bool) int {
	left, right := 0, n
	for left < right {
		mid := (left + right) / 2
		if fn(mid) {
			right = mid
		} else {
			left = mid + 1
		}
	}
	return right
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
	g.next = make([]int, n)
	g.to = make([]int, n)
	g.cur = 0
	return g
}

func (g *Graph) AddEdge(u, v int) {
	g.cur++
	g.next[g.cur] = g.nodes[u]
	g.nodes[u] = g.cur
	g.to[g.cur] = v
}

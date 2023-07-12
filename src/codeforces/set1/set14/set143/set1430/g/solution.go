package main

import (
	"bufio"
	"bytes"
	"fmt"
	"math"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	n, m := readTwoNums(reader)
	E := make([][]int, m)
	for i := 0; i < m; i++ {
		E[i] = readNNums(reader, 3)
	}
	res := solve(n, E)
	var buf bytes.Buffer
	for i := 0; i < n; i++ {
		buf.WriteString(fmt.Sprintf("%d ", res[i]))
	}
	buf.WriteByte('\n')
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

const inf = math.MaxInt64

func solve(n int, E [][]int) []int {
	// 0...n-1
	g := NewGraph(n, len(E))
	sum := make([]int64, n)
	for _, e := range E {
		u, v, w := e[0]-1, e[1]-1, e[2]
		g.AddEdge(u, v)
		sum[u] += int64(w)
		sum[v] -= int64(w)
	}

	vis := make([]bool, n)
	ord := make([]int, 0, n)

	var dfs1 func(u int)
	dfs1 = func(u int) {
		vis[u] = true
		for i := g.nodes[u]; i > 0; i = g.next[i] {
			v := g.to[i]
			if !vis[v] {
				dfs1(v)
			}
		}
		ord = append(ord, u)
	}

	for i := 0; i < n; i++ {
		if !vis[i] {
			dfs1(i)
		}
	}
	reverse(ord)

	for i := 0; i < n; i++ {
		vis[i] = false
	}

	var dfs2 func(u int)
	dfs2 = func(u int) {
		vis[u] = true
		for i := g.nodes[u]; i > 0; i = g.next[i] {
			v := g.to[i]
			if !vis[v] {
				dfs2(v)
			}
		}
	}
	need := make([]int, n)
	for i := 0; i < n; i++ {
		dfs2(i)
		for j := 0; j < n; j++ {
			if j != i && vis[j] {
				need[i] |= 1 << j
			}
			vis[j] = false
		}
	}

	dp := make([][][]int64, n+1)
	fp := make([][][]bool, n+1)
	for i := 0; i <= n; i++ {
		dp[i] = make([][]int64, n+1)
		fp[i] = make([][]bool, n+1)
		for j := 0; j <= n; j++ {
			dp[i][j] = make([]int64, 1<<n)
			fp[i][j] = make([]bool, 1<<n)
			for k := 0; k < len(dp[i][j]); k++ {
				dp[i][j][k] = inf
			}
		}
	}
	dp[0][0][0] = 0

	for i := 0; i < n; i++ {
		for j := 0; j <= n; j++ {
			for k := 0; k < (1 << n); k++ {
				if dp[i][j][k] > inf/2 {
					continue
				}
				if j == n {
					if dp[i+1][0][k] > dp[i][j][k] {
						dp[i+1][0][k] = dp[i][j][k]
						fp[i+1][0][k] = false
					}
				} else {
					if dp[i][j+1][k] > dp[i][j][k] {
						dp[i][j+1][k] = dp[i][j][k]
						fp[i][j+1][k] = false
					}

					v := ord[j]
					add := sum[v] * int64(i)
					if (k>>v)&1 == 0 && need[v]&k == need[v] {
						nk := k | (1 << v)
						if dp[i][j+1][nk] > dp[i][j][k]+add {
							dp[i][j+1][nk] = dp[i][j][k] + add
							fp[i][j+1][nk] = true
						}
					}
				}
			}
		}
	}

	ans := make([]int, n)
	for i, j, k := n, 0, (1<<n)-1; i > 0 || j > 0 || k > 0; {
		if j == 0 {
			j = n
			i--
		} else {
			if fp[i][j][k] {
				v := ord[j-1]
				ans[v] = i
				k ^= (1 << v)
			}
			j--
		}
	}
	return ans
}

func reverse(arr []int) {
	for i, j := 0, len(arr)-1; i < j; i, j = i+1, j-1 {
		arr[i], arr[j] = arr[j], arr[i]
	}
}

type Graph struct {
	nodes []int
	next  []int
	to    []int
	cur   int
}

func NewGraph(n int, e int) *Graph {
	nodes := make([]int, n)
	next := make([]int, e+1)
	to := make([]int, e+1)
	return &Graph{nodes, next, to, 0}
}

func (g *Graph) AddEdge(u, v int) {
	g.cur++
	g.next[g.cur] = g.nodes[u]
	g.nodes[u] = g.cur
	g.to[g.cur] = v
}

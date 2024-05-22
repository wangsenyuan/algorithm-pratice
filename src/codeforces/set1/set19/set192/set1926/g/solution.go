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
		a := readNNums(reader, n-1)
		s := readString(reader)
		res := solve(n, a, s)

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

const inf = 1 << 30

func solve(n int, a []int, s string) int {
	g := NewGraph(n, n*2)

	for i := 0; i < n-1; i++ {
		j := a[i] - 1
		g.AddEdge(j, i+1, i)
	}

	get := func(u int) int {
		if s[u] == 'P' {
			return -1
		}
		if s[u] == 'S' {
			return 1
		}
		return 0
	}

	dp := make([][]int, n)

	var dfs func(u int)
	dfs = func(u int) {
		dp[u] = []int{inf, inf, inf}
		if get(u) != 1 {
			dp[u][1] = 0
		}
		if get(u) != -1 {
			dp[u][2] = 0
		}
		var tot int

		for i := g.nodes[u]; i > 0; i = g.next[i] {
			v := g.to[i]
			dfs(v)
			dp[u][1] += min(dp[v][1], dp[v][2]+1, dp[v][0])
			dp[u][2] += min(dp[v][2], dp[v][1]+1, dp[v][0])
			tot += dp[v][0]
		}
		if get(u) != 0 {
			tot = inf
		}
		dp[u][0] = min(tot, dp[u][1]+1, dp[u][2]+1)
	}

	dfs(0)

	return min(dp[0][0], dp[0][1], dp[0][2])
}

func solve1(n int, a []int, s string) int {
	g := NewGraph(n, n*2)

	for i := 0; i < n-1; i++ {
		j := a[i] - 1
		g.AddEdge(i+1, j, i)
		g.AddEdge(j, i+1, i)
	}

	// 找到一个P的节点作为root
	root := -1
	for i := 0; i < n; i++ {
		if s[i] == 'P' {
			root = i
			break
		}
	}

	if root < 0 {
		return 0
	}

	get := func(u int) int {
		if s[u] == 'P' {
			return -1
		}
		if s[u] == 'S' {
			return 1
		}
		return 0
	}

	var dfs func(p int, u int) int

	var res int

	dfs = func(p int, u int) int {
		x := get(u)
		cnt := []int{0, 0}
		for i := g.nodes[u]; i > 0; i = g.next[i] {
			v := g.to[i]
			if p != v {
				tmp := dfs(u, v)
				if get(u) == 0 {
					if tmp > 0 {
						cnt[0]++
					} else if tmp < 0 {
						cnt[1]++
					}
				} else if sign(x)*sign(tmp) == -1 {
					res++
				}
			}
		}

		if x == 0 {
			res += min(cnt[0], cnt[1])
			if cnt[0] < cnt[1] {
				return -1
			} else if cnt[0] > cnt[1] {
				return 1
			}
		}

		return x
	}

	dfs(-1, root)

	return res
}

func sign(num int) int {
	if num < 0 {
		return -1
	}
	if num > 0 {
		return 1
	}
	return 0
}

type Graph struct {
	nodes []int
	next  []int
	to    []int
	val   []int
	cur   int
}

func NewGraph(n int, e int) *Graph {
	nodes := make([]int, n)
	next := make([]int, e)
	to := make([]int, e)
	val := make([]int, e)
	return &Graph{nodes, next, to, val, 0}
}

func (g *Graph) AddEdge(u, v, w int) {
	g.cur++
	g.next[g.cur] = g.nodes[u]
	g.nodes[u] = g.cur
	g.to[g.cur] = v
	g.val[g.cur] = w
}

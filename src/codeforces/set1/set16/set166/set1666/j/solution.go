package main

import (
	"bufio"
	"bytes"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	n := readNum(reader)
	c := make([][]int, n)
	for i := 0; i < n; i++ {
		c[i] = readNNums(reader, n)
	}

	res := solve(n, c)

	var buf bytes.Buffer

	for i := 0; i < n; i++ {
		buf.WriteString(fmt.Sprintf("%d ", res[i]))
	}

	fmt.Println(buf.String())
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

type Pair struct {
	first  int
	second int
}

const inf = 1 << 50

func solve(n int, c [][]int) []int {
	sum := make([][]int, n)
	for i := 0; i < n; i++ {
		sum[i] = make([]int, n)
		for j := 0; j < n; j++ {
			sum[i][j] = c[i][j]
			if i > 0 {
				sum[i][j] += sum[i-1][j]
			}
			if j > 0 {
				sum[i][j] += sum[i][j-1]
			}
			if i > 0 && j > 0 {
				sum[i][j] -= sum[i-1][j-1]
			}
		}
	}

	get := func(a, b, c, d int) int {
		if a > c || b > d {
			return 0
		}
		res := sum[c][d]
		if a > 0 {
			res -= sum[a-1][d]
		}
		if b > 0 {
			res -= sum[c][b-1]
		}
		if a > 0 && b > 0 {
			res += sum[a-1][b-1]
		}
		return res
	}

	dp := make([][]int, n)
	next := make([][]int, n)
	for i := 0; i < n; i++ {
		dp[i] = make([]int, n)
		next[i] = make([]int, n)
		for j := 0; j < n; j++ {
			dp[i][j] = -1
		}
	}

	var f func(l int, r int) int
	f = func(l int, r int) int {
		if l > r {
			return 0
		}
		res := &dp[l][r]
		if *res >= 0 {
			return *res
		}
		best := inf
		next[l][r] = -1
		for u := l; u <= r; u++ {
			x := f(l, u-1) + get(l, 0, u-1, l-1) + get(l, u, u-1, n-1)
			y := f(u+1, r) + get(u+1, 0, r, u) + get(u+1, r+1, r, n-1)
			if x+y < best {
				best = x + y
				next[l][r] = u
			}
		}
		*res = best
		return best
	}

	f(0, n-1)

	ans := make([]int, n)

	var dfs func(l int, r int, p int)

	dfs = func(l int, r int, p int) {
		if l > r {
			return
		}
		u := next[l][r]
		ans[u] = p + 1
		dfs(l, u-1, u)
		dfs(u+1, r, u)
	}

	dfs(0, n-1, -1)

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

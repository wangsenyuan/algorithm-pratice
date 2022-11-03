package main

import (
	"bufio"
	"bytes"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	n, mod := readTwoNums(reader)

	edges := make([][]int, n-1)

	for i := 0; i < n-1; i++ {
		edges[i] = readNNums(reader, 2)
	}

	res := solve(n, mod, edges)

	var buf bytes.Buffer

	for i := 0; i < n-1; i++ {
		buf.WriteString(fmt.Sprintf("%d ", res[i]))
	}

	buf.WriteByte('\n')
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

func normalize(s string) string {

	for i := len(s); i > 0; i-- {
		if s[i-1] >= 'a' && s[i-1] <= 'z' {
			return s[:i]
		}
	}
	return ""
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

func modAdd(a, b int, p int) int {
	a += b
	if a >= p {
		a -= p
	}
	return a
}

func modMul(a, b int, p int) int {
	r := int64(a) * int64(b)
	r %= int64(p)
	return int(r)
}

func pow(a, b int, p int) int {
	r := 1
	for b > 0 {
		if b&1 == 1 {
			r = modMul(r, a, p)
		}
		a = modMul(a, a, p)
		b >>= 1
	}
	return r
}

func solve(n int, mod int, E [][]int) []int {
	F := make([]int, n+1)
	F[0] = 1
	for i := 1; i <= n; i++ {
		F[i] = modMul(i, F[i-1], mod)
	}

	I := make([]int, n+1)
	I[n] = pow(F[n], mod-2, mod)
	for i := n - 1; i >= 0; i-- {
		I[i] = modMul(i+1, I[i+1], mod)
	}

	nCr := func(n int, r int) int {
		if n < r {
			return 0
		}
		res := F[n]
		res = modMul(res, I[r], mod)
		res = modMul(res, I[n-r], mod)
		return res
	}

	g := NewGraph(n, 2*n)

	for _, e := range E {
		u, v := e[0], e[1]
		u--
		v--
		g.AddEdge(u, v)
		g.AddEdge(v, u)
	}

	dp := make([][]int, n+1)
	S := make([][]int, n+1)
	pre := make([][]int, n+1)
	suf := make([][]int, n+1)

	for i := 0; i <= n; i++ {
		dp[i] = make([]int, n+1)
		S[i] = make([]int, n+1)
		pre[i] = make([]int, n+2)
		suf[i] = make([]int, n+2)
		for j := 0; j <= n+1; j++ {
			pre[i][j] = 1
			suf[i][j] = 1
		}
	}

	var dfs func(p int, u int)

	dfs = func(p int, u int) {
		for i := g.node[u]; i > 0; i = g.next[i] {
			v := g.to[i]
			if p != v {
				dfs(u, v)
			}
		}

		var cnt int

		for i := g.node[u]; i > 0; i = g.next[i] {
			v := g.to[i]
			if p != v {
				cnt++
				for j := 0; j <= n; j++ {
					suf[cnt][j] = S[v][j]
					pre[cnt][j] = S[v][j]
				}
			}
		}

		if cnt == 0 {
			for j := 1; j <= n; j++ {
				dp[u][j] = 1
				S[u][j] = modAdd(S[u][j-1], dp[u][j], mod)
			}
			return
		}

		for j := 0; j <= n; j++ {
			for i := 1; i <= cnt; i++ {
				pre[i][j] = modMul(pre[i-1][j], pre[i][j], mod)
			}
			for i := cnt; i >= 1; i-- {
				suf[i][j] = modMul(suf[i][j], suf[i+1][j], mod)
			}
		}

		for i := 1; i <= n; i++ {
			dp[u][i] = pre[cnt][i]
		}

		if u != 0 {
			var c int
			for i := g.node[u]; i > 0; i = g.next[i] {
				v := g.to[i]
				if p == v {
					continue
				}
				c++
				var sum int
				for mx := 1; mx <= n; mx++ {
					dp[u][mx] = modAdd(dp[u][mx], modMul(sum, dp[v][mx], mod), mod)
					sum = modAdd(sum, modMul(pre[c-1][mx], suf[c+1][mx], mod), mod)
				}
			}
		}
		for i := 1; i <= cnt; i++ {
			for j := 0; j <= n; j++ {
				suf[i][j] = 1
				pre[i][j] = 1
			}
		}
		for i := 1; i <= n; i++ {
			S[u][i] = modAdd(dp[u][i], S[u][i-1], mod)
		}
	}

	dfs(0, 0)

	res := make([]int, n-1)

	for i := 1; i < n; i++ {
		var ans int
		for j := 1; j <= i; j++ {
			tmp := modMul(nCr(i, j), dp[0][j], mod)
			if (i-j)&1 == 1 {
				ans = modAdd(ans, mod-tmp, mod)
			} else {
				ans = modAdd(ans, tmp, mod)
			}
		}
		res[i-1] = ans
	}

	return res
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
	g.next = make([]int, e+1)
	g.to = make([]int, e+1)
	g.cur = 0
	return g
}

func (g *Graph) AddEdge(u, v int) {
	g.cur++
	g.next[g.cur] = g.node[u]
	g.node[u] = g.cur
	g.to[g.cur] = v
}

func (g *Graph) Reset() {
	g.cur = 0
}

package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	n, k := readTwoNums(reader)
	E := make([][]int, n-1)
	for i := 0; i < n-1; i++ {
		E[i] = readNNums(reader, 2)
	}

	res := solve(n, E, k)

	fmt.Println(res)
}

func readString(reader *bufio.Reader) string {
	s, _ := reader.ReadString('\n')
	for i := 0; i < len(s); i++ {
		if s[i] == '\n' {
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

const MOD = 1e9 + 7

func modAdd(a, b int) int {
	a += b
	if a >= MOD {
		a -= MOD
	}
	return a
}

func modMul(a, b int) int {
	return int(int64(a) * int64(b) % MOD)
}

func modSub(a, b int) int {
	return modAdd(a, MOD-b)
}

func pow(a, b int) int {
	r := 1
	for b > 0 {
		if b&1 == 1 {
			r = modMul(r, a)
		}
		a = modMul(a, a)
		b >>= 1
	}
	return r
}

const N = 2*1e5 + 10

var F [N]int
var IF [N]int

func init() {
	F[0] = 1
	for i := 1; i < N; i++ {
		F[i] = modMul(i, F[i-1])
	}

	IF[N-1] = pow(F[N-1], MOD-2)

	for i := int(N - 2); i >= 0; i-- {
		IF[i] = modMul(i+1, IF[i+1])
	}
}

func nCr(n int, r int) int {
	if n < r {
		return 0
	}
	return modMul(F[n], modMul(IF[r], IF[n-r]))
}

func solve(n int, E [][]int, k int) int {
	g := NewGraph(n, 2*n)

	for i := 0; i < len(E); i++ {
		u, v := E[i][0], E[i][1]
		u--
		v--
		g.AddEdge(u, v)
		g.AddEdge(v, u)
	}

	sz := make([]int, n)
	dp := make([]int, n)
	cnt := make([]int, n)
	var dfs func(p int, u int)

	var cur int

	dfs = func(p int, u int) {
		sz[u]++
		var sub int
		for i := g.nodes[u]; i > 0; i = g.next[i] {
			v := g.to[i]
			if p != v {
				dfs(u, v)
				sz[u] += sz[v]
				sub = modAdd(sub, nCr(sz[v], k))
			}
		}
		cnt[u] = modSub(nCr(sz[u], k), sub)
		dp[u] = modMul(cnt[u], sz[u])
		cur = modAdd(cur, dp[u])
	}

	dfs(0, 0)

	var ans int

	var dfs2 func(p int, v int)

	dfs2 = func(p int, v int) {
		ans = modAdd(ans, cur)
		for i := g.nodes[v]; i > 0; i = g.next[i] {
			u := g.to[i]
			if p != u {
				sv := sz[v]
				cv := cnt[v]
				dv := dp[v]
				su := sz[u]
				cu := cnt[u]
				du := dp[u]
				sz[v] -= sz[u]
				sz[u] = n
				scur := cur

				cnt[v] = modSub(cnt[v], nCr(sv, k))
				cnt[v] = modAdd(cnt[v], nCr(sz[v], k))
				cnt[v] = modAdd(cnt[v], nCr(su, k))

				cnt[u] = modSub(cnt[u], nCr(su, k))
				cnt[u] = modAdd(cnt[u], nCr(sz[u], k))
				cnt[u] = modSub(cnt[u], nCr(sz[v], k))

				dp[v] = modMul(cnt[v], sz[v])
				dp[u] = modMul(cnt[u], sz[u])

				cur = modSub(cur, modAdd(du, dv))
				cur = modAdd(cur, modAdd(dp[u], dp[v]))

				dfs2(v, u)

				dp[u] = du
				cnt[u] = cu
				sz[u] = su

				dp[v] = dv
				cnt[v] = cv
				sz[v] = sv

				cur = scur
			}
		}
	}

	dfs2(0, 0)

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

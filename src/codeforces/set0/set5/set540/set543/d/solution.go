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
	P := readNNums(reader, n-1)
	res := solve(n, P)
	var buf bytes.Buffer

	for i := 0; i < n; i++ {
		buf.WriteString(fmt.Sprintf("%d ", res[i]))
	}
	buf.WriteByte('\n')
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

const MOD = 1_000_000_007

func modMul(a, b int) int {
	return int(int64(a) * int64(b) % MOD)
}

func modDiv(a, b int) int {
	return modMul(a, inverse(b))
}

func modAdd(a, b int) int {
	a += b
	if a >= MOD {
		a -= MOD
	}
	return a
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

func inverse(a int) int {
	return pow(a, MOD-2)
}

func solve(n int, P []int) []int {
	// 当root确定为r时
	// dp[u] = 以u为根节点的子树中，r--u 都是好路时，它的方案数
	// 考虑它的children
	// 对于child v，如果 u-v bad, 那么v里面所有的路都需要被修好
	// dp[u] = II(dp[v] + 1)
	// 那么在从根节点从u移动到v时
	// dp[u] / (dp[v] + 1), 然后对于 dp[v] *= (dp[u] + 1)
	g := NewGraph(n, 2*n)

	for i := 1; i < n; i++ {
		u := P[i-1] - 1
		g.AddEdge(u, i)
		g.AddEdge(i, u)
	}

	dp := make([]int, n)

	var dfs func(p int, u int)
	cnt := make([]int, n)

	dfs = func(p int, u int) {
		dp[u] = 1
		for i := g.nodes[u]; i > 0; i = g.next[i] {
			v := g.to[i]
			if p != v {
				dfs(u, v)
				dp[u] = modMul(dp[u], modAdd(dp[v], 1))
				cnt[u]++
			}
		}
	}

	dfs(0, 0)

	ans := make([]int, n)

	var dfs2 func(p int, u int)

	dfs2 = func(p int, u int) {
		ans[u] = dp[u]
		suf := make([]int, cnt[u]+1)
		var it int
		for i := g.nodes[u]; i > 0; i = g.next[i] {
			v := g.to[i]
			if p == v {
				continue
			}
			suf[it] = modAdd(dp[v], 1)
			it++
		}
		suf[it] = 1
		for i := it - 1; i >= 0; i-- {
			suf[i] = modMul(suf[i], suf[i+1])
		}
		//dp[u] = 1
		it = 0
		var prev = 1
		if p >= 0 {
			prev = modAdd(dp[p], 1)
		}
		for i := g.nodes[u]; i > 0; i = g.next[i] {
			v := g.to[i]
			if p == v {
				continue
			}
			x := modAdd(dp[v], 1)
			dp[u] = modMul(prev, suf[it+1])
			dp[v] = modMul(dp[v], modAdd(dp[u], 1))
			dfs2(u, v)
			prev = modMul(prev, x)
			it++
		}
	}

	dfs2(-1, 0)

	return ans
}

func max(a, b int) int {
	if a >= b {
		return a
	}
	return b
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

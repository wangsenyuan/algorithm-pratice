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
		N, M, L := readThreeNums(reader)
		S, _ := reader.ReadString('\n')
		X, _ := reader.ReadString('\n')
		U := readNNums(reader, M)
		V := readNNums(reader, M)
		res := solve(N, L, M, S, X, U, V)
		buf.WriteString(fmt.Sprintf("%d\n", res))
	}
	fmt.Print(buf.String())
}

const MOD = 1000000007

func solve(N int, L, M int, S string, X string, U, V []int) int {
	cnt := make(map[[2]int]int)
	for i := 0; i < M; i++ {
		a, b := U[i], V[i]
		a--
		b--
		if a > b {
			a, b = b, a
		}
		cnt[[...]int{a, b}]++
	}

	g := NewGraph(N, 2*len(cnt))

	for k, v := range cnt {
		a, b := k[0], k[1]
		g.AddEdge(a, b, v)
		g.AddEdge(b, a, v)
	}

	dp := make([][]int64, N)
	vis := make([][]bool, N)
	for i := 0; i < N; i++ {
		dp[i] = make([]int64, L)
		vis[i] = make([]bool, L)
		for j := 0; j < L; j++ {
			dp[i][j] = 0
		}
	}

	que := make([]int, N*L)
	var front, end int
	for i := 0; i < N; i++ {
		if X[i] == S[0] {
			que[end] = i * L
			end++
			dp[i][0] = 1
			vis[i][0] = true
		}
	}

	var res int64
	for front < end {
		cur := que[front]
		front++
		u, l := cur/L, cur%L
		if l == L-1 {
			res += dp[u][l]
			res %= MOD
			continue
		}
		for i := g.nodes[u]; i > 0; i = g.next[i] {
			v := g.to[i]
			if X[v] == S[l+1] {
				dp[v][l+1] += (dp[u][l] * int64(g.val[i])) % MOD
				dp[v][l+1] %= MOD
				if !vis[v][l+1] {
					vis[v][l+1] = true
					que[end] = v*L + l + 1
					end++
				}
			}
		}
	}

	if minElement(L, S) == maxElement(L, S) {
		for i := 0; i < N; i++ {
			for j := i + 1; j < N; j++ {
				if X[i] == S[0] && X[j] == S[0] {
					res = (res + MOD - int64(pow(cnt[[2]int{i, j}], L-1))) % MOD
				}
			}
		}
	}

	return int(res)
}

func pow(a, b int) int {
	A := int64(a)
	R := int64(1)
	for b > 0 {
		if b&1 == 1 {
			R = (R * A) % MOD
		}
		A = (A * A) % MOD
		b >>= 1
	}
	return int(R)
}

func minElement(n int, s string) byte {
	res := s[0]
	for i := 1; i < n; i++ {
		if s[i] < res {
			res = s[i]
		}
	}
	return res
}

func maxElement(n int, s string) byte {
	res := s[0]
	for i := 1; i < n; i++ {
		if s[i] > res {
			res = s[i]
		}
	}
	return res
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
	next := make([]int, e+3)
	to := make([]int, e+3)
	val := make([]int, e+3)
	var cur int
	return &Graph{nodes, next, to, val, cur}
}

func (g *Graph) AddEdge(u, v, w int) {
	g.cur++
	g.next[g.cur] = g.nodes[u]
	g.nodes[u] = g.cur
	g.to[g.cur] = v
	g.val[g.cur] = w
}

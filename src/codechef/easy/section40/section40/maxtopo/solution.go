package main

import (
	"bufio"
	"bytes"
	"fmt"
	"math"
	"os"
	"sort"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	var buf bytes.Buffer
	tc := readNum(reader)

	for tc > 0 {
		tc--
		n, k := readTwoNums(reader)
		E := make([][]int, n-1)
		for i := 0; i < n-1; i++ {
			E[i] = readNNums(reader, 2)
		}

		res := solve(n, E, k)

		buf.WriteString(fmt.Sprintf("%d %d\n", res[0], res[1]))
	}

	fmt.Print(buf.String())
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

const MOD = 1e9 + 7
const MAX_N = 500010

var F []int
var IF []int

func init() {
	F = make([]int, MAX_N)

	F[0] = 1

	for i := 1; i < MAX_N; i++ {
		F[i] = modMul(i, F[i-1])
	}

	IF = make([]int, MAX_N)
	IF[MAX_N-1] = pow(F[MAX_N-1], MOD-2)

	for i := MAX_N - 2; i >= 0; i-- {
		IF[i] = modMul(i+1, IF[i+1])
	}
}

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

func nCr(n, r int) int {
	if n < r {
		return 0
	}
	res := modMul(F[n], IF[r])
	return modMul(res, IF[n-r])
}

const EPS = 1e-9

func solve(n int, E [][]int, k int) []int {
	g := NewGraph(n, n*2)

	for _, e := range E {
		u, v := e[0]-1, e[1]-1
		g.AddEdge(u, v)
		g.AddEdge(v, u)
	}
	sz := make([]int, n)
	var dfs1 func(p int, u int)

	dfs1 = func(p int, u int) {
		sz[u] = 1
		for i := g.nodes[u]; i > 0; i = g.next[i] {
			v := g.to[i]
			if v != p {
				dfs1(u, v)
				sz[u] += sz[v]
			}
		}
	}

	dfs1(0, 0)

	type Pair struct {
		first  float64
		second int
	}

	P := make([]Pair, n)

	var dfs2 func(p int, u int, val float64)

	dfs2 = func(p int, u int, val float64) {
		if p >= 0 {
			val -= math.Log(float64(sz[u]))
			val += math.Log(float64(n - sz[u]))
		}
		P[u] = Pair{val, u}

		for i := g.nodes[u]; i > 0; i = g.next[i] {
			v := g.to[i]
			if p != v {
				dfs2(u, v, val)
			}
		}
	}

	var val float64

	for i := 0; i < n; i++ {
		val += math.Log(float64(sz[i]))
	}

	dfs2(-1, 0, val)

	sort.Slice(P, func(i, j int) bool {
		return P[i].first < P[j].first || math.Abs(P[i].first-P[j].first) < EPS && P[i].second > P[j].second
	})

	root := P[k-1].second

	dfs1(-1, root)

	ans := 1
	for i := 1; i <= n; i++ {
		ans = modMul(ans, i)
		ans = modMul(ans, pow(sz[i-1], MOD-2))
	}
	return []int{root + 1, ans}
}

func solve1(n int, E [][]int, k int) []int {
	// C[u] = 以u为root时的topo排序数, 所有edge都远离u
	// 如何计算呢?
	// X[u, i] 为以u为root时，i子树的计数
	// X[i] = C(sz(i)-1, sz(j)) * C(sz(i) - 1 - sz(j), sz(k))... * II(X[j])
	// 考虑一种情况， 如果 (u, v) u是父节点，v是子节点，且u只有v一个子节点
	// X[u] = C(sz(u)-1, sz(v)) * X[v] == X[v]
	// 如果反过来以v为节点 Y[v] = Y[u], 但是 X[u] 和 Y[u]的关系是什么呢？
	// 如果 X[u] 有两个子节点
	// X[u] = C(sz[u] - 1, sz(v)) * C(sz[u] - 1 - sz[v], sz[w])) * X[v] * X[w]
	// 一个直观的判断是，某个节点的deg越多，它更应该处于更接近root的位置
	// 所以其实是一个从外部剥洋葱的过程吗？
	// 最大的计算出来了，第二大的呢？肯定是它的某个子节点，所有子节点单独不包括root计算出来，
	// 然后进行枚举即可
	g := NewGraph(n, n*2)

	for _, e := range E {
		u, v := e[0]-1, e[1]-1
		g.AddEdge(u, v)
		g.AddEdge(v, u)
	}
	sz := make([]int, n)
	var dfs1 func(p int, u int)

	dfs1 = func(p int, u int) {
		sz[u]++
		for i := g.nodes[u]; i > 0; i = g.next[i] {
			v := g.to[i]
			if v != p {
				dfs1(u, v)
				sz[u] += sz[v]
			}
		}
	}

	dfs1(0, 0)

	var dfs2 func(p int, u int) int

	dfs2 = func(p int, u int) int {

		for i := g.nodes[u]; i > 0; i = g.next[i] {
			v := g.to[i]
			if p == v {
				continue
			}
			if sz[v]*2 > n || sz[v]*2 == n && v > u {
				return dfs2(u, v)
			}
		}

		return u
	}

	root := dfs2(0, 0)

	C := make([]int, n)
	for i := 0; i < n; i++ {
		C[i] = 1
	}

	var dfs3 func(p int, u int)

	dfs3 = func(p int, u int) {
		sz[u] = 1
		for i := g.nodes[u]; i > 0; i = g.next[i] {
			v := g.to[i]
			if p != v {
				dfs3(u, v)
				sz[u] += sz[v]
			}
		}
		cnt := sz[u] - 1
		for i := g.nodes[u]; i > 0; i = g.next[i] {
			v := g.to[i]
			if p != v {
				C[u] = modMul(C[u], C[v])
				C[u] = modMul(C[u], nCr(cnt, sz[v]))
				cnt -= sz[v]
			}
		}
	}

	dfs3(-1, root)

	if k == 1 {
		return []int{root + 1, C[root]}
	}
	//root, C[root] got
	second := -1
	for i := g.nodes[root]; i > 0; i = g.next[i] {
		v := g.to[i]
		if second == -1 || sz[second] < sz[v] || sz[second] == sz[v] && second < v {
			second = v
		}
	}

	second_val := modMul(modMul(C[root], sz[second]), pow(sz[root]-sz[second], MOD-2))

	return []int{second + 1, second_val}
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

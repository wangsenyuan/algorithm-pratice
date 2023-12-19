package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	n, k := readTwoNums(reader)

	routes := make([][]int, n-1)
	for i := 0; i < n-1; i++ {
		routes[i] = readNNums(reader, 2)
	}
	res := solve(n, routes, k)

	fmt.Println(res)
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

const MOD = 1000000007

func modAdd(a, b int) int {
	a += b
	if a >= MOD {
		a -= MOD
	}
	return a
}

func modSub(a, b int) int {
	return modAdd(a, MOD-b)
}

func modMul(a int, b int) int {
	r := int64(a) * int64(b)
	return int(r % MOD)
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

func solve(n int, routes [][]int, k int) int {
	// if k == 1, just n
	if k%2 == 1 {
		// (1 + 1 ... + 1) / n
		return 1
	}

	// when k = 2,
	// 如果朋友们在两个节点u, v上，那么最优的结果就在它们的路径上
	// 这个是否可以用贡献法？它的贡献是 sz[u] * (n - sz[u]-1)
	// 不对，因为两个节点都可以在它的子树中
	// cnt[u] * (n - cnt[u] - 1) 不断迭代u的子树
	// 在已迭代过的子树中，和外面；这样可以保证，肯定经过节点u

	// when k = 3
	// 分2中情况，一种是三点共线，似乎似乎只能在中间那个点
	// 三点不共线，也只能在三个的公共lca点
	// 共有nCr(n, 3)中选择，每次选择只有一个点，但有n个选择
	// 1/n
	return solve2(n, routes)
}

func solve2(n int, routes [][]int) int {
	g := NewGraph(n, 2*n)

	for _, e := range routes {
		u, v := e[0]-1, e[1]-1
		g.AddEdge(u, v)
		g.AddEdge(v, u)
	}

	var res int
	sz := make([]int, n)
	var dfs func(p int, u int)

	dfs = func(p int, u int) {
		sz[u] = 1
		for i := g.nodes[u]; i > 0; i = g.next[i] {
			v := g.to[i]
			if p == v {
				continue
			}
			dfs(u, v)
			// 其中一个在u的已处理的子树中
			// 另外一个在子树v中
			sz[u] += sz[v]
		}
		res = modAdd(res, modMul(sz[u], n-sz[u]))
	}

	dfs(-1, 0)

	// nCr(n, 2) * n
	div := modMul(modMul(n, n-1), inverse(2))
	res = modMul(res, inverse(div))
	res = modAdd(res, 1)

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

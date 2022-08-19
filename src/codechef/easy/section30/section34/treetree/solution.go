package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	n := readNum(reader)

	T := make([][][]int, n)

	for i := 0; i < n; i++ {
		m := readNum(reader)
		T[i] = make([][]int, m-1)
		for j := 0; j < m-1; j++ {
			T[i][j] = readNNums(reader, 3)
		}
	}

	A := readNNums(reader, n-1)

	res := solve(n, T, A)

	fmt.Println(res)
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

func readString(reader *bufio.Reader) string {
	s, _ := reader.ReadString('\n')
	for i := 0; i < len(s); i++ {
		if s[i] == '\n' || s[i] == '\r' {
			return s[:i]
		}
	}
	return s
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

const MOD = 998244353

func modAdd(a, b int) int {
	a += b
	if a >= MOD {
		a -= MOD
	}
	return a
}

func modMul(a, b int) int {
	r := int64(a) * int64(b)
	return int(r % MOD)
}

func solve(n int, T [][][]int, A []int) int {
	// 每棵树的连接点是可以先计算出来的
	// 使低weight的边尽量的出现在高处
	// 假设艺节点r为root，对于某条边e，设它接入的点为u，则贡献为 weight(e) * size(u)
	// 貌似假设以节点1计算一个结果出来，
	// 然后假设W(p)计算出来了， u是p的子节点，W(u) = W(p) + weight(edge(p, u)) * (size(u) - (n - size(u)))
	var N int
	var res int
	nodes := make([]*Node, n)
	for i := 0; i < n; i++ {
		t, w := processTree(T[i])
		res = modAdd(res, t)
		nodes[i] = &Node{len(T[i]) + 1, w, 0}
		N += len(T[i]) + 1
	}

	for i := 0; i < n; i++ {
		nodes[i].cost = int64(nodes[i].cnt) * int64(N-nodes[i].cnt)
	}

	// now connect nodes (have value) by A, to minizie res
	sort.Ints(A)

	sort.Sort(Nodes(nodes))

	for i := 0; i < n; i++ {
		node := nodes[i]
		w := int(node.weight % MOD)
		res = modAdd(res, modMul(w, N-node.cnt))
	}

	// 组成一个star tree
	for i := 0; i < n-1; i++ {
		a := A[i]
		node := nodes[i+1]
		res = modAdd(res, modMul(a, modMul(node.cnt, N-node.cnt)))
	}

	return res
}

type Node struct {
	cnt    int
	weight int64
	cost   int64
}

type Nodes []*Node

func (nodes Nodes) Len() int {
	return len(nodes)
}

func (nodes Nodes) Less(i, j int) bool {
	return nodes[i].cost > nodes[j].cost
}

func (nodes Nodes) Swap(i, j int) {
	nodes[i], nodes[j] = nodes[j], nodes[i]
}

func processTree(T [][]int) (int, int64) {
	n := len(T) + 1
	g := NewGraph(n, (n-1)*2)

	for _, edge := range T {
		u, v, w := edge[0], edge[1], edge[2]
		u--
		v--
		g.AddEdge(u, v, w)
		g.AddEdge(v, u, w)
	}
	size := make([]int, n)

	var total int

	var dfs func(p, u int)

	weights := make([]int64, n)

	dfs = func(p, u int) {
		size[u]++
		for i := g.node[u]; i > 0; i = g.next[i] {
			v := g.to[i]
			w := g.weight[i]
			if p != v {
				dfs(u, v)
				size[u] += size[v]
				weights[0] += int64(w) * int64(size[v])
				total = modAdd(total, modMul(w, modMul(n-size[v], size[v])))
			}
		}
	}

	dfs(0, 0)

	var dfs2 func(p, u int)

	dfs2 = func(p, u int) {
		for i := g.node[u]; i > 0; i = g.next[i] {
			v := g.to[i]
			w := g.weight[i]
			if p != v {
				// when weights[u] know, what is weights[v]
				weights[v] = weights[u] - int64(w)*int64(size[v]) + int64(w)*int64(n-size[v])
				dfs2(u, v)
			}
		}
	}

	dfs2(0, 0)

	best := weights[0]

	for i := 0; i < n; i++ {
		if weights[i] < best {
			best = weights[i]
		}
	}

	return total, best
}

type Graph struct {
	node   []int
	next   []int
	to     []int
	weight []int
	cur    int
}

func NewGraph(n int, e int) *Graph {
	g := new(Graph)
	g.node = make([]int, n)
	g.next = make([]int, e+1)
	g.to = make([]int, e+1)
	g.weight = make([]int, e+1)
	return g
}

func (g *Graph) AddEdge(u, v, w int) {
	g.cur++
	g.next[g.cur] = g.node[u]
	g.node[u] = g.cur
	g.to[g.cur] = v
	g.weight[g.cur] = w
}

package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	res := process(reader)
	fmt.Println(res)
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

func process(reader *bufio.Reader) int {
	n := readNum(reader)
	ops := readNNums(reader, n)
	p := readNNums(reader, n-1)
	return solve(n, ops, p)
}

func solve(n int, ops []int, fa []int) int {
	g := NewGraph(n, n)

	for i := range n - 1 {
		g.AddEdge(fa[i]-1, i+1)
	}
	// sz[u] = u子树的叶子节点的数量
	sz := make([]int, n)
	dp := make([]int, n)

	type pair struct {
		first  int
		second int
	}

	max_op := func(u int, arr []pair) {
		// 这个时候只需要把最大的 dp[u] = dp[v] + sz[u] - sz[v]
		for _, cur := range arr {
			v := cur.second
			dp[u] = max(dp[u], dp[v]+sz[u]-sz[v])
		}
	}

	// check := func(arr []pair, k int, x int) bool {
	// 	// 1....k进行分配，能否得到最大值的最小值为x
	// 	// x + 1, x + 2, x + 3
	// 	for _, cur := range arr {
	// 		v := cur.second
	// 		// dp[v]
	// 		w := cur.first
	// 		if w >= x {
	// 			continue
	// 		}
	// 		// w < x, 需要cnt个数
	// 		cnt := sz[v] - w + 1
	// 		// 需要把x,x+1, x + cnt - 1这些数字分配给v
	// 		if x+cnt-1 > k {
	// 			return false
	// 		}
	// 		x += cnt
	// 	}

	// 	return true
	// }

	min_op := func(u int, arr []pair) {
		var sum int
		for _, cur := range arr {
			sum += cur.first - 1
		}
		dp[u] = sum + 1
	}

	var dfs func(u int)
	dfs = func(u int) {
		var arr []pair
		for i := g.nodes[u]; i > 0; i = g.next[i] {
			v := g.to[i]
			dfs(v)
			sz[u] += sz[v]
			arr = append(arr, pair{dp[v], v})
		}

		if len(arr) == 0 {
			dp[u] = 1
			sz[u] = 1
			return
		}
		if ops[u] == 1 {
			max_op(u, arr)
		} else {
			min_op(u, arr)
		}
	}

	dfs(0)

	return dp[0]
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

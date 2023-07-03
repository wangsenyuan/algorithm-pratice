package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	n := readNum(reader)
	P := readNNums(reader, n-1)
	A := readNNums(reader, n)
	res := solve(A, P)
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

// 5937041487
// 101653164688

func solve(A []int, P []int) int64 {

	// 在一颗子树中的人，不能往外跑
	// 所以 f(0) = max(f(children of 0))
	// 但是对于A[0]来说，部分人选择
	// 先不考虑怎么跑，先算出所有children的f(u)
	// 然后通过给最小的检查是否能够分配后，使得bandit最多只能抓大x个人
	n := len(A)
	g := NewGraph(n, n)

	for i := 2; i <= n; i++ {
		j := P[i-2] - 1
		g.AddEdge(j, i-1)
	}

	var best int64

	var dfs func(u int) (int64, int64)

	dfs = func(u int) (int64, int64) {
		var leaf int64
		var val = int64(A[u])
		for i := g.nodes[u]; i > 0; i = g.next[i] {
			v := g.to[i]
			a, b := dfs(v)
			leaf += a
			val += b
		}
		if leaf == 0 {
			leaf++
		}
		best = max(best, (val+leaf-1)/leaf)
		return leaf, val
	}

	dfs(0)

	return best
}

func max(a, b int64) int64 {
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

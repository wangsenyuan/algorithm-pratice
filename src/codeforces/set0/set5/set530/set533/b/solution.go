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
	nodes := make([][]int, n)
	for i := 0; i < n; i++ {
		nodes[i] = readNNums(reader, 2)
	}
	res := solve(n, nodes)

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

func solve(n int, vertexs [][]int) int {
	val := make([]int, n)
	g := NewGraph(n, n)
	var root int
	for i, cur := range vertexs {
		p, v := cur[0], cur[1]
		if p < 0 {
			root = i
		} else {
			g.AddEdge(p-1, i)
		}
		val[i] = v
	}

	dp := make([][]int, n)

	var dfs func(u int)
	dfs = func(u int) {
		dp[u] = make([]int, 2)
		var sum int
		var adds []int

		for i := g.nodes[u]; i > 0; i = g.next[i] {
			v := g.to[i]
			dfs(v)
			sum += dp[v][0]
			adds = append(adds, dp[v][1]-dp[v][0])
		}
		// 还有一种case没有处理，就是有一部分子节点是奇数个的，但是u不选
		sort.Ints(adds)
		reverse(adds)
		dp[u][0] = sum
		for i := 0; i+2 <= len(adds); i += 2 {
			tmp := adds[i] + adds[i+1]
			dp[u][0] += max(tmp, 0)
		}

		dp[u][1] = dp[u][0] + val[u]

		for i := 0; i < len(adds); i++ {
			sum += adds[i]
			if i%2 == 0 {
				dp[u][1] = max(dp[u][1], sum)
			}
		}
	}

	dfs(root)

	return max(dp[root][0], dp[root][1])
}

func reverse(arr []int) {
	for i, j := 0, len(arr)-1; i < j; i, j = i+1, j-1 {
		arr[i], arr[j] = arr[j], arr[i]
	}
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

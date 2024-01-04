package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	n, m := readTwoNums(reader)
	mat := make([][]int, m)
	for i := 0; i < m; i++ {
		mat[i] = readNNums(reader, n)
	}
	res := solve(mat)
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
func solve(mat [][]int) int {
	k := len(mat)
	// k <= 5
	n := len(mat[0])
	// n <= 1000
	// 对于 a, b,如果在k个序列中，都满足 a在b的前面，就在它们之间连一条线
	// 那么就是生成图中的直径
	// n * n * k
	cnt := make([][]int, n+1)
	for i := 0; i <= n; i++ {
		cnt[i] = make([]int, n+1)
	}

	for _, row := range mat {
		for i := 0; i < n; i++ {
			for j := i + 1; j < n; j++ {
				cnt[row[i]][row[j]]++
			}
		}
	}
	// 每个点
	g := NewGraph(n+1, n*n/2)
	for a := 1; a <= n; a++ {
		for b := 1; b <= n; b++ {
			if cnt[a][b] == k {
				g.AddEdge(a, b)
			}
		}
	}

	// 这个图肯定是dag，所以可以用dfs
	height := make([]int, n+1)

	var dfs func(u int)

	dfs = func(u int) {
		if height[u] > 0 {
			return
		}
		for i := g.nodes[u]; i > 0; i = g.next[i] {
			v := g.to[i]
			dfs(v)
			height[u] = max(height[u], height[v])
		}
		height[u]++
	}

	for i := 1; i <= n; i++ {
		if height[i] == 0 {
			dfs(i)
		}
	}
	var best int
	for i := 1; i <= n; i++ {
		best = max(best, height[i])
	}

	return best
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
	e++
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

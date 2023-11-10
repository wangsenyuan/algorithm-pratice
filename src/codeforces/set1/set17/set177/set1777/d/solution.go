package main

import (
	"bufio"
	"bytes"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	tc := readNum(reader)

	var buf bytes.Buffer

	for tc > 0 {
		tc--
		n := readNum(reader)
		E := make([][]int, n-1)
		for i := 0; i < n-1; i++ {
			E[i] = readNNums(reader, 2)
		}
		res := solve(n, E)
		buf.WriteString(fmt.Sprintf("%d\n", res))
	}
	fmt.Print(buf.String())
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

const mod = 1000000007

func add(arr ...int) int {
	var res int
	for _, a := range arr {
		res += a
		if res >= mod {
			res -= mod
		}
	}
	return res
}

func mul(a, b int) int {
	return a * b % mod
}

func pow(a, b int) int {
	r := 1
	for b > 0 {
		if b&1 == 1 {
			r = mul(r, a)
		}
		a = mul(a, a)
		b >>= 1
	}
	return r
}

func solve(n int, E [][]int) int {
	// 叶子节点的贡献，只有1
	// 在时刻0，是原始的配置
	// 在时刻1， 是它的子树中时刻0的1的个数
	// 在时刻2，是它的子树中时刻1的个数
	// ....
	// 这里主要的问题是，在时刻t，它的影响影子并不是它的所有子树中的节点
	// 假设节点x有两个叶子节点， 在时刻0，它的状态由自己决定， 在时刻1，就有叶子节点决定了
	//   它在4中状态下， 有两种状态为1，其实是叶子节点决定的，而不是x决定的
	// 但是如果考虑x的父节点y，在时刻1，由x, a, b决定，但是在时刻2，有a, b决定
	// 需要知道任何一个节点，在时刻1，2，3...时，有多少个节点决定它的值

	g := NewGraph(n, 2*n)

	for _, e := range E {
		u, v := e[0]-1, e[1]-1
		g.AddEdge(u, v)
		g.AddEdge(v, u)
	}

	dp := make([]int, n)

	var dfs func(p int, u int, depth int) int

	dfs = func(p int, u int, depth int) int {
		tmp := depth
		for i := g.nodes[u]; i > 0; i = g.next[i] {
			v := g.to[i]
			if p != v {
				tmp = max(tmp, dfs(u, v, depth+1))
			}
		}
		dp[u] = tmp - depth + 1
		return tmp
	}

	dfs(0, 0, 0)

	var ans int

	for i := 0; i < n; i++ {
		ans = add(ans, dp[i])
	}

	return mul(ans, pow(2, n-1))
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

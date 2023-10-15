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
		a := readNNums(reader, n)
		b := readNNums(reader, n)
		res := solve(a, b)
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

const mod = 998244353

func add(a, b int) int {
	a += b
	if a >= mod {
		a -= mod
	}
	return a
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

func solve(a []int, b []int) int {
	n := len(a)
	// 如果 a[i] = b[i], c[i]可以是任意值，迫使对方选择a[i]
	// 如果 a[i] != b[i], 为了迫使对方选择 a[i], 可以让c[i] = a[i], 或者c[i] = b[i]， 让对方选择b[i]
	// 应该要变成一个图
	// 考虑 (a, b) 如果有两个（a, b) 那么它们只能有两种选择
	// （a, b)  (a, c) (b, d)
	// (a, b) 的选择，会影响到 (a, c), (b, d)的选择
	// （a, b) (b, c) (c, d), (d, a),
	// 组成环的情况，只有2种选择
	// 如果一个数字，出现在了多个环上，就没有答案
	// 在环上的数字，不能有分叉出去
	// 所有的数字都必须在环上, 要么在一条路径上，要么在一个环上
	ok := make([]bool, n)

	g := NewGraph(n, 2*n)

	for i := 0; i < n; i++ {
		a[i]--
		b[i]--
		g.AddEdge(a[i], b[i])
		g.AddEdge(b[i], a[i])
	}

	var edge int
	var vert int
	var self bool

	var dfs func(u int)

	dfs = func(u int) {
		if ok[u] {
			return
		}
		ok[u] = true
		vert++
		for i := g.nodes[u]; i > 0; i = g.next[i] {
			v := g.to[i]
			edge++
			dfs(v)
			if u == v {
				self = true
			}
		}
	}

	ans := 1

	for i := 0; i < n; i++ {
		if ok[i] {
			continue
		}
		self = false
		edge = 0
		vert = 0
		dfs(i)

		if edge != 2*vert {
			return 0
		} else if self {
			ans = mul(ans, n)
		} else {
			ans = mul(ans, 2)
		}
	}

	return ans
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

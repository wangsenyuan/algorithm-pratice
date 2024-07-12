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
		n, k := readTwoNums(reader)
		a := readNNums(reader, n)
		edges := make([][]int, n-1)
		for i := 0; i < n-1; i++ {
			edges[i] = readNNums(reader, 2)
		}
		res := solve(k, a, edges)
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

const inf = 1 << 30
const X = 1001

var divs [X][]int
var ceilSqrt [X]int

func init() {
	for i := X - 1; i > 0; i-- {
		for j := i; j < X; j += i {
			divs[j] = append(divs[j], i)
		}
	}

	for i := 1; i < X; i++ {
		ceilSqrt[i] = 1
		x := i
		for p := 2; p*p <= x; p++ {
			for p2 := p * p; x%p2 == 0; x /= p2 {
				ceilSqrt[i] *= p
			}
			if x%p == 0 {
				ceilSqrt[i] *= p
				x /= p
			}
		}
		ceilSqrt[i] *= x
	}
}

func solve(k int, a []int, edges [][]int) int {
	n := len(a)

	g := NewGraph(n+1, 2*n)

	for _, e := range edges {
		g.AddEdge(e[0], e[1])
		g.AddEdge(e[1], e[0])
	}

	gd := make([]int, n+1)

	leaf := make([]bool, n+1)

	var dfs func(p int, u int)
	dfs = func(p int, u int) {
		gd[u] = a[u-1]
		leaf[u] = true
		for i := g.nodes[u]; i > 0; i = g.next[i] {
			v := g.to[i]
			if p != v {
				leaf[u] = false
				dfs(u, v)
				gd[u] = gcd(gd[u], gd[v])
			}
		}
	}

	dfs(0, 1)

	var f func(p int, u int, d int) int
	f = func(p int, u int, d int) int {
		if gd[u]%d == 0 {
			return 0
		}
		if gd[u]*gd[u]%d == 0 {
			return 1
		}

		if leaf[u] || a[u-1]*a[u-1]%d != 0 {
			return inf
		}
		// 至少要操作1次
		res := 1

		for i := g.nodes[u]; i > 0; i = g.next[i] {
			v := g.to[i]
			if p != v {
				res += f(u, v, ceilSqrt[d])
			}
		}
		return res
	}

	if k > 0 {
		for _, d := range divs[a[0]] {
			var res int
			for i := g.nodes[1]; i > 0; i = g.next[i] {
				res += f(1, g.to[i], d)
			}
			if res < k {
				return a[0] * d
			}
		}
	}

	return a[0]
}

func gcd(a, b int) int {
	for b > 0 {
		a, b = b, a%b
	}
	return a
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

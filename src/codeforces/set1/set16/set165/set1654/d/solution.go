package main

import (
	"bufio"
	"bytes"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	var buf bytes.Buffer

	tc := readNum(reader)

	for tc > 0 {
		tc--
		n := readNum(reader)
		E := make([][]int, n-1)
		for i := 0; i < n-1; i++ {
			E[i] = readNNums(reader, 4)
		}
		res := solve(n, E)
		buf.WriteString(fmt.Sprintf("%d\n", res))
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

const N = 200010

var factors [][]int
var F [N]int
var WF [N]int

func init() {
	lpf := make([]int, N)
	for i := 2; i < N; i++ {
		if lpf[i] == 0 {
			for j := i; j < N; j += i {
				lpf[j] = i
			}
		}
	}
	factors = make([][]int, N)
	for i := 2; i < N; i++ {
		x := i
		factors[i] = make([]int, 0, 1)
		for x > 1 {
			factors[i] = append(factors[i], lpf[x])
			x /= lpf[x]
		}
	}
}

const MOD = 998244353

func solve(n int, conditions [][]int) int {
	g := NewGraph(n, 2*n)

	unique_factors := make(map[int]bool)

	for k, cond := range conditions {
		i, j, x, y := cond[0], cond[1], cond[2], cond[3]
		i--
		j--
		g.AddEdge(i, j, k+1)
		g.AddEdge(j, i, -(k + 1))
		for _, f := range factors[x] {
			unique_factors[f] = true
		}
		for _, f := range factors[y] {
			unique_factors[f] = true
		}
	}

	var dfs1 func(p int, u int)

	dfs1 = func(p int, u int) {
		for i := g.nodes[u]; i > 0; i = g.next[i] {
			v := g.to[i]
			if p == v {
				continue
			}
			w := g.val[i]
			l := abs(w)
			x, y := conditions[l-1][2], conditions[l-1][3]
			if w < 0 {
				x, y = y, x
			}
			for _, z := range factors[y] {
				F[z]--
			}
			for _, z := range factors[x] {
				F[z]++
				WF[z] = max(WF[z], F[z])
			}
			dfs1(u, v)
			for _, z := range factors[y] {
				F[z]++
			}
			for _, z := range factors[x] {
				F[z]--
			}
		}
	}

	var dfs2 func(p int, u int, r int)

	var ans int

	dfs2 = func(p int, u int, r int) {
		ans = modAdd(ans, r)
		tmp := r
		for i := g.nodes[u]; i > 0; i = g.next[i] {
			v := g.to[i]
			if p != v {
				w := g.val[i]
				l := abs(w)
				x, y := conditions[l-1][2], conditions[l-1][3]
				if w < 0 {
					x, y = y, x
				}
				w = tmp
				for _, z := range factors[y] {
					w = modMul(w, z)
				}
				for _, z := range factors[x] {
					w = modMul(w, inverse(z))
				}
				dfs2(u, v, w)
			}
		}
	}

	dfs1(-1, 0)

	r := 1

	for x := range unique_factors {
		r = modMul(r, pow(x, WF[x]))
		F[x] = 0
		WF[x] = 0
	}

	dfs2(-1, 0, r)

	return ans
}

func max(a, b int) int {
	if a >= b {
		return a
	}
	return b
}

func abs(num int) int {
	if num < 0 {
		return -num
	}
	return num
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

func gcd(a, b int) int {
	for b > 0 {
		a, b = b, a%b
	}
	return a
}

func lcm(a, b int) int {
	g := gcd(a, b)
	r := modMul(a, g)
	return modMul(r, inverse(b))
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

type Graph struct {
	nodes []int
	next  []int
	to    []int
	val   []int
	cur   int
}

func NewGraph(n int, e int) *Graph {
	nodes := make([]int, n)
	next := make([]int, e+3)
	to := make([]int, e+3)
	val := make([]int, e+3)
	var cur int
	return &Graph{nodes, next, to, val, cur}
}

func (g *Graph) AddEdge(u, v, w int) {
	g.cur++
	g.next[g.cur] = g.nodes[u]
	g.nodes[u] = g.cur
	g.to[g.cur] = v
	g.val[g.cur] = w
}

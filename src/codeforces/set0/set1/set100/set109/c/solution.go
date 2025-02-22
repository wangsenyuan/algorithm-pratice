package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	n := readNum(reader)
	E := make([][]int, n-1)
	for i := 0; i < n-1; i++ {
		E[i] = readNNums(reader, 3)
	}
	res := solve(n, E)
	fmt.Println(res)
}

func readFloat64(bs []byte, pos int, r *float64) int {
	var first float64
	var second float64
	exp := 1.0
	var dot bool
	for pos < len(bs) && (bs[pos] == '.' || isDigit(bs[pos])) {
		if bs[pos] == '.' {
			dot = true
		} else if !dot {
			first = first*10 + float64(bs[pos]-'0')
		} else {
			second = second*10 + float64(bs[pos]-'0')
			exp *= 10
		}
		pos++
	}
	*r = first + second/exp
	//fmt.Fprintf(os.Stderr, "%.6f ", *r)
	return pos
}

func isDigit(b byte) bool {
	return b >= '0' && b <= '9'
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

func solve(n int, E [][]int) int64 {
	set := NewDSU(n)

	for _, e := range E {
		u, v, w := e[0]-1, e[1]-1, e[2]
		if checkLuck(w) {
			continue
		}
		set.Union(u, v)
	}

	var res int

	for u := 0; u < n; u++ {
		p := set.Find(u)
		if p == u {
			sz := set.cnt[p]
			res += sz * (n - sz) * (n - sz - 1)
		}
	}

	return int64(res)
}

type DSU struct {
	arr []int
	cnt []int
}

func NewDSU(n int) *DSU {
	arr := make([]int, n)
	cnt := make([]int, n)
	for i := 0; i < n; i++ {
		arr[i] = i
		cnt[i] = 1
	}
	return &DSU{arr, cnt}
}

func (this *DSU) Find(x int) int {
	if this.arr[x] != x {
		this.arr[x] = this.Find(this.arr[x])
	}
	return this.arr[x]
}

func (this *DSU) Union(x int, y int) bool {
	px := this.Find(x)
	py := this.Find(y)

	if px == py {
		return false
	}
	if this.cnt[px] < this.cnt[py] {
		px, py = py, px
	}
	this.cnt[px] += this.cnt[py]
	this.arr[py] = px
	return true
}

func solve1(n int, E [][]int) int64 {
	//如果一条边是lucky的，
	// u, v 那么 i在u一侧，(j, k)属于 v一侧的组合，就满足条件
	// 另外一种情况是 (j, i) & (i, k)

	g := NewGraph(n, 2*n)

	for _, e := range E {
		u, v := e[0], e[1]
		u--
		v--
		l := checkLuck(e[2])
		g.AddEdge(u, v, l)
		g.AddEdge(v, u, l)
	}

	dp := make([]int, n)

	var dfs func(p int, u int)
	sz := make([]int, n)
	dfs = func(p int, u int) {
		sz[u]++
		for i := g.nodes[u]; i > 0; i = g.next[i] {
			v := g.to[i]
			if v != p {
				dfs(u, v)
				sz[u] += sz[v]
				if g.lucky[i] {
					dp[u] += sz[v]
				} else {
					dp[u] += dp[v]
				}
			}
		}
	}

	count := func(n int) int64 {
		return int64(n) * int64(n-1)
	}

	dfs(-1, 0)

	var ans int64

	var dfs2 func(p int, u int, cnt int)

	dfs2 = func(p int, u int, cnt int) {
		// 如果以u为一个端点，(j, k)都在u的子树中
		ans += count(dp[u])
		// 如果以u为断点，(j, k) 在子树外边
		ans += count(cnt)
		// 如果以u为中心点，(j, k) 分别在子树中
		ans += 2 * int64(cnt) * int64(dp[u])

		sum := cnt

		for i := g.nodes[u]; i > 0; i = g.next[i] {
			v := g.to[i]
			if p != v {
				if g.lucky[i] {
					sum += sz[v]
				} else {
					sum += dp[v]
				}
			}
		}

		for i := g.nodes[u]; i > 0; i = g.next[i] {
			v := g.to[i]
			if p != v {
				if g.lucky[i] {
					dfs2(u, v, n-sz[v])
				} else {
					dfs2(u, v, sum-dp[v])
				}
			}
		}
	}

	dfs2(-1, 0, 0)

	return ans
}

func checkLuck(num int) bool {
	for num > 0 {
		r := num % 10
		if r != 4 && r != 7 {
			return false
		}
		num /= 10
	}
	return true
}

type Graph struct {
	nodes []int
	next  []int
	to    []int
	lucky []bool
	cur   int
}

func NewGraph(n int, e int) *Graph {
	nodes := make([]int, n)
	next := make([]int, e+3)
	to := make([]int, e+3)
	lucky := make([]bool, e+3)
	var cur int
	return &Graph{nodes, next, to, lucky, cur}
}

func (g *Graph) AddEdge(u, v int, l bool) {
	g.cur++
	g.next[g.cur] = g.nodes[u]
	g.nodes[u] = g.cur
	g.to[g.cur] = v
	g.lucky[g.cur] = l
}

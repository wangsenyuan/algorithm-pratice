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
		A := readNNums(reader, n)
		edges := make([][]int, n-1)
		for i := 0; i < n-1; i++ {
			edges[i] = readNNums(reader, 2)
		}
		res := solve(n, A, edges)
		for i := 0; i < n; i++ {
			buf.WriteString(fmt.Sprintf("%d ", res[i]))
		}
		buf.WriteByte('\n')
	}
	fmt.Print(buf.String())
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

type Pair struct {
	first  int
	second int
}

const MAX_X = 1000010

const MOD = 1000000007

var divisors [MAX_X][]Pair
var ansd [MAX_X]int

func modAdd(a, b int) int {
	a += b
	if a >= MOD {
		a -= MOD
	}
	return a
}

func modMul(a, b int) int {
	res := int64(a) * int64(b)
	return int(res % MOD)
}

func pow(a, b int) int {
	res := 1
	for b > 0 {
		if b&1 == 1 {
			res = modMul(res, a)
		}
		a = modMul(a, a)
		b >>= 1
	}
	return res
}

func inverse(a int) int {
	return pow(a, MOD-2)
}

func modDiv(a, b int) int {
	return modMul(a, inverse(b))
}

func init() {
	for i := 0; i < MAX_X; i++ {
		divisors[i] = make([]Pair, 0, 1)
		ansd[i] = 1
	}
	set := make([]bool, MAX_X)

	for i := 2; i < MAX_X; i++ {
		if !set[i] {
			for j := i; j < MAX_X; j += i {
				set[j] = true
				tmp := j
				var cnt int
				for tmp%i == 0 {
					cnt++
					tmp /= i
				}
				ansd[j] = modMul(ansd[j], cnt+1)

				divisors[j] = append(divisors[j], Pair{i, cnt})
			}
		}
	}
}

func solve(n int, A []int, edges [][]int) []int {
	g := NewGraph(n, 2*(n-1))

	for _, edge := range edges {
		u, v := edge[0], edge[1]
		u--
		v--
		g.AddEdge(u, v)
		g.AddEdge(v, u)
	}

	sz := make([]int, n)

	in := make([]int, n)
	out := make([]int, n)
	ver := make([]int, n)

	var get_size func(p int, u int)
	var time int
	get_size = func(p int, u int) {
		ver[time] = u
		in[u] = time
		time++

		sz[u]++

		for i := g.node[u]; i > 0; i = g.next[i] {
			v := g.to[i]
			if p != v {
				get_size(u, v)
				sz[u] += sz[v]
			}
		}

		out[u] = time
	}

	get_size(-1, 0)

	var dfs func(p int, u int, keep bool)

	currdiv := 1

	currfac := make(map[int]int)

	ans := make([]int, n)

	dfs = func(p int, u int, keep bool) {
		var mx int
		big := -1
		for i := g.node[u]; i > 0; i = g.next[i] {
			v := g.to[i]
			if p != v && sz[v] > mx {
				mx = sz[v]
				big = v
			}
		}
		for i := g.node[u]; i > 0; i = g.next[i] {
			v := g.to[i]
			if v != p && v != big {
				dfs(u, v, false)
			}
		}

		if big >= 0 {
			dfs(u, big, true)
		}

		for i := g.node[u]; i > 0; i = g.next[i] {
			v := g.to[i]
			if v != p && v != big {
				for k := in[v]; k < out[v]; k++ {
					for _, j := range divisors[A[ver[k]]] {
						currdiv = modDiv(currdiv, currfac[j.first]+1)
						currfac[j.first] += j.second
						currdiv = modMul(currdiv, currfac[j.first]+1)
					}
				}
			}
		}

		for _, j := range divisors[A[u]] {
			currdiv = modDiv(currdiv, currfac[j.first]+1)
			currfac[j.first] += j.second
			currdiv = modMul(currdiv, currfac[j.first]+1)
		}
		ans[u] = currdiv
		if !keep {
			currdiv = 1
			currfac = make(map[int]int)
		}
	}

	dfs(0, 0, false)

	return ans
}

type Graph struct {
	node []int
	next []int
	to   []int
	cur  int
}

func NewGraph(n int, e int) *Graph {
	g := new(Graph)
	g.node = make([]int, n)
	g.next = make([]int, e+1)
	g.to = make([]int, e+1)
	g.cur = 0
	return g
}

func (g *Graph) AddEdge(u, v int) {
	g.cur++
	g.next[g.cur] = g.node[u]
	g.node[u] = g.cur
	g.to[g.cur] = v
}

package main

import (
	"bufio"
	"bytes"
	"fmt"
	"os"
	"sort"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	tc := readNum(reader)

	var buf bytes.Buffer
	for tc > 0 {
		tc--
		n := readNum(reader)
		A := readNNums(reader, n)
		B := readNNums(reader, n)
		res := solve(A, B)

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

func readNInt64s(reader *bufio.Reader, n int) []int64 {
	res := make([]int64, n)
	s, _ := reader.ReadBytes('\n')

	var pos int

	for i := 0; i < n; i++ {
		pos = readInt64(s, pos, &res[i]) + 1
	}

	return res
}

func readInt64(bytes []byte, from int, val *int64) int {
	i := from
	var sign int64 = 1
	if bytes[i] == '-' {
		sign = -1
		i++
	}
	var tmp int64
	for i < len(bytes) && bytes[i] >= '0' && bytes[i] <= '9' {
		tmp = tmp*10 + int64(bytes[i]-'0')
		i++
	}
	*val = tmp * sign
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

	for len(bs) == 0 || bs[0] == '\n' || bs[0] == '\r' {
		bs, _ = reader.ReadBytes('\n')
	}

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

const MAX_NUM = 31660
const MAX_N = 400
const MAX_NODE = MAX_N*MAX_N + MAX_N*MAX_N*10 + 2
const MAX_EDGE = MAX_NODE*2 + 10
const INF = 1000000000

var primes []int

func init() {
	set := make([]bool, MAX_NUM)

	for i := 2; i < MAX_NUM; i++ {
		if !set[i] {
			primes = append(primes, i)
			for j := i * i; j < MAX_NUM; j += i {
				set[j] = true
			}
		}
	}
}

func solve(A []int, B []int) int {
	var left []int
	var right []int

	for i := 0; i < len(A); i++ {
		for j := 0; j < len(B); j++ {
			g := gcd(A[i], B[j])
			if g == 1 {
				continue
			}
			if A[i] < B[j] {
				left = append(left, g)
			} else if A[i] > B[j] {
				right = append(right, g)
			}
		}
	}

	sort.Ints(left)
	sort.Ints(right)

	backupLeft := copyArray(left)
	backupRight := copyArray(right)
	left = unique(left)
	right = unique(right)

	var leftFactor [][]int
	var rightFactor [][]int
	var mid []int

	for i := 0; i < len(left); i++ {
		tmp := factorize(left[i])
		leftFactor = append(leftFactor, tmp)
		mid = append(mid, tmp...)
	}

	for i := 0; i < len(right); i++ {
		tmp := factorize(right[i])
		rightFactor = append(rightFactor, tmp)
		mid = append(mid, tmp...)
	}

	sort.Ints(mid)
	mid = unique(mid)

	g := NewGraph(MAX_NODE, MAX_EDGE)

	var source = len(left) + len(right) + len(mid)
	sink := source + 1

	addEdge := func(u, v, w int) {
		g.AddEdge(u, v, w)
		g.AddEdge(v, u, 0)
	}

	for i := 0; i < len(left); i++ {
		cnt := upperBound(backupLeft, left[i]) - lowerBound(backupLeft, left[i])
		addEdge(source, i, cnt)
		for j := 0; j < len(leftFactor[i]); j++ {
			k := lowerBound(mid, leftFactor[i][j])
			addEdge(i, len(left)+k, INF)
		}
	}

	for i := 0; i < len(right); i++ {
		cnt := upperBound(backupRight, right[i]) - lowerBound(backupRight, right[i])
		addEdge(len(left)+len(mid)+i, sink, cnt)
		for j := 0; j < len(rightFactor[i]); j++ {
			k := lowerBound(mid, rightFactor[i][j])
			addEdge(len(left)+k, len(left)+len(mid)+i, INF)
		}
	}

	return dinic(source, sink, sink+1, g)
}

func dinic(src, snk int, n int, g *Graph) int {
	level := make([]int, n)
	all_minus_one := make([]int, n)
	for i := 0; i < n; i++ {
		all_minus_one[i] = -1
	}

	que := make([]int, n)

	bfs := func() bool {
		var front, end int
		copy(level, all_minus_one)
		level[src] = 0
		que[end] = src
		end++
		for front < end {
			u := que[front]
			front++
			for i := g.node[u]; i > 0; i = g.next[i] {
				if g.limit[i] > g.flow[i] && level[g.to[i]] == -1 {
					v := g.to[i]
					level[v] = level[u] + 1
					que[end] = v
					end++
				}
			}
		}
		return level[snk] > 0
	}

	pos := make([]int, n)

	var dfs func(u int, flow int) int
	dfs = func(u int, flow int) int {
		if flow == 0 {
			return 0
		}
		if u == snk {
			return flow
		}

		for pos[u] > 0 {
			i := pos[u]
			v := g.to[i]
			if level[v] == level[u]+1 && g.flow[i] < g.limit[i] {
				tr := dfs(v, min(flow, g.limit[i]-g.flow[i]))
				if tr > 0 {
					g.flow[i] += tr
					g.flow[i^1] -= tr
					return tr
				}
			}

			pos[u] = g.next[i]
		}
		return 0
	}
	var flow int
	for bfs() {
		for i := 0; i < n; i++ {
			pos[i] = g.node[i]
		}
		for {
			cur := dfs(src, INF)
			if cur == 0 {
				break
			}
			flow += cur
		}
	}
	return flow
}

func min(a, b int) int {
	if a <= b {
		return a
	}
	return b
}

func upperBound(arr []int, num int) int {
	return search(len(arr), func(i int) bool {
		return arr[i] > num
	})
}

func lowerBound(arr []int, num int) int {
	return search(len(arr), func(i int) bool {
		return arr[i] >= num
	})
}

func search(n int, fn func(int) bool) int {
	l, r := 0, n
	for l < r {
		mid := (l + r) / 2
		if fn(mid) {
			r = mid
		} else {
			l = mid + 1
		}
	}
	return r
}

type Graph struct {
	node  []int
	next  []int
	to    []int
	flow  []int
	limit []int
	cur   int
}

func NewGraph(n int, e int) *Graph {
	g := new(Graph)
	g.node = make([]int, n)
	g.next = make([]int, e+3)
	g.to = make([]int, e+3)
	g.flow = make([]int, e+3)
	g.limit = make([]int, e+3)
	g.cur = 1
	return g
}

func (g *Graph) AddEdge(u, v, w int) {
	g.cur++
	g.next[g.cur] = g.node[u]
	g.node[u] = g.cur
	g.to[g.cur] = v
	g.limit[g.cur] = w
	g.flow[g.cur] = 0
}

func factorize(num int) []int {
	var res []int

	for j := 0; j < len(primes) && primes[j] <= num/primes[j]; j++ {
		if num%primes[j] == 0 {
			res = append(res, primes[j])
			for num%primes[j] == 0 {
				num /= primes[j]
			}
		}
	}

	if num > 1 {
		res = append(res, num)
	}
	return res
}

func unique(arr []int) []int {
	var n int
	for i := 1; i <= len(arr); i++ {
		if i == len(arr) || arr[i] > arr[i-1] {
			arr[n] = arr[i-1]
			n++
		}
	}
	return arr[:n]
}

func copyArray(arr []int) []int {
	res := make([]int, len(arr))
	copy(res, arr)
	return res
}

func gcd(a, b int) int {
	for b > 0 {
		a, b = b, a%b
	}
	return a
}

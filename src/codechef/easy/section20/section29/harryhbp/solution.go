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

	E := make([][]int, n-1)

	for i := 0; i < n-1; i++ {
		E[i] = readNNums(reader, 3)
	}

	res := solve(n, E)

	fmt.Println(res)
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
		if s[i] == '\n' {
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
	if n == 0 {
		return res
	}
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

func readFloat64(bytes []byte, from int, val *float64) int {
	i := from
	var sign float64 = 1
	if bytes[i] == '-' {
		sign = -1
		i++
	}
	var real int64

	for i < len(bytes) && bytes[i] >= '0' && bytes[i] <= '9' {
		real = real*10 + int64(bytes[i]-'0')
		i++
	}

	if i == len(bytes) || bytes[i] != '.' {
		*val = float64(real)
		return i
	}

	// bytes[i] == '.'
	i++

	var fraq float64
	var base float64 = 0.1
	for i < len(bytes) && bytes[i] >= '0' && bytes[i] <= '9' {
		fraq += base * float64(bytes[i]-'0')
		base /= 10
		i++
	}

	*val = (float64(real) + fraq) * sign

	return i
}

func readNFloats(reader *bufio.Reader, n int) []float64 {
	s, _ := reader.ReadBytes('\n')
	res := make([]float64, n)
	var pos int
	for i := 0; i < n; i++ {
		pos = readFloat64(s, pos, &res[i]) + 1
	}
	return res
}

const MOD = 1000000007

const MAX_N = 100010

type Bit []int64

func NewBit(n int) Bit {
	arr := make([]int64, n)
	return Bit(arr)
}

func (this Bit) Add(p int, v int64) {
	v %= MOD
	if v < 0 {
		v += MOD
	}
	for p > 0 {
		modAdd(&this[p], v)
		p -= p & -p
	}
}

func (this Bit) Get(p int) int64 {
	var res int64

	for p < len(this) {
		modAdd(&res, this[p])
		p += p & -p
	}

	return res
}

func solve(N int, E [][]int) int {

	g := NewGraph(N, len(E)*2)

	for _, e := range E {
		u, v, w := e[0], e[1], e[2]
		u--
		v--
		g.AddEdge(u, v, w)
		g.AddEdge(v, u, w)
	}

	cut := make([]bool, N)

	size := make([]int, N)

	var dfs_size func(p, u int)

	dfs_size = func(p, u int) {
		size[u] = 1

		for i := g.nodes[u]; i > 0; i = g.next[i] {
			v := g.to[i]
			if p != v && !cut[v] {
				dfs_size(u, v)
				size[u] += size[v]
			}
		}
	}

	var find_centroid func(p, u, tot int) int

	find_centroid = func(p, u, tot int) int {
		ret := u

		for i := g.nodes[u]; i > 0; i = g.next[i] {
			v := g.to[i]
			if p != v && !cut[v] && size[v]*2 >= tot {
				ret = find_centroid(u, v, tot)
			}
		}

		return ret
	}

	all := make([]Pair, N)

	var get_weight_pair func(p, u int, mi int, ma int, pos *int)

	get_weight_pair = func(p, u int, mi int, ma int, pos *int) {
		all[*pos] = Pair{mi, ma}
		*pos++

		for i := g.nodes[u]; i > 0; i = g.next[i] {
			v := g.to[i]
			w := g.cost[i]
			if p != v && !cut[v] {
				get_weight_pair(u, v, min(mi, w), max(ma, w), pos)
			}
		}
	}

	var ans int64

	var dfs func(u int)

	sum := NewBit(MAX_N)
	cnt := NewBit(MAX_N)

	dfs = func(u int) {
		dfs_size(u, u)

		tot := size[u]

		u = find_centroid(u, u, tot)

		var front, end int

		for i := g.nodes[u]; i > 0; i = g.next[i] {
			v := g.to[i]
			if cut[v] {
				continue
			}
			get_weight_pair(u, v, g.cost[i], g.cost[i], &end)

			modSub(&ans, calc(all[front:end], sum, cnt))

			front = end
		}

		modAdd(&ans, calc(all[:end], sum, cnt))

		for i := 0; i < end; i++ {
			a := all[i]
			modAdd(&ans, int64(a.first)*int64(a.second)%MOD)
		}

		cut[u] = true

		for i := g.nodes[u]; i > 0; i = g.next[i] {
			v := g.to[i]
			if !cut[v] {
				dfs(v)
			}
		}
	}

	dfs(0)

	return int(ans)
}

type Graph struct {
	nodes []int
	next  []int
	to    []int
	cost  []int
	cur   int
}

func NewGraph(n int, e int) *Graph {
	e++
	nodes := make([]int, n)
	next := make([]int, e)
	to := make([]int, e)
	cost := make([]int, e)
	return &Graph{nodes, next, to, cost, 0}
}

func (g *Graph) AddEdge(u, v, w int) {
	g.cur++
	g.next[g.cur] = g.nodes[u]
	g.nodes[u] = g.cur
	g.to[g.cur] = v
	g.cost[g.cur] = w
}

type Pair struct {
	first  int
	second int
}

type Pairs []Pair

func (this Pairs) Len() int {
	return len(this)
}

func (this Pairs) Less(i, j int) bool {
	return this[i].first < this[j].first
}

func (this Pairs) Swap(i, j int) {
	this[i], this[j] = this[j], this[i]
}

func min(a, b int) int {
	if a <= b {
		return a
	}
	return b
}

func max(a, b int) int {
	if a >= b {
		return a
	}
	return b
}

func modAdd(a *int64, b int64) {
	*a += b
	if *a >= MOD {
		*a -= MOD
	}
}

func modSub(a *int64, b int64) {
	modAdd(a, MOD-b)
}

/**
// Given a vector of pair find the summation of min(A[i].first, A[j].first) * max(A[i].second, A[j].second) for all pair i,j in the index of array.
*/
func calc(arr []Pair, sum Bit, cnt Bit) int64 {
	sort.Sort(Pairs(arr))

	n := len(arr)
	var ret int64

	for i := n - 1; i >= 0; i-- {
		// cur[i].first is min
		// cur[i].first 作为最小值, cur[i].second 不作为最大值时
		tmp := int64(arr[i].first) * sum.Get(arr[i].second) % MOD
		modAdd(&ret, tmp)

		// cur[i].first 作为最小值， cur[i].second 作为最大值时
		tmp = (int64(n-i-1) - cnt.Get(arr[i].second)) * int64(arr[i].second) % MOD
		tmp = int64(arr[i].first) * tmp % MOD

		modAdd(&ret, tmp)

		sum.Add(arr[i].second, int64(arr[i].second))
		cnt.Add(arr[i].second, 1)
	}

	for i := 0; i < n; i++ {
		// cancel
		sum.Add(arr[i].second, -int64(arr[i].second))
		cnt.Add(arr[i].second, -1)
	}

	return ret
}

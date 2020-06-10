package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	tc := readNum(reader)

	for tc > 0 {
		tc--
		n := readNum(reader)
		E := make([][]int, n-1)
		for i := 0; i < n-1; i++ {
			E[i] = readNNums(reader, 2)
		}
		A := readNNums(reader, n)
		q := readNum(reader)
		queries := make([][]int, q)
		for i := 0; i < q; i++ {
			queries[i] = readNNums(reader, 2)
		}
		res := solve(n, A, E, queries)

		for i := 0; i < len(res); i++ {
			fmt.Println(res[i])
		}
	}

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
		for x < len(bs) && (bs[x] < '0' || bs[x] > '9') {
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

const MX = 1000010
const MN = 100007
const MOD = 1000000007
const BLOCK_SIZE = 700

var spf []int
var inv []int64

func init() {
	spf = make([]int, MX)

	for x := 2; x < MX; x++ {
		if spf[x] == 0 {
			for y := x; y < MX; y += x {
				if spf[y] == 0 {
					spf[y] = x
				}
			}
		}
	}
	inv = make([]int64, 20*MN)

	for i := 1; i < len(inv); i++ {
		inv[i] = pow(i, MOD-2)
	}
}

func pow(a, b int) int64 {
	A := int64(a)
	R := int64(1)

	for b > 0 {
		if b&1 == 1 {
			R = (R * A) % MOD
		}
		A = (A * A) % MOD
		b >>= 1
	}
	return R
}

func solve(n int, A []int, E [][]int, queries [][]int) []int64 {
	degree := make([]int, n)
	for i := 0; i < len(E); i++ {
		u, v := E[i][0], E[i][1]
		u--
		v--
		degree[u]++
		degree[v]++
	}

	g := make([][]int, n)
	for i := 0; i < n; i++ {
		g[i] = make([]int, 0, degree[i])
	}
	for i := 0; i < n-1; i++ {
		u, v := E[i][0], E[i][1]
		u--
		v--
		g[u] = append(g[u], v)
		g[v] = append(g[v], u)
	}

	eu := make([]int, 2*n)
	st := make([]int, n)
	en := make([]int, n)

	var dfs func(p, u int, time *int)
	dfs = func(p, u int, time *int) {
		*time++
		eu[*time] = u
		st[u] = *time
		for _, v := range g[u] {
			if p == v {
				continue
			}
			dfs(u, v, time)
		}
		*time++
		eu[*time] = u
		en[u] = *time
	}
	time := new(int)
	*time = -1
	dfs(-1, 0, time)

	lca := NewLCA(g)

	qs := make([]Question, len(queries))

	for i := 0; i < len(queries); i++ {
		u, v := queries[i][0], queries[i][1]
		u--
		v--
		if st[u] > st[v] {
			u, v = v, u
		}
		l := lca.findLca(u, v)
		if l == u {
			qs[i] = Question{i, st[u], st[v], -1}
		} else {
			qs[i] = Question{i, en[u], st[v], l}
		}
	}

	sort.Sort(Questions(qs))

	f := make([]int64, MX)

	for i := 0; i < MX; i++ {
		f[i] = 1
	}

	ans := make([]int64, len(queries))

	var left, right = 0, -1
	var cur int64 = 1
	added := make([]bool, n)

	update := func(u int) {
		x := A[u]
		for x > 1 {
			p := spf[x]
			var cnt int64
			for x%p == 0 {
				x /= p
				cnt++
			}
			cur = (cur * inv[f[p]]) % MOD
			delta := int64(1)
			if added[u] {
				delta = -1
			}
			f[p] += delta * cnt
			cur = (cur * f[p]) % MOD
		}
		added[u] = !added[u]
	}

	for _, q := range qs {
		for right < q.second {
			right++
			update(eu[right])
		}

		for left > q.first {
			left--
			update(eu[left])
		}

		for right > q.second {
			update(eu[right])
			right--
		}

		for left < q.first {
			update(eu[left])
			left++
		}

		var fact int64 = 1

		if q.lca != -1 {
			u := q.lca
			x := A[u]
			for x > 1 {
				p := spf[x]
				var cnt int64
				for x%p == 0 {
					x /= p
					cnt++
				}
				fact = (fact * inv[f[p]]) % MOD
				fact = (fact * (f[p] + cnt)) % MOD
			}
		}
		ans[q.index] = (cur * fact) % MOD
	}

	return ans
}

type Question struct {
	index  int
	first  int
	second int
	lca    int
}

type Questions []Question

func (qs Questions) Len() int {
	return len(qs)
}

func (qs Questions) Less(i, j int) bool {
	a, b := qs[i], qs[j]
	if a.first/BLOCK_SIZE < b.first/BLOCK_SIZE {
		return true
	}
	if a.first/BLOCK_SIZE == b.first/BLOCK_SIZE {
		return a.second < b.second
	}
	return false
}

func (qs Questions) Swap(i, j int) {
	qs[i], qs[j] = qs[j], qs[i]
}

type LCA struct {
	n int
	d []int
	h int
	P [][]int
}

func NewLCA(g [][]int) LCA {
	n := len(g)
	d := make([]int, n)
	h := 20
	P := make([][]int, n)
	for i := 0; i < n; i++ {
		P[i] = make([]int, h)
	}

	var dfs func(p, u int)

	dfs = func(p, u int) {
		P[u][0] = p

		for _, v := range g[u] {
			if p == v {
				continue
			}
			d[v] = d[u] + 1
			dfs(u, v)
		}
	}

	dfs(-1, 0)

	for j := 1; j < h; j++ {
		for i := 0; i < n; i++ {
			p := P[i][j-1]
			if p >= 0 {
				P[i][j] = P[p][j-1]
			}
		}
	}

	return LCA{n, d, h, P}
}

func (lca LCA) findLca(u, v int) int {
	d := lca.d
	if d[u] < d[v] {
		u, v = v, u
	}
	P := lca.P
	h := lca.h
	for i := h - 1; i >= 0; i-- {
		if d[u]-(1<<uint(i)) >= d[v] {
			u = P[u][i]
		}
	}

	if u == v {
		return u
	}

	for i := h - 1; i >= 0; i-- {
		if P[u][i] != P[v][i] {
			u = P[u][i]
			v = P[v][i]
		}
	}

	return P[u][0]
}

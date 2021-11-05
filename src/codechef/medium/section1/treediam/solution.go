package main

import (
	"bufio"
	"bytes"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	n := readNum(reader)
	W := readNNums(reader, n)
	E := make([][]int, n-1)
	for i := 0; i < n-1; i++ {
		E[i] = readNNums(reader, 2)
	}
	S := make([]int, n-1)
	for i := 0; i < n-1; i++ {
		S[i] = readNum(reader)
	}
	res := solve(n, W, E, S)

	var buf bytes.Buffer
	for _, num := range res {
		buf.WriteString(fmt.Sprintf("%d\n", num))
	}

	fmt.Print(buf.String())
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

const MOD = 1000000007

func modAdd(a *int, b int) {
	*a += b
	if *a >= MOD {
		*a -= MOD
	}
}

func modMul(a *int64, b int64) {
	A := int64(*a) * int64(b)
	A %= MOD
	*a = A
}

func pow(a, b int64) int64 {
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

func inverse(a int64) int64 {
	return pow(a, MOD-2)
}

const H = 20

func solve(n int, W []int, E [][]int, S []int) []int {

	g := NewGraph(n, len(E)*2)

	for _, e := range E {
		u, v := e[0], e[1]
		u--
		v--
		g.AddEdge(u, v)
		g.AddEdge(v, u)
	}

	D := make([]int, n)
	sum := make([]int64, n)

	P := make([][]int, n)
	for i := 0; i < n; i++ {
		P[i] = make([]int, H)
	}

	var dfs func(p, u int)

	dfs = func(p, u int) {
		P[u][0] = p
		for i := 1; i < H; i++ {
			P[u][i] = P[P[u][i-1]][i-1]
		}

		for i := g.nodes[u]; i > 0; i = g.next[i] {
			v := g.to[i]
			if p != v {
				sum[v] = sum[u] + int64(W[v])
				D[v] = D[u] + 1
				dfs(u, v)
			}
		}
	}
	sum[0] = int64(W[0])
	dfs(0, 0)

	lca := func(u, v int) int {
		if D[u] < D[v] {
			u, v = v, u
		}
		// D[u] >= D[v]
		for i := H - 1; i >= 0; i-- {
			if D[u]-(1<<uint(i)) >= D[v] {
				u = P[u][i]
			}
		}
		if u == v {
			return u
		}
		for i := H - 1; i >= 0; i-- {
			if P[u][i] != P[v][i] {
				u = P[u][i]
				v = P[v][i]
			}
		}
		return P[u][0]
	}

	X := make([][]int, n)
	R := make([]int64, n)
	var prod int64 = 1
	for i := 0; i < n; i++ {
		R[i] = int64(W[i])
		X[i] = []int{i, i}
		modMul(&prod, int64(W[i]))
	}

	ans := make([]int, n)
	ans[n-1] = int(prod)

	set := NewUF(n)

	join := func(u, v int) int {
		pu := set.Find(u)
		pv := set.Find(v)

		modMul(&prod, inverse(R[pu]))
		modMul(&prod, inverse(R[pv]))

		best := R[pu]
		x, y := X[pu][0], X[pu][1]
		if R[pv] > best {
			best = R[pv]
			x, y = X[pv][0], X[pv][1]
		}

		for i := 0; i < 2; i++ {
			for j := 0; j < 2; j++ {
				a, b := X[pu][i], X[pv][j]
				c := lca(a, b)
				tmp := sum[a] - sum[c] + sum[b] - sum[c] + int64(W[c])
				if tmp > best {
					best = tmp
					x, y = a, b
				}
			}
		}

		set.Union(pu, pv)
		pu = set.Find(pu)
		R[pu] = best
		X[pu][0] = x
		X[pu][1] = y

		modMul(&prod, best)

		return int(prod)
	}

	for i := n - 2; i >= 0; i-- {
		e := E[S[i]-1]
		u, v := e[0], e[1]
		u--
		v--

		ans[i] = join(u, v)
	}
	return ans
}

type UF struct {
	arr []int
	cnt []int
}

func NewUF(n int) *UF {
	u := new(UF)
	u.arr = make([]int, n)
	u.cnt = make([]int, n)
	for i := 0; i < n; i++ {
		u.arr[i] = i
		u.cnt[i] = 1
	}
	return u
}

func (u *UF) Find(a int) int {
	if u.arr[a] != a {
		u.arr[a] = u.Find(u.arr[a])
	}
	return u.arr[a]
}

func (u *UF) Union(a, b int) bool {
	pa := u.Find(a)
	pb := u.Find(b)
	if pa == pb {
		return false
	}
	if u.cnt[pa] < u.cnt[pb] {
		pa, pb = pb, pa
	}
	u.cnt[pa] += u.cnt[pb]
	u.arr[pb] = pa
	return true
}

type Graph struct {
	nodes []int
	next  []int
	to    []int
	cur   int
}

func NewGraph(n int, e int) *Graph {
	g := new(Graph)
	g.nodes = make([]int, n)
	g.next = make([]int, e+2)
	g.to = make([]int, e+2)
	g.cur = 0
	return g
}

func (g *Graph) AddEdge(u, v int) {
	g.cur++
	g.next[g.cur] = g.nodes[u]
	g.nodes[u] = g.cur
	g.to[g.cur] = v
}

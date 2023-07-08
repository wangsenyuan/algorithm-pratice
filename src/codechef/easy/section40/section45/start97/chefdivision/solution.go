package main

import (
	"bufio"
	"bytes"
	"fmt"
	"math"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	n, k := readTwoNums(reader)
	A := readNNums(reader, n)
	E := make([][]int, n-1)
	for i := 0; i < n-1; i++ {
		E[i] = readNNums(reader, 2)
	}
	Q := make([][]int, k)
	for i := 0; i < k; i++ {
		s, _ := reader.ReadBytes('\n')
		if s[0] != '3' {
			var u, p int
			pos := readInt(s, 2, &u)
			readInt(s, pos+1, &p)
			Q[i] = []int{int(s[0] - '0'), u, p}
		} else {
			var u int
			readInt(s, 2, &u)
			Q[i] = []int{3, u}
		}
	}
	res := solve(n, E, A, Q)
	var buf bytes.Buffer

	for _, x := range res {
		buf.WriteString(fmt.Sprintf("%d\n", x))
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
	for i := 0; i < n; i++ {
		for x < len(bs) && (bs[x] < '0' || bs[x] > '9') && bs[x] != '-' {
			x++
		}
		x = readInt(bs, x, &res[i])
	}
	return res
}

type Item struct {
	first  uint64
	second uint64
}

func All() Item {
	return Item{math.MaxUint64, math.MaxUint64}
}

func None() Item {
	return Item{0, 0}
}

func (it *Item) Set(i int) {
	if i < 63 {
		it.first |= 1 << i
	} else {
		it.second |= 1 << (i - 63)
	}
}

func (it *Item) Unset(i int) {
	if i < 63 {
		it.first ^= 1 << i
	} else {
		it.second ^= 1 << (i - 63)
	}
}

func (it Item) Count() int {
	return digit_count(it.first) + digit_count(it.second)
}

func digit_count(num uint64) int {
	var res int
	for i := 0; i < 63; i++ {
		if (num>>i)&1 == 1 {
			res++
		}
	}
	return res
}

func (this Item) Or(that Item) Item {
	a := this.first | that.first
	b := this.second | that.second
	return Item{a, b}
}

func (this Item) And(that Item) Item {
	a := this.first & that.first
	b := this.second & that.second
	return Item{a, b}
}

type SegTree struct {
	nodes []Item
	lazy  [][]Item
	n     int
}

func BuildTree(a []Item) *SegTree {
	n := len(a)
	nodes := make([]Item, 4*n)
	lazy := make([][]Item, 4*n)
	var build func(i int, l int, r int)
	build = func(i int, l int, r int) {
		lazy[i] = make([]Item, 2)
		lazy[i][0] = All()
		lazy[i][1] = None()
		if l == r {
			nodes[i] = a[l]
			return
		}
		mid := (l + r) / 2
		build(i*2, l, mid)
		build(i*2+1, mid+1, r)
		nodes[i] = nodes[2*i].And(nodes[2*i+1])
	}

	build(1, 0, n-1)

	return &SegTree{nodes, lazy, n}
}

func (t *SegTree) pushValue(v int, it []Item) {
	t.nodes[v] = t.nodes[v].And(it[0])
	t.lazy[v][0] = t.lazy[v][0].And(it[0])
	t.lazy[v][1] = t.lazy[v][1].And(it[0])
	t.nodes[v] = t.nodes[v].Or(it[1])
	t.lazy[v][0] = t.lazy[v][0].Or(it[1])
	t.lazy[v][1] = t.lazy[v][1].Or(it[1])
}

func (t *SegTree) push(v int) {
	t.pushValue(2*v, t.lazy[v])
	t.pushValue(2*v+1, t.lazy[v])
	t.lazy[v][0] = All()
	t.lazy[v][1] = None()
}

func (t *SegTree) update(L int, R int, a Item, f func(Item, Item) Item) {
	var loop func(i int, l, r, L, R int)

	loop = func(i int, l, r, L, R int) {
		if L > R {
			return
		}
		if l == L && r == R {
			t.nodes[i] = f(t.nodes[i], a)
			t.lazy[i][0] = f(t.lazy[i][0], a)
			t.lazy[i][1] = f(t.lazy[i][1], a)
			return
		}
		t.push(i)
		mid := (l + r) / 2
		loop(2*i, l, mid, L, min(R, mid))
		loop(2*i+1, mid+1, r, max(mid+1, L), R)
		t.nodes[i] = t.nodes[2*i].And(t.nodes[2*i+1])
	}

	loop(1, 0, t.n-1, L, R)
}

func (t *SegTree) AndUpdate(L int, R int, a Item) {
	t.update(L, R, a, func(a, b Item) Item {
		return a.And(b)
	})
}

func (t *SegTree) OrUpdate(L int, R int, a Item) {
	t.update(L, R, a, func(a, b Item) Item {
		return a.Or(b)
	})
}

func (t *SegTree) Query(L int, R int) Item {
	var loop func(i int, l int, r int, L int, R int) Item
	loop = func(i int, l int, r int, L int, R int) Item {
		if L > R {
			return All()
		}
		if l == L && r == R {
			return t.nodes[i]
		}
		t.push(i)
		mid := (l + r) / 2
		a := loop(2*i, l, mid, L, min(R, mid))
		b := loop(2*i+1, mid+1, r, max(mid+1, L), R)
		return a.And(b)
	}
	return loop(1, 0, t.n-1, L, R)
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

var primes []int
var primes_pos []int

func init() {
	primes_pos = make([]int, 700)
	set := make([]bool, 700)
	for i := 2; i < 700; i++ {
		if !set[i] {
			primes_pos[i] = len(primes)
			primes = append(primes, i)
			for j := i * i; j < 700; j += i {
				set[j] = true
			}
		}
	}
}

func solve(n int, E [][]int, A []int, Q [][]int) []int {
	g := NewGraph(n, 2*n)
	for _, e := range E {
		u, v := e[0]-1, e[1]-1
		g.AddEdge(u, v)
		g.AddEdge(v, u)
	}

	in := make([]int, n)
	out := make([]int, n)
	var time int
	var dfs func(p int, u int)

	dfs = func(p int, u int) {
		in[u] = time
		time++
		for i := g.nodes[u]; i > 0; i = g.next[i] {
			v := g.to[i]
			if p != v {
				dfs(u, v)
			}
		}
		out[u] = time - 1
	}

	dfs(0, 0)

	B := make([]Item, n)
	for i := 0; i < n; i++ {
		u := in[i]
		it := None()
		for j := 0; j < 125; j++ {
			if A[i]%primes[j] == 0 {
				it.Set(j)
			}
		}
		B[u] = it
	}

	st := BuildTree(B)

	var ans []int

	for _, q := range Q {
		if q[0] == 1 {
			u, p := q[1], q[2]
			u--
			tmp := None()
			tmp.Set(primes_pos[p])
			st.OrUpdate(in[u], out[u], tmp)
		} else if q[0] == 2 {
			u, p := q[1], q[2]
			u--
			tmp := All()
			tmp.Unset(primes_pos[p])
			st.AndUpdate(in[u], out[u], tmp)
		} else {
			u := q[1]
			u--
			tmp := st.Query(in[u], out[u])
			ans = append(ans, tmp.Count())
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

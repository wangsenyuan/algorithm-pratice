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

	var buf bytes.Buffer

	n := readNum(reader)
	C := readNNums(reader, n)
	X := readNNums(reader, n)
	E := make([][]int, n-1)
	for i := 0; i < n-1; i++ {
		E[i] = readNNums(reader, 2)
	}

	res := solve(n, E, C, X)

	for _, num := range res {
		buf.WriteString(fmt.Sprintf("%d ", num))
	}

	buf.WriteByte('\n')
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

const N = 200005

func solve(n int, E [][]int, C []int, X []int) []int {

	g := NewGraph(n, 2*len(E)+1)
	for _, e := range E {
		u, v := e[0], e[1]
		u--
		v--
		g.AddEdge(u, v)
		g.AddEdge(v, u)
	}

	in := make([]int, n)
	out := make([]int, n)

	var dfs func(p, u int, time *int)

	ord := make([]int, 0, n)

	dfs = func(p, u int, time *int) {
		ord = append(ord, u)
		in[u] = *time
		*time++
		for i := g.node[u]; i > 0; i = g.next[i] {
			v := g.to[i]
			if p != v {
				dfs(u, v, time)
			}
		}
		out[u] = *time
	}

	dfs(-1, 0, new(int))

	cp := make([]Pair, n)
	for i := 0; i < n; i++ {
		cp[i] = Pair{C[i], i}
	}

	sort.Sort(Pairs(cp))

	my := make(map[int]int)

	for i := 1; i <= n; i++ {
		if i == n || cp[i].first != cp[i-1].first {
			my[cp[i-1].first] = len(my)
		}
	}

	for i := 0; i < n; i++ {
		C[i] = my[C[i]]
	}

	for i := 0; i < n; i++ {
		ord[i] = C[ord[i]]
	}

	Q := make([]Query, n)

	for i := 0; i < n; i++ {
		Q[i] = Query{in[i], out[i] - 1, i, X[i]}
	}

	sort.Sort(Queries(Q))

	ans := make([]int, n)

	freq := make([]int, n+1)
	cnt := make([]int, n+1)

	currL, currR := 0, 0
	for i := 0; i < n; i++ {
		L, R := Q[i].l, Q[i].r
		for currR <= R {
			// currSum += a[currR];
			cnt[freq[ord[currR]]]--
			freq[ord[currR]]++
			cnt[freq[ord[currR]]]++
			currR++
		}

		for currR > R+1 {
			//  currSum -= a[currR-1];
			cnt[freq[ord[currR-1]]]--
			freq[ord[currR-1]]--
			cnt[freq[ord[currR-1]]]++
			currR--
		}

		for currL < L {
			cnt[freq[ord[currL]]]--
			freq[ord[currL]]--
			cnt[freq[ord[currL]]]++
			currL++
		}
		for currL > L {
			// currSum += a[currL-1];
			cnt[freq[ord[currL-1]]]--
			freq[ord[currL-1]]++
			cnt[freq[ord[currL-1]]]++
			currL--
		}

		if Q[i].x < len(cnt) {
			ans[Q[i].id] = cnt[Q[i].x]
		}
	}

	return ans
}

type Pair struct {
	first  int
	second int
}

func (this Pair) Less(that Pair) bool {
	return this.first < that.first
}

type Pairs []Pair

func (ps Pairs) Len() int {
	return len(ps)
}

func (ps Pairs) Less(i, j int) bool {
	return ps[i].Less(ps[j])
}

func (ps Pairs) Swap(i, j int) {
	ps[i], ps[j] = ps[j], ps[i]
}

const BLOCK_SIZE = 448

type Query struct {
	l  int
	r  int
	id int
	x  int
}

func (this Query) Less(that Query) bool {
	if this.l/BLOCK_SIZE != that.l/BLOCK_SIZE {
		return this.l/BLOCK_SIZE < that.l/BLOCK_SIZE
	}
	return this.r < that.r
}

type Queries []Query

func (qs Queries) Len() int {
	return len(qs)
}

func (qs Queries) Less(i, j int) bool {
	return qs[i].Less(qs[j])
}

func (qs Queries) Swap(i, j int) {
	qs[i], qs[j] = qs[j], qs[i]
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
	g.next = make([]int, e)
	g.to = make([]int, e)
	g.cur = 0
	return g
}

func (g *Graph) AddEdge(u, v int) {
	g.cur++
	g.next[g.cur] = g.node[u]
	g.node[u] = g.cur
	g.to[g.cur] = v
}

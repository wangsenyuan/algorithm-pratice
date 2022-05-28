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

	n, h := readTwoNums(reader)
	E := make([][]int, n-1)
	for i := 0; i < n-1; i++ {
		E[i] = readNNums(reader, 2)
	}

	mat := make([][]int, h)
	for i := 0; i < h; i++ {
		mat[i] = readNNums(reader, h)
	}
	q := readNum(reader)
	F := make([][]int, q)
	S := make([][]int, q)

	for i := 0; i < q; i++ {
		F[i] = readNNums(reader, 2)
		S[i] = readNNums(reader, F[i][0])
	}

	res := solve(n, E, mat, q, F, S)

	var buf bytes.Buffer

	for _, cur := range res {
		buf.WriteString(cur)
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

const H = 20

func solve(n int, E [][]int, mat [][]int, q int, F [][]int, S [][]int) []string {
	g := NewGraph(n, len(E)*2)

	for _, e := range E {
		u, v := e[0], e[1]
		u--
		v--
		g.AddEdge(u, v)
		g.AddEdge(v, u)
	}

	D := make([]int, n)
	P := make([][]int, n)
	in := make([]int, n)
	out := make([]int, n)
	var dfs func(p, u int)
	var time int
	dfs = func(p, u int) {
		in[u] = time
		time++
		P[u] = make([]int, H)
		P[u][0] = p
		for i := 1; i < H; i++ {
			P[u][i] = P[P[u][i-1]][i-1]
		}
		for i := g.node[u]; i > 0; i = g.next[i] {
			v := g.to[i]
			if p != v {
				D[v] = D[u] + 1
				dfs(u, v)
			}
		}
		out[u] = time
	}

	dfs(0, 0)

	isAncestor := func(u, v int) bool {
		return in[u] <= in[v] && out[v] <= out[u]
	}

	lca := func(u, v int) int {
		if D[u] < D[v] {
			u, v = v, u
		}
		for i := H - 1; i >= 0; i-- {
			if (D[u] - (1 << uint(i))) >= D[v] {
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

	orig := make([]int, n+5)

	wt := make([]int, n+5)

	sort_by_preorder := func(arr []int) {
		x := make([]Pair, len(arr))
		for i := 0; i < len(arr); i++ {
			x[i] = Pair{int64(in[arr[i]]), int64(arr[i])}
		}
		sort.Sort(Pairs(x))
		for i := 0; i < len(arr); i++ {
			arr[i] = int(x[i].second)
		}
	}

	stack := make([]int, n)

	tree := make([]map[int]bool, n+5)

	for i := 0; i < len(tree); i++ {
		tree[i] = make(map[int]bool)
	}

	carry := make([]int64, n+5)

	ans := make([]int64, n+5)

	var dfs_aux func(s int) []Pair

	dfs_aux = func(s int) []Pair {
		var res []Pair
		for u, _ := range tree[s] {
			tmp := dfs_aux(u)
			carry[s] += carry[u]
			res = append(res, tmp...)
		}
		ans[s] += carry[s]
		res = append(res, Pair{ans[s], int64(wt[s])})
		return res
	}

	build_tree := func(k int, mp map[int]int, arr []int) []Pair {
		done := make(map[int]bool)

		sort_by_preorder(arr)

		for i := 0; i < k; i++ {
			done[arr[i]] = true
		}

		for i := 0; i < k-1; i++ {
			v := lca(arr[i], arr[i+1])
			if !done[v] {
				arr = append(arr, v)
				done[v] = true
			}
		}
		for i := k; i < len(arr); i++ {
			mp[arr[i]] = i
			wt[i] = 1
		}
		sort_by_preorder(arr)

		cur := len(arr)
		p := 0
		stack[p] = arr[0]
		p++

		for i := 1; i < len(arr); i++ {
			for !isAncestor(stack[p-1], arr[i]) {
				p--
			}
			tree[mp[stack[p-1]]][cur] = true
			tree[cur][mp[arr[i]]] = true
			wt[cur] = D[arr[i]] - D[stack[p-1]] + 1
			cur++
			stack[p] = arr[i]
			p++
		}

		for i := 0; i < k; i++ {
			for j := i + 1; j < k; j++ {
				carry[mp[orig[i]]] += int64(mat[i][j])
				carry[mp[orig[j]]] += int64(mat[i][j])
				carry[mp[lca(orig[i], orig[j])]] -= 2 * int64(mat[i][j])
				ans[mp[lca(orig[i], orig[j])]] += int64(mat[i][j])
			}
		}

		tbs := dfs_aux(mp[arr[0]])

		for i := 0; i < cur; i++ {
			tree[i] = make(map[int]bool)
			carry[i] = 0
			ans[i] = 0
			orig[i] = 0
			wt[i] = 0
		}

		return tbs
	}

	calc := func(X int, arr []int) string {
		k := len(arr)
		mp := make(map[int]int)
		for j := 0; j < k; j++ {
			arr[j]--
			orig[j] = arr[j]
			mp[arr[j]] = j
			wt[j] = 1
		}

		tbs := build_tree(k, mp, arr)

		sort.Sort(Pairs(tbs))

		var cur int
		var cost int64

		for i := 0; i < len(tbs); i++ {
			cur += int(tbs[i].second)
		}
		if cur < X {
			return "No"
		}
		cur = 0
		var ind int
		for ind < len(tbs) && (cur+int(tbs[ind].second)) <= X {
			cur += int(tbs[ind].second)
			cost += tbs[ind].first * tbs[ind].second
			ind++
		}

		cost += int64(X-cur) * (tbs[ind].first)

		return fmt.Sprintf("%d", cost)
	}

	res := make([]string, len(F))

	for i := 0; i < len(F); i++ {
		k, x := F[i][0], F[i][1]

		if k == 1 {
			if x == 1 {
				res[i] = "0"
			} else {
				res[i] = "No"
			}
			continue
		}

		res[i] = calc(x, S[i])
	}

	return res
}

type Pair struct {
	first  int64
	second int64
}

type Pairs []Pair

func (ps Pairs) Len() int {
	return len(ps)
}

func (ps Pairs) Less(i, j int) bool {
	return ps[i].first < ps[j].first
}

func (ps Pairs) Swap(i, j int) {
	ps[i], ps[j] = ps[j], ps[i]
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

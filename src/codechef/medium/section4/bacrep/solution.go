package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

func readInt(bytes []byte, from int, val *int) int {
	i := from
	sign := 1
	if bytes[i] == '-' {
		sign = -1
		i++
	}
	tmp := 0
	for i < len(bytes) && bytes[i] != ' ' {
		tmp = tmp*10 + int(bytes[i]-'0')
		i++
	}
	*val = tmp * sign
	return i
}

func readNum(scanner *bufio.Scanner) (a int) {
	scanner.Scan()
	readInt(scanner.Bytes(), 0, &a)
	return
}

func readTwoNums(scanner *bufio.Scanner) (a int, b int) {
	res := readNNums(scanner, 2)
	a, b = res[0], res[1]
	return
}

func readNNums(scanner *bufio.Scanner, n int) []int {
	res := make([]int, n)
	x := 0
	scanner.Scan()
	for i := 0; i < n; i++ {
		for x < len(scanner.Bytes()) && scanner.Bytes()[x] == ' ' {
			x++
		}
		x = readInt(scanner.Bytes(), x, &res[i])
	}
	return res
}

func readInt64(bytes []byte, from int, val *int64) int {
	i := from
	sign := int64(1)
	if bytes[i] == '-' {
		sign = -1
		i++
	}
	tmp := int64(0)
	for i < len(bytes) && bytes[i] != ' ' {
		tmp = tmp*10 + int64(bytes[i]-'0')
		i++
	}
	*val = tmp * sign
	return i
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	N, Q := readTwoNums(scanner)
	edges := make([][]int, N-1)
	for i := 0; i < N-1; i++ {
		edges[i] = readNNums(scanner, 2)
	}
	A := readNNums(scanner, N)
	queries := make([]string, Q)

	for i := 0; i < Q; i++ {
		scanner.Scan()
		queries[i] = scanner.Text()
	}

	res := solve(N, Q, A, edges, queries)

	for i := 0; i < len(res); i++ {
		fmt.Println(res[i])
	}
}

func solve(N, Q int, A []int, edges [][]int, queries []string) []int64 {
	conns := make([][]int, N)

	for i := 0; i < N; i++ {
		conns[i] = make([]int, 0, 3)
	}

	for _, edge := range edges {
		u, v := edge[0]-1, edge[1]-1
		conns[u] = append(conns[u], v)
		conns[v] = append(conns[v], u)
	}

	open := make([]int, N)
	close := make([]int, N)
	depth := make([]int, N)
	leaf := make([]bool, N)

	var dfs func(p int, u, d int, time *int)

	dfs = func(p, u, d int, time *int) {
		open[u] = *time
		*time++
		depth[u] = d
		leaf[u] = true
		for _, v := range conns[u] {
			if p == v {
				continue
			}
			leaf[u] = false
			dfs(u, v, d+1, time)
		}
		close[u] = *time
	}

	dfs(-1, 0, 0, new(int))

	QS := make([]Query, 0, N+Q)

	for i := 0; i < N; i++ {
		QS = append(QS, NewQuery(-1, -depth[i], i, A[i], 0, '+'))
	}

	for i := 0; i < Q; i++ {
		query := []byte(queries[i])
		op := query[0]
		var v int
		var k int
		if op == '+' {
			j := readInt(query, 2, &v)
			readInt(query, j+1, &k)
		} else {
			readInt(query, 2, &v)
		}
		v--
		QS = append(QS, NewQuery(i, i+1-depth[v], v, k, 0, op))
	}

	sort.Sort(Queries(QS))

	tree := make([]*FenwickTree, 2)
	tree[0] = NewFenwickTree()
	tree[1] = NewFenwickTree()

	for i := 0; i < len(QS); i++ {
		if i > 0 && QS[i].time > QS[i-1].time {
			tree[1].Clear()
		}

		if QS[i].op == '+' {
			tree[0].Add(open[QS[i].v], close[QS[i].v], int64(QS[i].k))
			tree[1].Add(open[QS[i].v], close[QS[i].v], int64(QS[i].k))
		} else {
			var j = 1
			if leaf[QS[i].v] {
				j = 0
			}
			QS[i].answer = tree[j].Get(open[QS[i].v])
		}
	}

	res := make([]int64, Q)

	for i := 0; i < Q; i++ {
		res[i] = -1
	}

	for i := 0; i < len(QS); i++ {
		if QS[i].op == '?' {
			res[QS[i].id] = QS[i].answer
		}
	}

	var j int

	for i := 0; i < Q; i++ {
		if res[i] < 0 {
			continue
		}
		res[j] = res[i]
		j++
	}

	return res[:j]
}

type Query struct {
	id, time, v, k int
	answer         int64
	op             byte
}

func NewQuery(id, time, v, k int, answer int64, op byte) Query {
	return Query{id, time, v, k, answer, op}
}

type Queries []Query

func (qs Queries) Len() int {
	return len(qs)
}

func (qs Queries) Less(i, j int) bool {
	if qs[i].time < qs[j].time {
		return true
	}
	if qs[i].time == qs[j].time {
		return qs[i].op == '+' && qs[j].op == '?'
	}
	return false
}

func (qs Queries) Swap(i, j int) {
	qs[i], qs[j] = qs[j], qs[i]
}

type FenwickTree struct {
	arr     []int64
	mark    []bool
	history []int
	p       int
}

const MAX_N = 2000010

func NewFenwickTree() *FenwickTree {
	arr := make([]int64, MAX_N)
	mark := make([]bool, MAX_N)
	history := make([]int, MAX_N)

	return &FenwickTree{arr, mark, history, 0}
}

func (tree *FenwickTree) Add(L, R int, V int64) {

	add := func(x int, v int64) {
		x += 5
		if !tree.mark[x] {
			tree.mark[x] = true
			tree.history[tree.p] = x
			tree.p++
		}
		for i := x; i < MAX_N; i += i & -i {
			tree.arr[i] += v
		}
	}

	add(L, V)
	add(R, -V)
}

func (tree *FenwickTree) Get(x int) int64 {
	x += 5
	var res int64

	for i := x; i > 0; i -= i & -i {
		res += tree.arr[i]
	}
	return res
}

func (tree *FenwickTree) Clear() {
	for i := 0; i < tree.p; i++ {
		a := tree.history[i]
		tree.mark[a] = false

		for j := a; j < MAX_N; j += j & -j {
			tree.arr[j] = 0
		}
	}
	tree.p = 0
}

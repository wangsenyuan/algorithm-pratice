package main

import (
	"bufio"
	"fmt"
	"os"
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

func fillNNums(scanner *bufio.Scanner, n int, res []int) {
	x := 0
	scanner.Scan()
	for i := 0; i < n; i++ {
		for x < len(scanner.Bytes()) && scanner.Bytes()[x] == ' ' {
			x++
		}
		x = readInt(scanner.Bytes(), x, &res[i])
	}
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	tc := readNum(scanner)

	for tc > 0 {
		tc--
		n, q := readTwoNums(scanner)
		A := readNNums(scanner, n)
		E := make([][]int, n-1)
		for i := 0; i < n-1; i++ {
			E[i] = readNNums(scanner, 2)
		}
		solver := NewSolver1(n, A, E)

		var v, k int

		for q > 0 {
			q--
			a, b := readTwoNums(scanner)
			v ^= a
			k ^= b
			v, k = solver.Query(v, k)
			fmt.Printf("%d %d\n", v, k)
		}

		solver.Reset()
	}
}

const MAX_N = 200005

var nodes []*Node
var visit []bool
var G [][]int

var L []int
var R []int
var mn []int
var rt []int

const MAX_M = 1 << 20
const MAX_S = 2 * 20 * 20 * MAX_N

var sub []map[int]bool

func init() {
	nodes = make([]*Node, MAX_N)
	visit = make([]bool, MAX_N)
	G = make([][]int, MAX_N)
	for i := 0; i < MAX_N; i++ {
		G[i] = make([]int, 0, 3)
	}
	L = make([]int, MAX_S)
	R = make([]int, MAX_S)
	mn = make([]int, MAX_S)
	for i := 0; i < MAX_S; i++ {
		mn[i] = MAX_M
	}
	rt = make([]int, MAX_N)
	sub = make([]map[int]bool, MAX_N)
}

type Solver1 struct {
	L, R, mn, rt []int
	sz           int
}

func NewSolver1(n int, W []int, E [][]int) Solver1 {
	sz := 1

	for i := 0; i <= n; i++ {
		G[i] = G[i][:0]
		sub[i] = make(map[int]bool)
	}

	for _, e := range E {
		u, v := e[0], e[1]
		G[u] = append(G[u], v)
		G[v] = append(G[v], u)
	}

	copy := func(v int) int {
		L[sz] = L[v]
		R[sz] = R[v]
		mn[sz] = mn[v]
		sz++
		return sz - 1
	}

	var update func(p, c, v int, l, r int)

	update = func(p, c, v int, l, r int) {
		mn[v] = min(mn[v], c)
		if r-l > 1 {
			m := (l + r) / 2
			if p < m {
				L[v] = copy(L[v])
				update(p, c, L[v], l, m)
			} else {
				R[v] = copy(R[v])
				update(p, c, R[v], m, r)
			}
		}
	}

	var dfs func(p, u int)

	dfs = func(p, u int) {
		rt[u] = copy(0)
		sub[u][u] = true

		update(W[u-1], u, rt[u], 0, MAX_M)

		for _, v := range G[u] {
			if p == v {
				continue
			}
			dfs(u, v)

			if len(sub[u]) < len(sub[v]) {
				rt[u] = copy(rt[v])
				sub[u], sub[v] = sub[v], sub[u]
			}

			for it := range sub[v] {
				sub[u][it] = true
				update(W[it-1], it, rt[u], 0, MAX_M)
			}
		}
	}

	dfs(0, 1)

	return Solver1{L, R, mn, rt, sz}
}

func (solver Solver1) Query(v, k int) (int, int) {
	var get func(v int, l, r int) (int, int)
	get = func(v int, l, r int) (int, int) {
		if r-l == 1 {
			return mn[v], l
		}
		m := (l + r) / 2
		if ((k^(m-l)) < k && L[v] > 0) || R[v] == 0 {
			return get(L[v], l, m)
		}
		return get(R[v], m, r)
	}

	x, y := get(rt[v], 0, MAX_M)
	y ^= k
	return x, y
}

func (solver *Solver1) Reset() {
	for i := 0; i < solver.sz; i++ {
		L[i] = 0
		R[i] = 0
		mn[i] = MAX_M
	}
}

func min(a, b int) int {
	if a <= b {
		return a
	}
	return b
}

type Solver struct {
	nodes []*Node
}

func NewSolver(n int, W []int, E [][]int) Solver {
	for i := 0; i < n; i++ {
		visit[i] = false
		G[i] = G[i][:0]
	}

	for _, e := range E {
		u, v := e[0]-1, e[1]-1
		G[u] = append(G[u], v)
		G[v] = append(G[v], u)
	}

	var dfs func(v int)

	dfs = func(v int) {
		visit[v] = true
		nodes[v] = NewNode(v, W[v])
		nodes[v].length = 20

		for _, u := range G[v] {
			if !visit[u] {
				dfs(u)
				nodes[v] = merge(nodes[v], nodes[u])
			}
		}
	}

	dfs(0)

	return Solver{nodes[:n]}
}

func (solver Solver) Query(v, k int) (int, int) {
	nodes := solver.nodes

	v--
	node := nodes[v]

	j := 1 << 19

	for node.left != nil || node.right != nil {
		if k&j > 0 && node.left != nil || node.right == nil {
			node = node.left
		} else {
			node = node.right
		}
		j >>= 1
	}
	v = node.vertex + 1
	k ^= node.weight

	return v, k
}

type Node struct {
	left, right *Node
	vertex      int
	weight      int
	length      int
}

func NewNode(vertex, weight int) *Node {
	node := new(Node)
	node.vertex = vertex
	node.weight = weight
	node.length = 0
	return node
}

func (a *Node) Split() {
	tmp := NewNode(a.vertex, a.weight)
	tmp.length = a.length - 1
	if a.weight&(1<<uint(a.length-1)) > 0 {
		a.right = tmp
	} else {
		a.left = tmp
	}
	a.length = 0
	a.vertex = 0
}

func merge(a, b *Node) *Node {
	if a == nil {
		return b
	}
	if b == nil {
		return a
	}

	if a.length > 0 {
		a.Split()
	}

	if b.length > 0 {
		b.Split()
	}

	if a.left == nil && a.right == nil {
		if a.vertex > b.vertex {
			return b
		}
		return a
	}

	tmp := new(Node)
	tmp.left = merge(a.left, b.left)
	tmp.right = merge(a.right, b.right)
	tmp.length = 0
	return tmp
}

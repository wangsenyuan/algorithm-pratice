package main

import (
	"bufio"
	"os"
	"fmt"
	"math"
)

func readInt(bytes []byte, from int, val *int) int {
	i := from
	tmp := 0
	for i < len(bytes) && bytes[i] != ' ' {
		tmp = tmp*10 + int(bytes[i]-'0')
		i++
	}
	*val = tmp
	return i
}

func readNum(scanner *bufio.Scanner) (a int) {
	scanner.Scan()
	readInt(scanner.Bytes(), 0, &a)
	return
}

func readTwoNums(scanner *bufio.Scanner) (a int, b int) {
	scanner.Scan()
	x := readInt(scanner.Bytes(), 0, &a)
	readInt(scanner.Bytes(), x+1, &b)
	return
}

func readNNums(scanner *bufio.Scanner, n int) []int {
	res := make([]int, n)
	x := -1
	scanner.Scan()
	for i := 0; i < n; i++ {
		x = readInt(scanner.Bytes(), x+1, &res[i])
	}
	return res
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	n := readNum(scanner)
	pairs := make([][]int, n-1)
	for i := 0; i < n-1; i++ {
		pairs[i] = readNNums(scanner, 2)
	}
	m := readNum(scanner)
	queries := make([][]int, m)
	for i := 0; i < m; i++ {
		queries[i] = readNNums(scanner, 3)
	}
	res := solve(n, pairs, queries)
	for i := 0; i < m; i++ {
		fmt.Println(res[i])
	}
}

func solve(n int, pairs [][]int, queries [][]int) []int {
	neighbors := make([]map[int]bool, n)

	for i := 0; i < n; i++ {
		neighbors[i] = make(map[int]bool)
	}

	for _, pair := range pairs {
		u, v := pair[0]-1, pair[1]-1
		neighbors[u][v] = true
		neighbors[v][u] = true
	}

	var eulerTour func(p, u int, level int)
	E := make([]int, 2*n-1)
	L := make([]int, 2*n-1)
	H := make([]int, n)

	for i := 0; i < n; i++ {
		H[i] = -1
	}

	var index int

	eulerTour = func(p, u int, level int) {
		index++
		E[index-1] = u
		L[index-1] = level

		if H[u] < 0 {
			H[u] = index - 1
		}

		for v := range neighbors[u] {
			if p == v {
				continue
			}
			eulerTour(u, v, level+1)
			index++
			E[index-1] = u
			L[index-1] = level
		}
	}

	eulerTour(-1, 0, 0)

	rmq := Construct(L)

	lca := func(u, v int) int {
		a, b := H[u], H[v]
		if a > b {
			a, b = b, a
		}
		res := rmq.Query(a, b)
		return E[res]
	}

	res := make([]int, len(queries))

	for i, query := range queries {
		r, u, v := query[0]-1, query[1]-1, query[2]-1
		uv := lca(u, v)
		ru := lca(r, u)
		rv := lca(r, v)
		ruv := lca(r, uv)

		if ruv != uv {
			res[i] = uv + 1
		} else {
			if ru == uv && rv == uv {
				res[i] = uv + 1
			} else if rv == uv {
				res[i] = ru + 1
			} else {
				res[i] = rv + 1
			}
		}
	}

	return res
}

type RMQ struct {
	size int
	tree []int
	arr  []int
}

func Construct(arr []int) *RMQ {
	size := len(arr)

	var rec func(i int, left, right int)

	tree := make([]int, 4*size)
	rec = func(i int, left, right int) {
		if left == right {
			tree[i] = left
			return
		}
		mid := (left + right) / 2
		rec(2*i+1, left, mid)
		rec(2*i+2, mid+1, right)
		if arr[tree[2*i+1]] < arr[tree[2*i+2]] {
			tree[i] = tree[2*i+1]
		} else {
			tree[i] = tree[2*i+2]
		}
	}

	rec(0, 0, size-1)
	return &RMQ{size, tree, arr}
}

func (rmq *RMQ) Query(left, right int) int {
	var rec func(i int, s, e int) int

	rec = func(i int, s, e int) int {
		if s > right || e < left {
			return -1
		}
		if s >= left && e <= right {
			return rmq.tree[i]
		}
		mid := (s + e) / 2
		q1 := rec(2*i+1, s, mid)
		q2 := rec(2*i+2, mid+1, e)
		if q1 < 0 {
			return q2
		}
		if q2 < 0 {
			return q1
		}
		if rmq.arr[q1] < rmq.arr[q2] {
			return q1
		}
		return q2
	}

	return rec(0, 0, rmq.size-1)
}

func solve2(n int, pairs [][]int, queries [][]int) []int {
	neighbors := make([]map[int]bool, n)

	for i := 0; i < n; i++ {
		neighbors[i] = make(map[int]bool)
	}

	for _, pair := range pairs {
		u, v := pair[0]-1, pair[1]-1
		neighbors[u][v] = true
		neighbors[v][u] = true
	}

	level := make([]int, n)
	parent := make([]int, n)

	var height func(p, u, l int) int

	height = func(p, u, l int) int {
		ans := 0
		level[u] = l
		for v := range neighbors[u] {
			if v == p {
				continue
			}
			parent[v] = u
			tmp := height(u, v, l+1)
			if tmp > ans {
				ans = tmp
			}
		}

		return ans + 1
	}

	maxHeight := height(-1, 0, 0)

	nr := int(math.Sqrt(float64(maxHeight)))

	var dfs func(p, u int)

	P := make([]int, n)

	dfs = func(p, u int) {
		if level[u] < nr {
			P[u] = 0
		} else {
			if level[u]%nr == 0 {
				P[u] = parent[u]
			} else {
				P[u] = P[parent[u]]
			}
		}

		for v := range neighbors[u] {
			if p == v {
				continue
			}
			dfs(u, v)
		}
	}

	dfs(-1, 0)

	lca := func(u, v int) int {
		for P[u] != P[v] {
			if level[u] > level[v] {
				u = P[u]
			} else {
				v = P[v]
			}
		}

		for u != v {
			if level[u] > level[v] {
				u = parent[u]
			} else {
				v = parent[v]
			}
		}
		return u
	}

	res := make([]int, len(queries))

	for i, query := range queries {
		r, u, v := query[0]-1, query[1]-1, query[2]-1
		uv := lca(u, v)
		ru := lca(r, u)
		rv := lca(r, v)
		ruv := lca(r, uv)

		if ruv != uv {
			res[i] = uv + 1
		} else {
			if ru == uv && rv == uv {
				res[i] = uv + 1
			} else if rv == uv {
				res[i] = ru + 1
			} else {
				res[i] = rv + 1
			}
		}
	}

	return res
}

func solve1(n int, pairs [][]int, queries [][]int) []int {
	neighbors := make([]map[int]bool, n)

	for i := 0; i < n; i++ {
		neighbors[i] = make(map[int]bool)
	}

	for _, pair := range pairs {
		u, v := pair[0]-1, pair[1]-1
		neighbors[u][v] = true
		neighbors[v][u] = true
	}

	level := make([]int, n)
	parent := make([]int, n)

	maxLevel := 0
	for 1<<uint(maxLevel) <= n {
		maxLevel++
	}

	pp := make([][]int, maxLevel)
	for i := 0; i < maxLevel; i++ {
		pp[i] = make([]int, n)
		for j := 0; j < n; j++ {
			pp[i][j] = -1
		}
	}

	que := make([]int, n)
	front, tail := 0, 0
	que[tail] = 0
	tail++
	seen := make([]bool, n)
	seen[0] = true
	for front < tail {
		tt := tail
		for front < tt {
			u := que[front]
			front++
			for v := range neighbors[u] {
				if !seen[v] {
					seen[v] = true
					parent[v] = u
					level[v] = level[u] + 1
					pp[0][v] = u
					que[tail] = v
					tail++
				}
			}
		}
	}

	for i := 1; i < maxLevel; i++ {
		for j := 0; j < n; j++ {
			if pp[i-1][j] >= 0 {
				pp[i][j] = pp[i-1][pp[i-1][j]]
			}
		}
	}

	lca := func(p, q int) int {
		if level[p] < level[q] {
			p, q = q, p
		}

		var log int

		for 1<<uint(log) <= level[p] {
			log++
		}
		log--

		for i := log; i >= 0; i-- {
			if level[p]-(1<<uint(i)) >= level[q] {
				p = pp[i][p]
			}
		}

		if p == q {
			return p
		}

		for i := log; i >= 0; i-- {
			if pp[i][p] != -1 && pp[i][p] != pp[i][q] {
				p = pp[i][p]
				q = pp[i][q]
			}
		}
		return parent[p]
	}

	res := make([]int, len(queries))

	for i, query := range queries {
		r, u, v := query[0]-1, query[1]-1, query[2]-1
		uv := lca(u, v)
		ru := lca(r, u)
		rv := lca(r, v)
		ruv := lca(r, uv)

		if ruv != uv {
			res[i] = uv + 1
		} else {
			if ru == uv && rv == uv {
				res[i] = uv + 1
			} else if rv == uv {
				res[i] = ru + 1
			} else {
				res[i] = rv + 1
			}
		}
	}

	return res
}

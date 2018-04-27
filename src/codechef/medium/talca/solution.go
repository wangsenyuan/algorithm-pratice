package main

import (
	"bufio"
	"os"
	"fmt"
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

	var eularTour func(p, u int, level int)

	first := make([]int, n)
	for i := 0; i < n; i++ {
		first[i] = -1
	}

	var index int

	nodes := make([]Node, 0, 3*n)

	eularTour = func(p, u int, level int) {
		index++
		if first[u] < 0 {
			first[u] = index - 1
		}
		nodes = append(nodes, Node{u, level})
		for v := range neighbors[u] {
			if v == p {
				continue
			}
			eularTour(u, v, level+1)
			level++
			nodes = append(nodes, Node{u, level})
		}
	}

	N := len(nodes)
	trees := make([]int, 4*N)

	var construct func(i int, left, right int)
	construct = func(i int, left, right int) {
		if left+1 == right {
			trees[i] = left
			return
		}
		mid := (left + right) / 2
		construct(2*i+1, left, mid)
		construct(2*i+2, mid, right)
		a := nodes[trees[2*i+1]]
		b := nodes[trees[2*i+2]]
		if a.level < b.level {
			trees[i] = trees[2*i+1]
		} else {
			trees[i] = trees[2*i+2]
		}
	}

	construct(0, 0, N)

	var query func(i int, s, e int, left, right int) int
	query = func(i int, s, e int, left, right int) int {
		if e <= left || s >= right {
			return -1
		}
		if s >= left && e <= right {
			return trees[i]
		}

		mid := (s + e) / 2
		a := query(2*i+1, s, mid, left, right)
		b := query(2*i+2, mid, e, left, right)
		if a < 0 {
			return b
		}
		if b < 0 {
			return a
		}
		c, d := nodes[a], nodes[b]
		if c.level < d.level {
			return a
		}
		return b
	}

	lca := func(u, v int) int {
		return query(0, 0, N, first[u], first[v])
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

type Node struct {
	index, level int
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

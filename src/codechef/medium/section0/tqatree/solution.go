package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
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
	edges := make([][]int, n-1)
	for i := 0; i < n-1; i++ {
		edges[i] = readNNums(scanner, 3)
	}
	q := readNum(scanner)
	queries := make([][]int, q)
	for i := 0; i < q; i++ {
		queries[i] = readNNums(scanner, 3)
	}
	res := solve(n, edges, queries)

	for _, ans := range res {
		fmt.Println(ans)
	}
}

func solve(n int, edges [][]int, queries [][]int) []int {
	neigbhors := make([]map[int]int, n)
	for i := 0; i < n; i++ {
		neigbhors[i] = make(map[int]int)
	}
	costs := make([]int, len(edges))
	for i, edge := range edges {
		u, v, w := edge[0]-1, edge[1]-1, edge[2]
		neigbhors[u][v] = w
		neigbhors[v][u] = w
		costs[i] = w
	}

	var height func(p, u int) int

	height = func(p, u int) int {
		var h int
		for v := range neigbhors[u] {
			if p == v {
				continue
			}
			hh := height(u, v)
			if hh > h {
				h = hh
			}
		}
		return h + 1
	}

	maxHeight := height(-1, 0)

	sqrtHeight := int(math.Sqrt(float64(maxHeight)))
	if sqrtHeight*sqrtHeight < maxHeight {
		sqrtHeight++
	}

	var dfs func(p, u, h int)

	dists := make([]int, n)
	P := make([]int, n)
	T := make([]int, n)
	L := make([]int, n)
	// O := make([]int, n)
	start := make([]int, n)
	end := make([]int, n)
	var idx int
	dfs = func(p, u, h int) {
		L[u] = h
		if h < sqrtHeight {
			T[u] = 0
		} else if h%sqrtHeight == 0 {
			T[u] = P[u]
		} else {
			T[u] = T[P[u]]
		}
		start[u] = idx
		// O[idx] = u
		idx++
		for v, w := range neigbhors[u] {
			if v == p {
				continue
			}
			dists[v] = dists[u] + w
			P[v] = u
			dfs(u, v, h+1)
		}
		end[u] = idx
	}
	dfs(-1, 0, 0)

	lca := func(u, v int) int {
		for T[u] != T[v] {
			if L[u] > L[v] {
				u = T[u]
			} else {
				v = T[v]
			}
		}

		for u != v {
			if L[u] > L[v] {
				u = P[u]
			} else {
				v = P[v]
			}
		}
		return u
	}

	st := CreateSegTree(n)

	res := make([]int, len(queries))
	var ri int

	for i := 0; i < len(queries); i++ {
		if queries[i][0] == 1 {
			j := queries[i][1] - 1
			u, v := edges[j][0]-1, edges[j][1]-1
			if P[u] == v {
				u, v = v, u
			}
			w := queries[i][2]
			a, b := start[v], end[v]-1
			st.Update(a, b, w-costs[j])
			costs[j] = w
		} else {
			u, v := queries[i][1]-1, queries[i][2]-1
			c := lca(u, v)
			du := st.Get(start[u])
			dv := st.Get(start[v])
			dc := st.Get(start[c])
			res[ri] = du + dv - 2*dc + dists[u] + dists[v] - 2*dists[c]
			ri++
		}
	}
	return res[:ri]
}

type SegTree struct {
	tree []int
	size int
}

func CreateSegTree(n int) *SegTree {
	tree := make([]int, 4*n)

	return &SegTree{tree, n}
}

func (st *SegTree) Update(start, end int, val int) {
	tree := st.tree
	size := st.size

	var loop func(i int, left, right int)
	loop = func(i int, left, right int) {
		if right < start || left > end {
			return
		}
		if start <= left && right <= end {
			tree[i] += val
			return
		}
		mid := (left + right) / 2
		loop(2*i+1, left, mid)
		loop(2*i+2, mid+1, right)
	}
	loop(0, 0, size-1)
}

func (st SegTree) Get(idx int) int {
	tree := st.tree

	var loop func(i int, left, right int) int
	loop = func(i int, left, right int) int {
		if left == right {
			return tree[i]
		}
		mid := (left + right) / 2
		if idx <= mid {
			return loop(2*i+1, left, mid) + tree[i]
		}
		return loop(2*i+2, mid+1, right) + tree[i]
	}

	return loop(0, 0, st.size-1)
}

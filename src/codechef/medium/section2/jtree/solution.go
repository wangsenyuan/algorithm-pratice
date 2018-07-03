package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	n, m := readTwoNums(scanner)

	edges := make([][]int, n-1)
	for i := 0; i < n-1; i++ {
		edges[i] = readNNums(scanner, 2)
	}
	tickets := make([][]int, m)
	for i := 0; i < m; i++ {
		tickets[i] = readNNums(scanner, 3)
	}
	Q := readNum(scanner)
	friends := make([]int, Q)

	for i := 0; i < Q; i++ {
		friends[i] = readNum(scanner)
	}

	res := solve(n, m, edges, tickets, Q, friends)
	for _, ans := range res {
		fmt.Println(ans)
	}
}

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

const INF = math.MaxInt32

func solve(n, m int, edges [][]int, tickets [][]int, Q int, friends []int) []int {
	outs := make([]map[int]bool, n)
	for i := 0; i < n; i++ {
		outs[i] = make(map[int]bool)
	}

	for _, edge := range edges {
		u, v := edge[0]-1, edge[1]-1
		outs[v][u] = true
	}

	T := make([]map[int]int, n)
	for i := 0; i < n; i++ {
		T[i] = make(map[int]int)
	}

	for _, ticket := range tickets {
		v, k, w := ticket[0]-1, ticket[1], ticket[2]
		if w1, found := T[v][k]; !found || w < w1 {
			T[v][k] = w
		}
	}

	// at most n nodes added
	st := CreateSegTree(n)

	F := make([]int, n)

	var dfs func(u int, h int)

	dfs = func(u int, h int) {
		st.Set(h, F[u])

		for v := range outs[u] {
			ans := INF

			for k, w := range T[v] {
				tmp := st.Get(max(0, h+1-k), h) + w
				ans = min(ans, tmp)
			}
			F[v] = ans
			dfs(v, h+1)
		}
		st.Set(h, INF)
	}

	dfs(0, 0)

	res := make([]int, Q)

	for i := 0; i < Q; i++ {
		res[i] = F[friends[i]-1]
	}
	return res
}

type SegTree struct {
	tree []int
	size int
}

func CreateSegTree(n int) SegTree {
	tree := make([]int, 4*n)
	for i := 0; i < len(tree); i++ {
		tree[i] = INF
	}
	return SegTree{tree: tree, size: n}
}

func (st *SegTree) Set(pos int, val int) {
	tree := st.tree

	var loop func(i int, start, end int)
	loop = func(i int, start, end int) {
		if start == end {
			tree[i] = val
			return
		}
		mid := start + (end-start)>>1
		if pos <= mid {
			loop(2*i+1, start, mid)
		} else {
			loop(2*i+2, mid+1, end)
		}
		tree[i] = min(tree[2*i+1], tree[2*i+2])
	}
	loop(0, 0, st.size-1)
}

func min(a, b int) int {
	if a <= b {
		return a
	}
	return b
}

func (st SegTree) Get(left, right int) int {
	tree := st.tree

	var loop func(i, start, end int) int

	loop = func(i, start, end int) int {
		if start > right || end < left {
			return INF
		}
		if left <= start && end <= right {
			return tree[i]
		}
		mid := start + (end-start)>>1
		a := loop(2*i+1, start, mid)
		b := loop(2*i+2, mid+1, end)
		return min(a, b)
	}

	return loop(0, 0, st.size-1)
}

func max(a, b int) int {
	if a >= b {
		return a
	}
	return b
}

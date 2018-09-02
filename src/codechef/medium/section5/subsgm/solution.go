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

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	N, M := readTwoNums(scanner)

	A := readNNums(scanner, N)

	queries := make([]int, 2*M)
	for i := 0; i < M; i++ {
		x, y := readTwoNums(scanner)
		queries[i<<1] = x
		queries[i<<1|1] = y
	}
	res := solve(N, A, M, queries)
	for _, ans := range res {
		fmt.Println(ans)
	}
}

func solve(N int, A []int, M int, queries []int) []int {
	B := make([]int, N)
	for i := 0; i < N; i++ {
		B[i] = A[i] - i
	}

	tree := NewSegTree(B)

	res := make([]int, M+1)
	res[0] = tree.GetBest()
	for i := 0; i < M; i++ {
		x, y := queries[i<<1], queries[i<<1|1]
		tree.Update(x-1, y-(x-1))
		res[i+1] = tree.GetBest()
	}
	return res
}

type Node struct {
	l, r       int
	lnum, rnum int
	lcnt, rcnt int
	best       int
}

type Tree []*Node

func NewSegTree(num []int) Tree {
	n := len(num)
	nodes := make(Tree, 4*n)

	var loop func(i int, start, end int)
	loop = func(i int, start, end int) {
		node := new(Node)
		node.l = start
		node.r = end
		nodes[i] = node

		if start == end {
			node.lnum = num[start]
			node.rnum = num[start]
			node.lcnt = 1
			node.rcnt = 1
			node.best = 1
			return
		}
		mid := (start + end) >> 1
		loop(2*i+1, start, mid)
		loop(2*i+2, mid+1, end)
		merge(nodes[i], nodes[2*i+1], nodes[2*i+2])
	}

	loop(0, 0, n-1)

	return nodes
}

func (tree Tree) GetBest() int {
	return tree[0].best
}

func (tree Tree) Update(pos int, val int) {
	n := len(tree) / 4

	var loop func(i int, start, end int)
	loop = func(i int, start, end int) {
		node := tree[i]
		if start == end {
			node.lnum = val
			node.rnum = val
			return
		}
		mid := (start + end) >> 1
		if pos <= mid {
			loop(2*i+1, start, mid)
		} else {
			loop(2*i+2, mid+1, end)
		}
		merge(tree[i], tree[2*i+1], tree[2*i+2])
	}
	loop(0, 0, n-1)
}

func merge(a, b, c *Node) {
	a.lnum = b.lnum
	if b.lnum != c.lnum || b.lcnt != b.r-b.l+1 {
		a.lcnt = b.lcnt
	} else {
		a.lcnt = b.lcnt + c.lcnt
	}

	a.rnum = c.rnum
	if c.rnum != b.rnum || c.rcnt != c.r-c.l+1 {
		a.rcnt = c.rcnt
	} else {
		a.rcnt = c.rcnt + b.rcnt
	}

	a.best = max(b.best, c.best)
	a.best = max(a.best, max(a.lcnt, a.rcnt))
	if b.rnum == c.lnum {
		a.best = max(a.best, b.rcnt+c.lcnt)
	}
}

func max(a, b int) int {
	if a >= b {
		return a
	}
	return b
}

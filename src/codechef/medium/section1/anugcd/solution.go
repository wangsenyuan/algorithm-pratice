package main

import (
	"bufio"
	"fmt"
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

	N, M := readTwoNums(scanner)
	A := readNNums(scanner, N)

	queries := make([][]int, M)

	for i := 0; i < M; i++ {
		queries[i] = readNNums(scanner, 3)
	}

	res := solve(N, A, M, queries)

	for _, ans := range res {
		fmt.Printf("%d %d\n", ans[0], ans[1])
	}
}

var primes []int

const MAX_N = 100000

var factors [][]int

func init() {
	// 9,592 primes
	primes = make([]int, 10000)
	set := make([]bool, MAX_N)
	factors = make([][]int, MAX_N+1)

	for i := 0; i <= MAX_N; i++ {
		factors[i] = make([]int, 0, 10)
	}

	var i int
	for x := 2; x < MAX_N; x++ {
		if !set[x] {
			primes[i] = x
			i++
			// factors[x] = append(factors[x], i-1)
			for y := x; y < MAX_N; y += x {
				set[y] = true
				factors[y] = append(factors[y], i-1)
			}
		}
	}

	primes = primes[:i]
}

func solve(N int, A []int, M int, queries [][]int) [][]int {
	PN := len(primes)

	nodes := make([]*Node, PN)

	for i := 0; i < N; i++ {
		a := A[i]
		for _, j := range factors[a] {
			nodes[j] = nodes[j].Insert(i, A[i], N)
		}
	}

	res := make([][]int, M)

	for i := 0; i < M; i++ {
		g, l, r := queries[i][0], queries[i][1]-1, queries[i][2]-1

		mx, mc := -1, -1

		for _, j := range factors[g] {
			a, b := nodes[j].Query(l, r, N)
			if a < 0 {
				continue
			}
			if a > mx {
				mx, mc = a, b
			} else if mx == a {
				mc += b
			}
		}

		res[i] = []int{mx, mc}
	}

	return res
}

type Node struct {
	maxVal, cnt int
	left, right *Node
}

func merge(node *Node, left, right *Node) *Node {
	if left == nil && right == nil {
		return node
	}
	if left == nil {
		node.maxVal = right.maxVal
		node.cnt = right.cnt
	} else if right == nil {
		node.maxVal = left.maxVal
		node.cnt = left.cnt
	} else if left.maxVal > right.maxVal {
		node.maxVal = left.maxVal
		node.cnt = left.cnt
	} else if right.maxVal > left.maxVal {
		node.maxVal = right.maxVal
		node.cnt = right.cnt
	} else {
		node.maxVal = left.maxVal
		node.cnt = left.cnt + right.cnt
	}

	return node
}

func (root *Node) Insert(pos int, val int, size int) *Node {
	var loop func(node *Node, start, end int) *Node
	loop = func(node *Node, start, end int) *Node {
		if node == nil {
			node = new(Node)
		}

		if start == end {
			node.maxVal = val
			node.cnt = 1
			return node
		}

		mid := (start + end) / 2
		if pos <= mid {
			node.left = loop(node.left, start, mid)
		} else {
			node.right = loop(node.right, mid+1, end)
		}

		return merge(node, node.left, node.right)
	}

	return loop(root, 0, size-1)
}

func (root *Node) Query(left, right int, size int) (int, int) {
	var loop func(node *Node, start, end int) *Node

	loop = func(node *Node, start, end int) *Node {
		if node == nil {
			return nil
		}
		if end < left || start > right {
			return nil
		}
		if left <= start && end <= right {
			return node
		}
		mid := (start + end) / 2
		a := loop(node.left, start, mid)
		b := loop(node.right, mid+1, end)
		c := new(Node)
		c.maxVal = -1
		c.cnt = -1
		return merge(c, a, b)
	}

	res := loop(root, 0, size-1)
	if res == nil {
		return -1, -1
	}
	return res.maxVal, res.cnt
}

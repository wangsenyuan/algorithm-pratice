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

	firstLine := readNNums(scanner, 4)
	N, Q, L, R := firstLine[0], firstLine[1], firstLine[2], firstLine[3]
	queries := make([][]int, Q)
	for i := 0; i < Q; i++ {
		queries[i] = readNNums(scanner, 3)
	}
	res := solve(N, Q, L, R, queries)

	for _, ans := range res {
		fmt.Println(ans)
	}
}

func solve(N, Q, L, R int, queries [][]int) []uint64 {
	A := createSegTree(N, L)
	B := createSegTree(N, R+1)

	res := make([]uint64, Q)
	var j int

	for i := 0; i < Q; i++ {
		query := queries[i]
		if query[0] == 1 {
			x, y := query[1]-1, query[2]
			A.Update(x, y)
			B.Update(x, y)
		} else {
			l, r := query[1]-1, query[2]-1
			res[j] = B.Query(l, r) - A.Query(l, r)
			j++
		}
	}

	return res[:j]
}

type Node struct {
	num_ones_subarrays                 uint64
	len, len_ones_left, len_ones_right int
	left, right                        *Node
}

func leafNode(pos int, val int) *Node {
	node := new(Node)
	node.len = 1
	node.num_ones_subarrays = uint64(val)
	node.len_ones_left = val
	node.len_ones_right = val
	return node
}

func copyNode(dst, src *Node) {
	dst.len = src.len
	dst.num_ones_subarrays = src.num_ones_subarrays
	dst.len_ones_left = src.len_ones_left
	dst.len_ones_right = src.len_ones_right
	dst.num_ones_subarrays = src.num_ones_subarrays
}

func merge(res, a, b *Node) {
	if a == nil {
		a = new(Node)
	}

	if b == nil {
		b = new(Node)
	}

	res.len = a.len + b.len
	res.len_ones_left = a.len_ones_left
	if a.len_ones_left == a.len {
		res.len_ones_left += b.len_ones_left
	}
	res.len_ones_right = b.len_ones_right
	if b.len_ones_right == b.len {
		res.len_ones_right += a.len_ones_right
	}

	res.num_ones_subarrays = a.num_ones_subarrays +
		b.num_ones_subarrays +
		uint64(a.len_ones_right)*uint64(b.len_ones_left)

}

type SegTree struct {
	maxBound int
	root     *Node
	size     int
}

func createSegTree(n int, bound int) *SegTree {
	st := new(SegTree)
	st.maxBound = bound
	st.size = n

	for i := 0; i < n; i++ {
		st.Update(i, 0)
	}

	return st
}

func (st *SegTree) Update(pos int, num int) {
	val := 0
	if num < st.maxBound {
		val = 1
	}
	var loop func(node *Node, left, right int) *Node
	loop = func(node *Node, left, right int) *Node {
		if node == nil {
			node = new(Node)
		}
		if left == right {
			node = leafNode(left, val)
			return node
		}
		mid := (left + right) / 2
		if pos <= mid {
			node.left = loop(node.left, left, mid)
		} else {
			node.right = loop(node.right, mid+1, right)
		}

		merge(node, node.left, node.right)
		return node
	}

	st.root = loop(st.root, 0, st.size-1)
}

func (st *SegTree) Query(l, r int) uint64 {
	var loop func(node *Node, start, end int) *Node
	loop = func(node *Node, start, end int) *Node {
		if end < l || r < start {
			return new(Node)
		}

		if l <= start && end <= r {
			return node
		}
		mid := (start + end) / 2
		res := new(Node)
		merge(res, loop(node.left, start, mid), loop(node.right, mid+1, end))
		return res
	}
	res := loop(st.root, 0, st.size-1)
	return res.num_ones_subarrays
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

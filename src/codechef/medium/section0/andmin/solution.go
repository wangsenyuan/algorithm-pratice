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
	var i int
	for i < n {
		x = readInt(scanner.Bytes(), x+1, &res[i])
		i++
		if x == len(scanner.Bytes()) {
			break
		}
	}
	return res[:i]
}

const INF = 1 << 30

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	n, q := readTwoNums(scanner)

	A := readNNums(scanner, n)
	queries := make([][]int, q)

	for i := 0; i < q; i++ {
		queries[i] = readNNums(scanner, 4)
	}
	res := solve(n, A, q, queries)

	for _, ans := range res {
		fmt.Println(ans)
	}
}

func solve(n int, A []int, q int, queries [][]int) []int {
	rmq := CreateTree(n, A)

	res := make([]int, q)
	var j int
	for i := 0; i < q; i++ {
		query := queries[i]
		if query[0] == 0 {
			res[j] = rmq.Get(query[1]-1, query[2]-1, n)
			j++
		} else {
			rmq.Update(query[1]-1, query[2]-1, query[3], n)
		}
	}

	return res[:j]
}

type Node struct {
	minVal      int
	bit         []int
	left, right *Node
}

func CreateTree(n int, arr []int) *Node {
	var loop func(left, right int) *Node

	loop = func(left, right int) *Node {
		if left == right {
			minVal := arr[left]
			bit := make([]int, 31)
			for i := 0; i < 31; i++ {
				bit[i] = (minVal >> uint(i)) & 1
			}
			return &Node{minVal: minVal, bit: bit}
		}

		mid := (left + right) >> 1
		a := loop(left, mid)
		b := loop(mid+1, right)
		minVal := min(a.minVal, b.minVal)
		bit := make([]int, 31)
		for i := 0; i < 31; i++ {
			bit[i] = a.bit[i] | b.bit[i]
		}
		return &Node{minVal, bit, a, b}
	}
	return loop(0, n-1)
}

func (root *Node) Update(start, end int, x int, size int) {

	var loop func(node *Node, left, right int)

	loop = func(node *Node, left, right int) {
		if end < left || right < start {
			return
		}
		flag := false
		for i := 0; i < 31; i++ {
			if x&(1<<uint(i)) == 0 && node.bit[i]&(1<<uint(i)) > 0 {
				flag = true
				break
			}
		}
		if !flag {
			return
		}
		if left == right {
			node.minVal &= x
			for i := 0; i < 31; i++ {
				node.bit[i] = (node.minVal >> uint(i)) & 1
			}
			return
		}
		mid := (left + right) >> 1
		loop(node.left, left, mid)
		loop(node.right, mid+1, right)
		node.minVal = min(node.left.minVal, node.right.minVal)
		for i := 0; i < 31; i++ {
			node.bit[i] = node.left.bit[i] | node.right.bit[i]
		}
	}

	loop(root, 0, size-1)
}

func (root Node) Get(start int, end int, size int) int {
	var loop func(node Node, left, right int) int
	loop = func(node Node, left, right int) int {
		if end < left || right < start {
			return INF
		}
		if start <= left && right <= end {
			return node.minVal
		}
		mid := (left + right) >> 1
		return min(loop(*node.left, left, mid), loop(*node.right, mid+1, right))
	}
	return loop(root, 0, size-1)
}

func solve1(n int, A []int, q int, queries [][]int) []int {
	rmq := CreateRMQ(n, A)

	res := make([]int, q)
	var j int
	for i := 0; i < q; i++ {
		query := queries[i]
		if query[0] == 0 {
			res[j] = rmq.Get(query[1]-1, query[2]-1)
			j++
		} else {
			rmq.Update(query[1]-1, query[2]-1, query[3])
		}
	}

	return res[:j]
}

type RMQ struct {
	tree []int
	bit  []int
	size int
}

func CreateRMQ(n int, arr []int) *RMQ {
	tree := make([]int, 4*n)
	bit := make([]int, 4*n)
	var loop func(i int, left, right int)

	loop = func(i int, left, right int) {
		if left == right {
			tree[i] = arr[left]
			bit[i] = arr[left]
			return
		}
		mid := (left + right) / 2
		loop(2*i+1, left, mid)
		loop(2*i+2, mid+1, right)
		tree[i] = min(tree[2*i+1], tree[2*i+2])
		bit[i] = bit[2*i+1] | bit[2*i+2]
	}

	loop(0, 0, n-1)
	return &RMQ{tree, bit, n}
}

func (this *RMQ) Update(start, end int, x int) {
	bit := this.bit
	tree := this.tree
	var loop func(i int, left, right int)
	loop = func(i int, left, right int) {
		if end < left || right < start {
			return
		}

		y := x & bit[i]

		if bit[i] == y {
			return
		}

		if left == right {
			bit[i] &= x
			tree[i] &= x
			return
		}

		mid := (left + right) / 2
		loop(2*i+1, left, mid)
		loop(2*i+2, mid+1, right)

		tree[i] = min(tree[2*i+1], tree[2*i+2])
		bit[i] = bit[2*i+1] | bit[2*i+2]
	}
	loop(0, 0, this.size-1)
}

func (this RMQ) Get(start, end int) int {
	tree := this.tree
	var loop func(i int, left, right int) int
	loop = func(i int, left, right int) int {
		if end < left || right < start {
			return INF
		}
		if start <= left && right <= end {
			return tree[i]
		}
		mid := (left + right) / 2
		a := loop(2*i+1, left, mid)
		b := loop(2*i+2, mid+1, right)
		return min(a, b)
	}
	return loop(0, 0, this.size-1)
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

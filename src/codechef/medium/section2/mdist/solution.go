package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	n := readNum(scanner)

	points := make([][]int, n)

	for i := 0; i < n; i++ {
		points[i] = readNNums(scanner, 2)
	}

	q := readNum(scanner)

	queries := make([][]int, q)

	for i := 0; i < q; i++ {
		scanner.Scan()
		bs := scanner.Bytes()
		if bs[0] == 'U' {
			queries[i] = make([]int, 4)
			x := 2
			for j := 1; j < 4; j++ {
				for x < len(bs) && bs[x] == ' ' {
					x++
				}
				x = readInt(bs, x, &queries[i][j])
			}
		} else {
			queries[i] = make([]int, 3)
			queries[i][0] = 1
			x := 2
			for j := 1; j < 3; j++ {
				for x < len(bs) && bs[x] == ' ' {
					x++
				}
				x = readInt(bs, x, &queries[i][j])
			}
		}
	}
	// fmt.Printf("[debug]%d %v %d %v\n", n, points, q, queries)
	res := solve(n, points, q, queries)

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

func solve(n int, points [][]int, q int, queries [][]int) []int64 {
	roots := make([]*Node, 4)

	for i, point := range points {
		roots[0] = roots[0].Update(i, int64(point[0])+int64(point[1]), n)
		roots[1] = roots[1].Update(i, int64(point[0])-int64(point[1]), n)
		roots[2] = roots[2].Update(i, int64(-point[0])+int64(point[1]), n)
		roots[3] = roots[3].Update(i, int64(-point[0])-int64(point[1]), n)
	}

	res := make([]int64, q)
	var j int

	for i := 0; i < q; i++ {
		query := queries[i]
		// update
		if query[0] == 0 {
			idx := query[1]
			x, y := query[2], query[3]
			// fmt.Printf("[debug] update %d %d %d\n", idx, x, y)

			roots[0] = roots[0].Update(idx, int64(x)+int64(y), n)
			roots[1] = roots[1].Update(idx, int64(x)-int64(y), n)
			roots[2] = roots[2].Update(idx, int64(-x)+int64(y), n)
			roots[3] = roots[3].Update(idx, int64(-x)-int64(y), n)
		} else {
			// query
			l, r := query[1], query[2]
			// fmt.Printf("[debug] query %d %d\n", l, r)
			res[j] = 0
			for k := 0; k < 4; k++ {
				node := roots[k].Get(l, r, n)
				// node can't be nil
				if node == nil {
					panic("node can't be nil")
				}
				tmp := node.maxVal - node.minVal
				if tmp > res[j] {
					res[j] = tmp
				}
			}
			j++
		}
	}
	return res[:j]
}

type Node struct {
	minVal, maxVal int64
	left, right    *Node
}

func (root *Node) Get(left, right int, size int) *Node {
	var loop func(node *Node, start, end int) *Node
	loop = func(node *Node, start, end int) *Node {
		if node == nil || end < left || start > right {
			return nil
		}
		if left <= start && end <= right {
			return node
		}
		mid := (start + end) >> 1
		a := loop(node.left, start, mid)
		b := loop(node.right, mid+1, end)
		if a == nil {
			return b
		}
		if b == nil {
			return a
		}
		c := new(Node)
		c.minVal = findMinVal(a, b)
		c.maxVal = findMaxVal(a, b)
		return c
	}

	return loop(root, 0, size-1)
}

func (root *Node) Update(pos int, val int64, size int) *Node {
	var loop func(node *Node, start, end int) *Node

	loop = func(node *Node, start, end int) *Node {
		if node == nil {
			node = new(Node)
			node.minVal = val
			node.maxVal = val
		}

		if start == end {
			return node
		}
		mid := (start + end) >> 1
		if pos <= mid {
			node.left = loop(node.left, start, mid)
		} else {
			node.right = loop(node.right, mid+1, end)
		}
		node.minVal = findMinVal(node.left, node.right)
		node.maxVal = findMaxVal(node.left, node.right)
		return node
	}

	return loop(root, 0, size-1)
}

func findMinVal(a, b *Node) int64 {
	if a == nil {
		return b.minVal
	}
	if b == nil {
		return a.minVal
	}
	if a.minVal <= b.minVal {
		return a.minVal
	}
	return b.minVal
}

func findMaxVal(a, b *Node) int64 {
	if a == nil {
		return b.maxVal
	}
	if b == nil {
		return a.maxVal
	}
	if a.maxVal >= b.maxVal {
		return a.maxVal
	}
	return b.maxVal
}

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

func readInt64(bytes []byte, from int, val *int64) int {
	i := from
	var tmp int64
	for i < len(bytes) && bytes[i] != ' ' {
		tmp = tmp*10 + int64(bytes[i]-'0')
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

func readNInt64Nums(scanner *bufio.Scanner, n int) []int64 {
	res := make([]int64, n)
	x := -1
	scanner.Scan()
	for i := 0; i < n; i++ {
		x = readInt64(scanner.Bytes(), x+1, &res[i])
	}
	return res
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	n, m := readTwoNums(scanner)
	A := readNNums(scanner, int(n))

	queries := make([][]int, m)
	for i := 0; i < m; i++ {
		queries[i] = readNNums(scanner, 3)
	}
	res := solve(n, A, queries)
	for _, tmp := range res {
		fmt.Println(tmp)
	}
}

func solve(n int, A []int, queries [][]int) []int64 {
	sg := ConstructSegTree(n, A)

	res := make([]int64, len(queries))
	k := 0

	for _, query := range queries {
		a, b, c := query[0], query[1], query[2]
		if a == 1 {
			sg.Update(b-1, c)
		} else {
			mk := sg.Get(b-1, c-1)
			res[k] = findPerimeter(mk)
			k++
		}
	}
	return res[:k]
}

func findPerimeter(A []int) int64 {
	for i := 0; i+3 <= len(A); i++ {
		a, b, c := A[i], A[i+1], A[i+2]
		if a-c < b {
			return int64(a) + int64(b) + int64(c)
		}
	}
	return 0
}

type Node struct {
	nums []int
	end  int
}

type SegTree struct {
	nodes []*Node
	size  int
}

func ConstructSegTree(n int, arr []int) *SegTree {
	nodes := make([]*Node, 4*n)
	for i := 0; i < len(nodes); i++ {
		nodes[i] = &Node{make([]int, 50), 0}
	}
	merge := func(i int) {
		p, a, b := nodes[i], nodes[2*i+1], nodes[2*i+2]
		var end int
		var j, k int
		for j < a.end && k < b.end && end < 50 {
			if a.nums[j] >= b.nums[k] {
				p.nums[end] = a.nums[j]
				end++
				j++
			} else {
				p.nums[end] = b.nums[k]
				end++
				k++
			}
		}

		for end < 50 && j < a.end {
			p.nums[end] = a.nums[j]
			end++
			j++
		}
		for end < 50 && k < b.end {
			p.nums[end] = b.nums[k]
			end++
			k++
		}
		p.end = end
	}

	var loop func(i int, left int, right int)
	loop = func(i int, left int, right int) {
		if left == right {
			nodes[i].nums[0] = arr[left]
			nodes[i].end = 1
			return
		}
		mid := (left + right) / 2
		loop(2*i+1, left, mid)
		loop(2*i+2, mid+1, right)
		merge(i)
	}

	loop(0, 0, n-1)

	return &SegTree{nodes, n}
}

func (sg *SegTree) Update(pos int, val int) {
	nodes := sg.nodes

	merge := func(i int) {
		p, a, b := nodes[i], nodes[2*i+1], nodes[2*i+2]
		var end int
		var j, k int
		for j < a.end && k < b.end && end < 50 {
			if a.nums[j] >= b.nums[k] {
				p.nums[end] = a.nums[j]
				end++
				j++
			} else {
				p.nums[end] = b.nums[k]
				end++
				k++
			}
		}

		for end < 50 && j < a.end {
			p.nums[end] = a.nums[j]
			end++
			j++
		}
		for end < 50 && k < b.end {
			p.nums[end] = b.nums[k]
			end++
			k++
		}
		p.end = end
	}

	var loop func(i int, left int, right int)
	loop = func(i int, left int, right int) {
		if left == right {
			nodes[i].nums[0] = val
			nodes[i].end = 1
			return
		}
		mid := (left + right) / 2
		if mid >= pos {
			loop(2*i+1, left, mid)
		} else {
			loop(2*i+2, mid+1, right)
		}
		merge(i)
	}

	loop(0, 0, sg.size-1)
}

func (sg SegTree) Get(left, right int) []int {
	nodes := sg.nodes

	merge := func(a, b []int) []int {
		res := make([]int, 50)
		end := 0
		var i, j int
		for i < len(a) && j < len(b) && end < 50 {
			if a[i] >= b[j] {
				res[end] = a[i]
				end++
				i++
			} else {
				res[end] = b[j]
				end++
				j++
			}
		}
		for i < len(a) && end < 50 {
			res[end] = a[i]
			end++
			i++
		}
		for j < len(b) && end < 50 {
			res[end] = b[j]
			end++
			j++
		}
		return res[:end]
	}

	var loop func(i int, begin int, end int) []int

	loop = func(i int, begin int, end int) []int {
		if end < left || begin > right {
			return nil
		}

		if left <= begin && end <= right {
			return nodes[i].nums[:nodes[i].end]
		}
		mid := (begin + end) / 2
		a := loop(2*i+1, begin, mid)
		b := loop(2*i+2, mid+1, end)
		if a == nil {
			return b
		}
		if b == nil {
			return a
		}
		return merge(a, b)
	}

	return loop(0, 0, sg.size-1)
}

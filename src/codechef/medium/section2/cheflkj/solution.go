package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	N, Q := readTwoNums(scanner)

	A := readNNums(scanner, N)

	queries := make([][]int, Q)

	for i := 0; i < Q; i++ {
		queries[i] = readNNums(scanner, 3)
	}

	res := solve(N, Q, A, queries)

	for _, ans := range res {
		if ans {
			fmt.Println("Yes")
		} else {
			fmt.Println("No")
		}
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

func solve(N, Q int, A []int, queries [][]int) []bool {
	idx := make(map[int]int)
	var cur int
	for i := 0; i < N; i++ {
		if _, found := idx[A[i]]; !found {
			idx[A[i]] = cur
			cur++
		}
	}
	for _, query := range queries {
		if query[0] == 1 {
			y := query[2]
			if _, found := idx[y]; !found {
				idx[y] = cur
				cur++
			}
		}
	}
	for i := 0; i < N; i++ {
		A[i] = idx[A[i]]
	}
	for i := 0; i < Q; i++ {
		if queries[i][0] == 1 {
			queries[i][2] = idx[queries[i][2]]
		}
	}
	// compact the vals
	trees := make([]*BIT, cur)
	for i := 0; i < cur; i++ {
		trees[i] = BuildBIT(N)
	}

	for i := 0; i < N; i++ {
		trees[A[i]].Update(i, 1)
	}

	res := make([]bool, Q)
	var q int
	for i := 0; i < Q; i++ {
		kind, x, y := queries[i][0], queries[i][1], queries[i][2]
		if kind == 1 {
			x--
			trees[A[x]].Update(x, -1)
			A[x] = y
			trees[A[x]].Update(x, 1)
		} else {
			x--
			y--
			var ok bool
			for i := 0; i < 25 && !ok; i++ {
				j := randBetween(x, y)
				cnt := trees[A[j]].GetRange(x, y)
				if cnt*2 > y-x+1 {
					ok = true
				}
			}
			res[q] = ok
			q++
		}
	}
	return res[:q]
}

func randBetween(left, right int) int {
	length := right - left + 1
	x := rand.Intn(length)
	return x + left
}

type BIT struct {
	arr  []int
	size int
}

func BuildBIT(n int) *BIT {
	arr := make([]int, n+1)
	size := n
	return &BIT{arr, size}
}
func (bit *BIT) Update(pos int, val int) {
	pos++
	for pos <= bit.size {
		bit.arr[pos] += val
		pos += pos & (-pos)
	}
}

func (bit BIT) GetRange(left int, right int) int {

	get := func(pos int) int {
		var res int
		pos++
		for pos > 0 {
			res += bit.arr[pos]
			pos -= pos & (-pos)
		}
		return res
	}

	if left > 0 {
		return get(right) - get(left-1)
	}
	return get(right)
}

func solve1(N, Q int, A []int, queries [][]int) []bool {
	root := BuildSegTree(A)

	res := make([]bool, Q)
	var j int

	for i := 0; i < Q; i++ {
		kind, x, y := queries[i][0], queries[i][1], queries[i][2]
		if kind == 1 {
			x--
			root = root.Update(x, y, A[x], N)
			A[x] = y
		} else {
			x--
			y--
			rgs := root.Get(x, y, N)

			for i := 0; i < len(rgs); i++ {
				var cnt int
				if rgs[i].dominant > 0 {
					for j := 0; j < len(rgs); j++ {
						cnt += rgs[j].freq[rgs[i].dominant]
					}
				}
				if cnt*2 > y-x+1 {
					res[j] = true
					break
				}
			}

			j++
		}
	}

	return res[:j]
}

type Node struct {
	dominant    int
	length      int
	freq        map[int]int
	left, right *Node
}

func BuildSegTree(A []int) *Node {

	var loop func(node *Node, start, end int) *Node
	loop = func(node *Node, start, end int) *Node {
		if node == nil {
			node = new(Node)
			node.length = end - start + 1
		}
		if start == end {
			node.dominant = A[start]
			node.freq = make(map[int]int)
			node.freq[A[start]]++
			return node
		}

		mid := (start + end) >> 1
		node.left = loop(node.left, start, mid)
		node.right = loop(node.right, mid+1, end)
		freq := make(map[int]int)
		for k, c := range node.left.freq {
			freq[k] += c
		}
		for k, c := range node.right.freq {
			freq[k] += c
		}
		node.freq = freq
		node.dominant = -1
		if node.left.dominant > 0 && freq[node.left.dominant] > node.length/2 {
			node.dominant = node.left.dominant
		}
		if node.right.dominant > 0 && freq[node.right.dominant] > node.length/2 {
			node.dominant = node.right.dominant
		}
		return node
	}

	return loop(nil, 0, len(A)-1)
}

func (root *Node) Update(pos int, val, oldVal int, size int) *Node {
	var loop func(node *Node, start, end int) *Node

	loop = func(node *Node, start, end int) *Node {
		if start == end {
			node.freq = make(map[int]int)
			node.dominant = val
			node.freq[val] = 1
			return node
		}
		mid := (start + end) >> 1
		if pos <= mid {
			node.left = loop(node.left, start, mid)
		} else {
			node.right = loop(node.right, mid+1, end)
		}
		node.freq[oldVal]--
		node.freq[val]++
		if node.left.dominant > 0 && node.freq[node.left.dominant] > node.length/2 {
			node.dominant = node.left.dominant
		} else if node.right.dominant > 0 && node.freq[node.right.dominant] > node.length/2 {
			node.dominant = node.right.dominant
		} else {
			node.dominant = -1
		}

		return node
	}

	return loop(root, 0, size-1)
}

func (root *Node) Get(left, right int, size int) []*Node {
	rgs := make([]*Node, 0, 10)

	var loop func(node *Node, start, end int)

	loop = func(node *Node, start, end int) {
		if node == nil || start > right || end < left {
			// out of range
			return
		}
		if left <= start && end <= right {
			rgs = append(rgs, node)
			return
		}

		if start < end {
			mid := (start + end) >> 1
			loop(node.left, start, mid)
			loop(node.right, mid+1, end)
		}
	}
	loop(root, 0, size-1)
	return rgs
}

package main

import (
	"sort"
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
	A := readNNums(scanner, n)
	m := readNum(scanner)
	queries := make([][]int, m)
	for i := 0; i < m; i++ {
		queries[i] = readNNums(scanner, 2)
	}
	res := solve(n, A, queries)
	for i := 0; i < m; i++ {
		fmt.Println(res[i])
	}
}

func solve(n int, A []int, queries [][]int) []int64 {
	ps := make(PS, n)

	for i := 0; i < n; i++ {
		ps[i] = Pair{i, A[i]}
	}

	sort.Sort(ps)

	sts := make([]*SegTree, n+1)
	sts[0] = nil

	for i := 0; i < n; i++ {
		sts[i+1] = Insert(sts[i], ps[i].pos, ps[i].val, 0, n-1)
	}

	findForbiddenSum := func(l, r int) int64 {
		var sum int64

		for {
			i := sort.Search(n, func(i int) bool {
				return int64(ps[i].val) > sum+1
			})
			nextSum := Query(sts[i], l, r)
			if nextSum == sum {
				return sum + 1
			}
			sum = nextSum
		}

		return -1
	}

	res := make([]int64, len(queries))

	for i := 0; i < len(queries); i++ {
		query := queries[i]
		l, r := query[0]-1, query[1]-1
		res[i] = findForbiddenSum(l, r)
	}
	return res
}

type Pair struct {
	pos int
	val int
}

type PS []Pair

func (ps PS) Len() int {
	return len(ps)
}

func (ps PS) Less(i, j int) bool {
	return ps[i].val < ps[j].val
}

func (ps PS) Swap(i, j int) {
	ps[i], ps[j] = ps[j], ps[i]
}

type SegTree struct {
	start, end  int
	sum         int64
	left, right *SegTree
}

func cloneSegTree(src *SegTree) *SegTree {
	return &SegTree{src.start, src.end, src.sum, src.left, src.right}
}

func Insert(st *SegTree, pos int, val int, start, end int) *SegTree {
	if st == nil {
		st = &SegTree{start: start, end: end}
	} else {
		st = cloneSegTree(st)
	}
	st.sum += int64(val)
	if start == end {
		return st
	}
	mid := (start + end) >> 1
	if pos <= mid {
		st.left = Insert(st.left, pos, val, start, mid)
	} else {
		st.right = Insert(st.right, pos, val, mid+1, end)
	}
	return st
}

func Query(st *SegTree, start, end int) int64 {
	if st == nil {
		return 0
	}
	if st.end < start || end < st.start {
		return 0
	}
	if start <= st.start && st.end <= end {
		return st.sum
	}
	return Query(st.left, start, end) + Query(st.right, start, end)
}

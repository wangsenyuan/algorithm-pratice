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

	fmt.Println(solve(n, A))
}

func solve(n int, A []int) uint64 {
	pos := make([]int, n+1)
	for i := 0; i < n; i++ {
		pos[A[i]] = i
	}

	sts := make([]*SegTree, n+1)
	sts[0] = create(0, n-1)

	for i := 1; i <= n; i++ {
		sts[i] = change(sts[i-1], pos[i])
	}

	qL := make([]int, 4*(n+1))
	qR := make([]int, 4*(n+1))
	var q1, q2 int
	qL[q1] = 1
	qR[q1] = n
	q1++
	var ans uint64
	for q1 != q2 {
		L, R := qL[q2], qR[q2]

		if R-L+1 >= 2 {
			ans += uint64(R - L + 1)
			k := getKth(sts[L-1], sts[R], (R-L+2)/2)
			median := A[k]
			qL[q1] = L
			qR[q1] = median - 1
			q1++
			if median+1 < R {
				qL[q1] = median + 1
				qR[q1] = R
				q1++
			}
		}

		q2++
	}
	return ans
}

type SegTree struct {
	start, end  int
	sum         int
	left, right *SegTree
}

func (st SegTree) clone() *SegTree {
	return &SegTree{st.start, st.end, st.sum, st.left, st.right}
}

func create(start, end int) *SegTree {
	if start == end {
		return &SegTree{start, end, 0, nil, nil}
	}
	mid := (start + end) / 2
	return &SegTree{start, end, 0, create(start, mid), create(mid+1, end)}
}

func change(st *SegTree, p int) *SegTree {
	st = st.clone()
	st.sum++

	if st.start < st.end {
		if st.left.end >= p {
			st.left = change(st.left, p)
		} else {
			st.right = change(st.right, p)
		}
	}
	return st
}

func getKth(low, high *SegTree, key int) int {
	if low.start == low.end {
		return low.start
	}
	if high.left.sum-low.left.sum >= key {
		return getKth(low.left, high.left, key)
	}
	return getKth(low.right, high.right, key-(high.left.sum-low.left.sum))
}

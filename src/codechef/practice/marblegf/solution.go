package main

import "fmt"

func main() {
	var n, q int
	fmt.Scanf("%d %d", &n, &q)

	marbles := make([]int, n)

	for i := 0; i < n; i++ {
		fmt.Scanf("%d", &marbles[i])
	}

	st := Construct(marbles)

	var tp string
	var a, b int
	for i := 0; i < q; i++ {
		fmt.Scanf("%s %d %d", &tp, &a, &b)
		if tp == "S" {
			fmt.Println(st.sum(a, b))
		} else if tp == "G" {
			st.add(a, b)
		} else {
			st.sub(a, b)
		}
	}
}

type SegTree struct {
	n    int
	vals []int
	sums []int
}

func Construct(vals []int) *SegTree {
	n := len(vals)
	m := 2*n + 2
	sums := make([]int, m)
	st := &SegTree{len(vals), vals, sums}
	for i := 0; i < n; i++ {
		st.add(i, vals[i])
	}

	return st
}

func (st *SegTree) add(i, diff int) {
	st.update(0, st.n, 1, i, diff)
	st.vals[i] += diff
}

func (st *SegTree) sub(i, diff int) {
	st.add(i, -diff)
}

func (st *SegTree) update(i, j, k, m, diff int) {
	if i > m || j < m || i == j {
		return
	}
	st.sums[k] += diff
	st.update(i, i+(j-i)/2, 2*k, m, diff)
	st.update(i+(j-i)/2+1, j, 2*k+1, m, diff)
}

func (st *SegTree) sumRange(i, j, k, a, b int) int {
	if j < a || i > b || i == j {
		return 0
	}

	if a <= i && j <= b {
		return st.sums[k]
	}
	return st.sumRange(i, i+(j-i)/2, 2*k, a, b) + st.sumRange(i+(j-i)/2+1, 2*k+1, 2*k+1, a, b)
}

func (st *SegTree) sum(a, b int) int {
	return st.sumRange(0, st.n, 1, a, b)
}

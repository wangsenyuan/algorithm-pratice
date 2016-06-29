package main

import (
	"fmt"
	"math"
	"sort"
)

func main() {
	matrix := [][]int{
		[]int{2, 2, -1},
	}

	res := maxSumSubmatrix(matrix, 0)

	fmt.Printf("%d\n", res)
}

func maxSumSubmatrix(matrix [][]int, target int) int {
	m := len(matrix)
	if m == 0 {
		return 0
	}
	n := len(matrix[0])
	row := false
	if m > n {
		m, n = n, m
		row = true
	}

	res := math.MinInt32
	for i := 0; i < m; i++ {
		cols := make([]int, n, n)
		for j := i; j >= 0; j-- {
			vals := make([]int, 1, m)
			vals[0] = 0
			sum := 0
			for k := 0; k < n; k++ {
				if row {
					cols[k] += matrix[k][j]
				} else {
					cols[k] += matrix[j][k]
				}

				sum += cols[k]
				x := findCelling(vals, sum-target)
				if x >= 0 && x < len(vals) && res < sum-vals[x] {
					res = sum - vals[x]
				}
				vals = insertAndSort(vals, sum)
			}
		}
	}

	return res
}

func findCelling(vals []int, target int) int {
	return sort.Search(len(vals), func(i int) bool {
		return vals[i] >= target
	})
}

func insertAndSort(vals []int, target int) []int {
	if len(vals)+1 == cap(vals) {
		tmp := make([]int, len(vals), 2*cap(vals))
		copy(tmp, vals)
		vals = tmp
	}

	x := sort.Search(len(vals), func(i int) bool {
		return vals[i] > target
	})

	if x < 0 || x >= len(vals) {
		return append(vals, target)
	}

	a := vals[:x]
	b := vals[x:]

	c := make([]int, x, x+1)
	copy(c, a)
	c = append(c, target)
	return append(c, b...)
}

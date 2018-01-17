package main

import "sort"

func main() {

}

func solve(n, m int, A, B [][]int) int64 {
	AB := make([][]int, n)
	for i := 0; i < n; i++ {
		AB[i] = make([]int, m)
		for j := 0; j < m; j++ {
			AB[i][j] = A[i][j] - B[i][j]
		}
	}

	for i := 1; i < n; i++ {
		for j := 1; j < m; j++ {
			if AB[i][j]-AB[0][j] != AB[i][0]-AB[0][0] {
				return -1
			}
		}
	}

	ps := make([]int, 0, n+m)
	for i := 0; i < n; i++ {
		ps = append(ps, AB[0][0]-AB[i][0])
	}

	for i := 0; i < m; i++ {
		ps = append(ps, AB[0][i])
	}

	sort.Ints(ps)
	mid := (n + m) / 2
	var ans int64

	for _, x := range ps {
		ans += abs(int64(ps[mid] - x))
	}

	return ans
}

func abs(a int64) int64 {
	if a < 0 {
		return -a
	}
	return a
}

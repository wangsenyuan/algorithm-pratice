package main

import (
	"testing"
)

func runSample(t *testing.T, n int, m int, X []int, A []string, expect []int) {
	res := solve(n, m, X, A)

	a := suprise(X, A, expect)
	b := suprise(X, A, res)

	if a != b {
		t.Errorf("Sample expect %v => %d , but got %v => %d", expect, a, res, b)
	}
}

func suprise(X []int, A []string, P []int) int {
	n := len(X)
	m := len(A[0])
	var res int
	for i := 0; i < n; i++ {
		var tmp int
		for j := 0; j < m; j++ {
			if A[i][j] == '1' {
				tmp += P[j]
			}
		}
		res += abs(tmp - X[i])
	}
	return res
}

func abs(num int) int {
	if num < 0 {
		return -num
	}
	return num
}

func TestSample1(t *testing.T) {
	n, m := 4, 3
	X := []int{5, 1, 2, 2}
	A := []string{
		"110",
		"100",
		"101",
		"100",
	}
	expect := []int{3, 1, 2}
	runSample(t, n, m, X, A, expect)
}

func TestSample2(t *testing.T) {
	n, m := 4, 4
	X := []int{6, 2, 0, 10}
	A := []string{
		"1001",
		"0010",
		"0110",
		"0101",
	}
	expect := []int{2, 3, 4, 1}
	runSample(t, n, m, X, A, expect)
}

func TestSample3(t *testing.T) {
	n, m := 3, 6
	X := []int{20, 3, 15}
	A := []string{
		"010110",
		"000101",
		"111111",
	}
	expect := []int{3, 1, 4, 5, 2, 6}
	runSample(t, n, m, X, A, expect)
}

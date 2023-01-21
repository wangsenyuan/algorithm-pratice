package main

import "testing"

func runSample(t *testing.T, n int, m int, L, R []int) {
	p, q := solve(n, m, L, R)

	for i := 0; i < m; i++ {
		l, r := L[i], R[i]
		l--
		r--
		if sign(p[l]-p[r]) != sign(q[l]-q[r]) {
			t.Fatalf("Sample result %v %v, not correct at %d", p, q, i)
		}
	}
}

func sign(num int) int {
	if num < 0 {
		return -1
	}
	if num == 0 {
		return 0
	}
	return 1
}

func TestSample1(t *testing.T) {
	n, m := 4, 2
	L := []int{1, 3}
	R := []int{2, 4}
	runSample(t, n, m, L, R)
}

func TestSample2(t *testing.T) {
	n, m := 6, 4
	L := []int{1, 1, 3, 3}
	R := []int{2, 3, 5, 6}
	runSample(t, n, m, L, R)
}

func TestSample3(t *testing.T) {
	n, m := 2, 1
	L := []int{1}
	R := []int{2}
	runSample(t, n, m, L, R)
}

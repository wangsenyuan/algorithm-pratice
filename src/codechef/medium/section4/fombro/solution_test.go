package main

import "testing"

func runSample(t *testing.T, n, m int, R []int, expect []int64) {
	solver := NewSolver(n, m)

	for i := 0; i < len(R); i++ {
		res := solver.Query(R[i])

		if res != expect[i] {
			t.Fatalf("%d %d, expect %v, but got %d at %d", n, m, expect, res, i)
		}
	}

}

func TestSample1(t *testing.T) {
	n, m := 5, 114
	R := []int{2}
	expect := []int64{72}
	runSample(t, n, m, R, expect)
}

func TestSample2(t *testing.T) {
	n, m := 50, 3874
	R := []int{31, 17, 21}
	expect := []int64{3718, 624, 1144}
	runSample(t, n, m, R, expect)
}

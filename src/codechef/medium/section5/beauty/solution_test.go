package main

import "testing"

func runSample(t *testing.T, n, k int, s string, Q []string, expect []int64) {
	solver := solve(n, k, s)

	for i := 0; i < len(Q); i++ {
		ans := solver.Query(len(Q[i]), Q[i])
		if ans != expect[i] {
			t.Fatalf("Sample expect %v, but got %d at %d", expect, ans, i)
		}
	}
}

func TestSample1(t *testing.T) {
	n, k := 7, 4
	S := "acabada"
	Q := []string{"b", "ac"}
	expect := []int64{8, 2}
	runSample(t, n, k, S, Q, expect)
}

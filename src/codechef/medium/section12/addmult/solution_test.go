package main

import "testing"

func runSample(t *testing.T, n int, st string, A []int, expect string) {
	res := solve(n, st, A)

	if res != expect {
		t.Errorf("Sample %d %s %v, expect %s, but got %s", n, st, A, expect, res)
	}
}

func TestSample1(t *testing.T) {
	n := 3
	st := "Chef"
	A := []int{1, 2, 3}
	expect := "Chefu"
	runSample(t, n, st, A, expect)
}

package main

import "testing"

func runSample(t *testing.T, n int, A []int64, X int64, notSad int, happy int) {
	a, b := solve(n, A, X)

	if a != notSad || b != happy {
		t.Errorf("sample %d %v %d, expect %d %d, but got %d %d", n, A, X, notSad, happy, a, b)
	}
}

func TestSample1(t *testing.T) {
	n := 3
	A := []int64{3, 4, 5}
	var X int64 = 10
	notSad, happy := 0, 2
	runSample(t, n, A, X, notSad, happy)
}

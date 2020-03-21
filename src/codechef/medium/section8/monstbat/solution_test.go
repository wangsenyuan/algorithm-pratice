package main

import "testing"

func runSample(t *testing.T, N int, V []int, D string, M int, U []int, E string, expect int64) {
	res := solve(N, V, D, M, U, E)

	if res != expect {
		t.Errorf("Sample %d %v %s %d %v %s, expect %d, but got %d", N, V, D, M, U, E, expect, res)
	}
}

func TestSample1(t *testing.T) {
	N := 3
	V := []int{5, 4, 60}
	D := "AAA"
	M := 3
	U := []int{2, 15, 16}
	E := "DAA"
	runSample(t, N, V, D, M, U, E, 38)
}

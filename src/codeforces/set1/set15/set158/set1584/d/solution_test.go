package main

import "testing"

func runSample(t *testing.T, n int, i int, j int, k int) {
	var cnt int
	ask := func(l int, r int) int {
		cnt++
		if cnt > 40 {
			t.Fatalf("Sample took too much queries")
		}
		var res int
		a, b := max(l, i), min(r, j-1)
		res += max(0, b-a) * (b - a + 1) / 2
		c, d := max(l, j), min(r, k)
		res += max(0, d-c) * (d - c + 1) / 2
		return res
	}

	res := solve(n, ask)

	if res[0] != i || res[1] != j || res[2] != k {
		t.Fatalf("Sample expect %d %d %d, but got %v", i, j, k, res)
	}
}

func TestSample1(t *testing.T) {
	n, i, j, k := 5, 1, 3, 5
	// [2,1,5,4,3]
	runSample(t, n, i, j, k)
}

func TestSample2(t *testing.T) {
	n, i, j, k := 5, 2, 4, 5
	runSample(t, n, i, j, k)
}

func TestSample3(t *testing.T) {
	n, i, j, k := 999999953, 999999941, 999999952, 999999953
	runSample(t, n, i, j, k)
}

func TestSample4(t *testing.T) {
	n, i, j, k := 1000000000, 999999997, 999999999, 1000000000
	runSample(t, n, i, j, k)
}

func TestSample5(t *testing.T) {
	n, i, j, k := 999999954, 1, 499999977, 999999944
	runSample(t, n, i, j, k)
}

func TestSample6(t *testing.T) {
	n, i, j, k := 999999964, 499999982, 999999953, 999999964
	runSample(t, n, i, j, k)
}

func TestSample7(t *testing.T) {
	n, i, j, k := 999999980, 1, 228, 499999990
	runSample(t, n, i, j, k)
}

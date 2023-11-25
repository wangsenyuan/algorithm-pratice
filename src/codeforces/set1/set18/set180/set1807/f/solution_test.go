package main

import "testing"

func runSample(t *testing.T, n int, m int, i1, j1 int, i2, j2 int, d string, expect int) {
	res := solve(n, m, i1, j1, i2, j2, d)

	if res != expect {
		t.Fatalf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	n, m, i1, j1, i2, j2, d := 5, 7, 1, 7, 2, 4, "DL"
	expect := 3
	runSample(t, n, m, i1, j1, i2, j2, d, expect)
}

func TestSample2(t *testing.T) {
	n, m, i1, j1, i2, j2, d := 5, 7, 1, 7, 3, 2, "DL"
	expect := -1
	runSample(t, n, m, i1, j1, i2, j2, d, expect)
}

func TestSample3(t *testing.T) {
	n, m, i1, j1, i2, j2, d := 3, 3, 1, 3, 2, 2, "UR"
	expect := 1
	runSample(t, n, m, i1, j1, i2, j2, d, expect)
}

func TestSample4(t *testing.T) {
	n, m, i1, j1, i2, j2, d := 2, 4, 2, 1, 2, 2, "DR"
	expect := -1
	runSample(t, n, m, i1, j1, i2, j2, d, expect)
}

func TestSample5(t *testing.T) {
	n, m, i1, j1, i2, j2, d := 4, 3, 1, 1, 1, 3, "UL"
	expect := 4
	runSample(t, n, m, i1, j1, i2, j2, d, expect)
}

func TestSample6(t *testing.T) {
	n, m, i1, j1, i2, j2, d := 6, 4, 1, 2, 3, 4, "DR"
	expect := 0
	runSample(t, n, m, i1, j1, i2, j2, d, expect)
}

package main

import "testing"

func runSample(t *testing.T, n, m int, Z, L, R, B int, expect int) {
	res := solve(n, m, Z, L, R, B)
	if res != expect {
		t.Errorf("sample %d %d %d %d %d %d, expect %d, but got %d", n, m, Z, L, R, B, expect, res)
	}
}

func TestSample1(t *testing.T) {
	n, m, Z, L, R, B := 2, 2, 3, 2, 1, 1
	runSample(t, n, m, Z, L, R, B, 4)
}

func TestSample2(t *testing.T) {
	n, m, Z, L, R, B := 3, 3, 1, 2, 0, 9
	runSample(t, n, m, Z, L, R, B, 8)
}

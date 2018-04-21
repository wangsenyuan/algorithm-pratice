package main

import "testing"

func runSample(t *testing.T, n, x, y int, A []int, B []int, expect int) {
	res := solve(n, x, y, A, B)
	if res != expect {
		t.Errorf("sample %d %d %d %v %v, expect %d, but got %d", n, x, y, A, B, expect, res)
	}
}

func TestSample1(t *testing.T) {
	n, x, y := 5, 3, 3
	A := []int{1, 2, 3, 4, 5}
	B := []int{5, 4, 3, 2, 1}
	expect := 21
	runSample(t, n, x, y, A, B, expect)
}

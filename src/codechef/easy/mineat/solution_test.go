package main

import "testing"

func runSample(t *testing.T, n int, h int, A []int, expect int) {
	res := solve(n, h, A)
	if res != expect {
		t.Errorf("sample %d %d %v, expect %d, but got %d", n, h, A, expect, res)
	}
}

func TestSample1(t *testing.T) {
	n, h := 3, 3
	A := []int{1, 2, 3}
	runSample(t, n, h, A, 3)
}

func TestSample2(t *testing.T) {
	n, h := 3, 4
	A := []int{1, 2, 3}
	runSample(t, n, h, A, 2)
}

func TestSample3(t *testing.T) {
	n, h := 4, 5
	A := []int{4, 3, 2, 7}
	runSample(t, n, h, A, 4)
}

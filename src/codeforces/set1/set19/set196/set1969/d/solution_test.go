package main

import "testing"

func runSample(t *testing.T, k int, a []int, b []int, expect int) {
	res := solve(k, a, b)

	if res != expect {
		t.Fatalf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	k := 0
	a := []int{2, 1}
	b := []int{1, 2}
	expect := 1
	runSample(t, k, a, b, expect)
}

func TestSample2(t *testing.T) {
	k := 1
	a := []int{1, 2, 1, 4}
	b := []int{3, 3, 2, 3}
	expect := 1
	runSample(t, k, a, b, expect)
}

func TestSample3(t *testing.T) {
	k := 2
	a := []int{2, 1, 1, 1}
	b := []int{4, 2, 3, 2}
	expect := 0
	runSample(t, k, a, b, expect)
}

func TestSample4(t *testing.T) {
	k := 2
	a := []int{1, 3, 4, 9, 1, 3}
	b := []int{7, 6, 8, 10, 6, 8}
	// 不选择[9, 10]的时候，可以获得最大的收益7
	expect := 7
	runSample(t, k, a, b, expect)
}

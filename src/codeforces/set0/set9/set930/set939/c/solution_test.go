package main

import "testing"

func runSample(t *testing.T, s int, f int, a []int, expect int) {
	res := solve(a, s, f)

	if res != expect {
		t.Fatalf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	a := []int{1, 2, 3}
	// 在0的时候，1的时间是 0
	s, f := 1, 3
	expect := 3
	runSample(t, s, f, a, expect)
}

func TestSample2(t *testing.T) {
	a := []int{1, 2, 3, 4, 1}
	// 在0的时候，1的时间是 0
	s, f := 1, 3
	expect := 4
	runSample(t, s, f, a, expect)
}

func TestSample3(t *testing.T) {
	a := []int{7171, 2280, 6982, 9126, 9490, 2598, 569, 6744, 5754, 1855}
	// 在0的时候，1的时间是 0
	s, f := 7, 9
	expect := 4
	runSample(t, s, f, a, expect)
}

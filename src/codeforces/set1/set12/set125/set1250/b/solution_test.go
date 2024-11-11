package main

import "testing"

func runSample(t *testing.T, a []int, k int, expect int) {
	res := solve(a, k)

	if res != expect {
		t.Fatalf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	a := []int{3, 1, 2, 3, 2, 3}
	k := 3
	expect := 6
	runSample(t, a, k, expect)
}

func TestSample2(t *testing.T) {
	a := []int{1, 1, 1, 1, 1, 1, 1, 1, 1, 1}
	k := 1
	expect := 10
	runSample(t, a, k, expect)
}

func TestSample3(t *testing.T) {
	a := []int{1, 2, 3, 1, 2, 3, 4, 1, 2, 1, 2, 1}
	k := 4
	expect := 12
	runSample(t, a, k, expect)
}

// 3 1 2 2 3
func TestSample4(t *testing.T) {
	a := []int{3, 1, 2, 2, 3}
	k := 3
	expect := 6
	runSample(t, a, k, expect)
}

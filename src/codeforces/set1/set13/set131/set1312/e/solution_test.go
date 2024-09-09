package main

import "testing"

func runSample(t *testing.T, a []int, expect int) {
	res := solve(a)

	if res != expect {
		t.Fatalf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	a := []int{4, 3, 2, 2, 3}
	expect := 2
	runSample(t, a, expect)
}

func TestSample2(t *testing.T) {
	a := []int{3, 3, 4, 4, 4, 3, 3}
	expect := 2
	runSample(t, a, expect)
}

func TestSample3(t *testing.T) {
	a := []int{1, 3, 5}
	expect := 3
	runSample(t, a, expect)
}

func TestSample4(t *testing.T) {
	a := []int{67, 67, 65, 65, 66, 66, 66, 68, 67, 67, 67, 65, 65, 66, 70}
	// 67, 67, 67, 67, 68
	expect := 2
	runSample(t, a, expect)
}

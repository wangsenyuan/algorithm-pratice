package main

import "testing"

func runSample(t *testing.T, l []int, c []int, expect int) {
	res := solve(l, c)

	if res != expect {
		t.Fatalf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	l := []int{100, 99, 9900}
	c := []int{1, 1, 1}
	expect := 2
	runSample(t, l, c, expect)
}

func TestSample2(t *testing.T) {
	l := []int{10, 20, 30, 40, 50}
	c := []int{1, 1, 1, 1, 1}
	expect := -1
	runSample(t, l, c, expect)
}

func TestSample3(t *testing.T) {
	l := []int{15015, 10010, 6006, 4290, 2730, 2310, 1}
	c := []int{1, 1, 1, 1, 1, 1, 10}
	expect := 6
	runSample(t, l, c, expect)
}

func TestSample4(t *testing.T) {
	l := []int{4264, 4921, 6321, 6984, 2316, 8432, 6120, 1026}
	c := []int{4264, 4921, 6321, 6984, 2316, 8432, 6120, 1026}
	expect := 7237
	runSample(t, l, c, expect)
}

package main

import "testing"

func runSample(t *testing.T, p int, x []int, expect int) {
	res := solve(p, x)

	if res != expect {
		t.Fatalf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	p := 3
	x := []int{1, 2}
	expect := 1
	runSample(t, p, x, expect)
}

func TestSample2(t *testing.T) {
	p := 1000
	x := []int{4, 1, 3, 2, 5}
	expect := 995
	runSample(t, p, x, expect)
}

func TestSample3(t *testing.T) {
	p := 3
	x := []int{2}
	expect := 1
	runSample(t, p, x, expect)
}

func TestSample4(t *testing.T) {
	p := 4
	x := []int{3, 0, 2}
	expect := 2
	runSample(t, p, x, expect)
}

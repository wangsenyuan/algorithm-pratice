package main

import "testing"

func runSample(t *testing.T, p int, k int, a []int, expect int) {
	res := solve(p, k, a)

	if res != expect {
		t.Fatalf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	p, k := 6, 2
	a := []int{2, 4, 3, 5, 7}
	expect := 3
	runSample(t, p, k, a, expect)
}

func TestSample2(t *testing.T) {
	p, k := 11, 2
	a := []int{2, 4, 3, 5, 7}
	expect := 4
	runSample(t, p, k, a, expect)
}

func TestSample3(t *testing.T) {
	p, k := 2, 3
	a := []int{4, 2, 6}
	expect := 1
	runSample(t, p, k, a, expect)
}

func TestSample4(t *testing.T) {
	p, k := 2, 3
	a := []int{10, 1, 3, 9, 2}
	expect := 1
	runSample(t, p, k, a, expect)
}

func TestSample5(t *testing.T) {
	p, k := 6, 4
	a := []int{3, 2, 3, 2}
	expect := 4
	runSample(t, p, k, a, expect)
}

func TestSample6(t *testing.T) {
	p, k := 5, 3
	a := []int{1, 2, 2, 1, 2}
	expect := 5
	runSample(t, p, k, a, expect)
}

package main

import "testing"

func runSample(t *testing.T, p int, a []int, expect int) {
	res := solve(p, a)

	if res != expect {
		t.Fatalf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	a := []int{10, 10}
	p := 5
	expect := 5
	runSample(t, p, a, expect)
}

func TestSample2(t *testing.T) {
	a := []int{3, 3}
	p := 5
	expect := 3
	runSample(t, p, a, expect)
}

func TestSample3(t *testing.T) {
	a := []int{5, 2, 4, 9}
	p := 5
	expect := 12
	runSample(t, p, a, expect)
}

func TestSample4(t *testing.T) {
	a := []int{5, 3, 3, 6, 10, 100, 9, 15}
	p := 8
	expect := 46
	runSample(t, p, a, expect)
}

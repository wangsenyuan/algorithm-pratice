package main

import "testing"

func runSample(t *testing.T, m int, b int, mod int, a []int, expect int) {
	res := solve(m, b, mod, a)

	if res != expect {
		t.Fatalf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	m, b, mod := 3, 3, 100
	a := []int{1, 1, 1}
	expect := 10
	runSample(t, m, b, mod, a, expect)
}

func TestSample2(t *testing.T) {
	m, b, mod := 6, 5, 1000000007
	a := []int{1, 2, 3}
	expect := 0
	runSample(t, m, b, mod, a, expect)
}

func TestSample3(t *testing.T) {
	m, b, mod := 6, 5, 11
	a := []int{1, 2, 1}
	expect := 0
	runSample(t, m, b, mod, a, expect)
}

package main

import "testing"

func runSample(t *testing.T, m int, a []int, c []int, expect int) {
	res := solve(m, a, c)

	if res != expect {
		t.Fatalf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	m := 10
	a := []int{1, 2, 3}
	c := []int{2, 2, 1}
	expect := 7
	runSample(t, m, a, c, expect)
}

func TestSample2(t *testing.T) {
	m := 20
	a := []int{4, 2, 7, 5, 6, 1}
	c := []int{1, 2, 1, 3, 1, 7}
	expect := 19
	runSample(t, m, a, c, expect)
}

func TestSample3(t *testing.T) {
	m := 1033
	a := []int{206, 207, 1000}
	c := []int{3, 4, 1}
	expect := 1033
	runSample(t, m, a, c, expect)
}

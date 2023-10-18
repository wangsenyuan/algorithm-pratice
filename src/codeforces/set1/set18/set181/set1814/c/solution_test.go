package main

import "testing"

func runSample(t *testing.T, s1, s2 int, r []int, a, b []int) {
	c, d := solve(s1, s2, r)

	expect := check(s1, s2, r, a, b)
	res := check(s1, s2, r, c, d)

	if expect != res {
		t.Fatalf("Sample result %v, %v, not correct", c, d)
	}
}

func check(s1, s2 int, r []int, a, b []int) int {
	var res int

	for i := 0; i < len(r); i++ {
		res += r[i] * find(s1, s2, a, b, i+1)
	}

	return res
}

func find(s1 int, s2 int, a, b []int, id int) int {
	for i := 0; i < len(a); i++ {
		if a[i] == id {
			return s1 * (i + 1)
		}
	}

	for i := 0; i < len(b); i++ {
		if b[i] == id {
			return s2 * (i + 1)
		}
	}

	return -1
}

func TestSample1(t *testing.T) {
	s1, s2 := 3, 1
	r := []int{8, 6, 4, 4, 4, 1, 7}
	a := []int{5, 6}
	b := []int{1, 7, 2, 4, 3}
	runSample(t, s1, s2, r, a, b)
}

func TestSample2(t *testing.T) {
	s1, s2 := 1, 10
	r := []int{1, 1, 1, 1, 1}
	a := []int{4, 3, 5, 2, 1}
	b := []int{}
	runSample(t, s1, s2, r, a, b)
}

func TestSample3(t *testing.T) {
	s1, s2 := 1, 1
	r := []int{4, 5, 6, 8, 1, 7, 3, 2}
	a := []int{4, 2, 7, 5}
	b := []int{6, 3, 1, 8}
	runSample(t, s1, s2, r, a, b)
}

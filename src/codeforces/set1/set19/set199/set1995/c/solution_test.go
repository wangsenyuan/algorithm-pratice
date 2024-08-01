package main

import "testing"

func runSample(t *testing.T, a []int, expect int) {
	res := solve(a)

	if expect != res {
		t.Fatalf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	a := []int{1, 2, 3}
	expect := 0
	runSample(t, a, expect)
}

func TestSample2(t *testing.T) {
	a := []int{3, 2}
	expect := 1
	runSample(t, a, expect)
}

func TestSample3(t *testing.T) {
	a := []int{3, 1, 5}
	expect := -1
	runSample(t, a, expect)
}

func TestSample4(t *testing.T) {
	a := []int{1, 1, 2, 3}
	expect := 0
	runSample(t, a, expect)
}

func TestSample5(t *testing.T) {
	a := []int{4, 3, 2}
	expect := 3
	runSample(t, a, expect)
}

func TestSample6(t *testing.T) {
	a := []int{16, 2, 4, 2, 256, 2, 4, 2, 8}
	expect := 15
	runSample(t, a, expect)
}

func TestSample7(t *testing.T) {
	a := []int{10010, 10009, 10008, 10007, 10006, 10005, 10004, 10003, 10002, 10001, 10000}
	expect := 55
	runSample(t, a, expect)
}

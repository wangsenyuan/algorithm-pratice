package main

import "testing"

func runSample(t *testing.T, n int, k int, l int, d []int, expect bool) {
	res := solve(n, k, l, d)
	if res != expect {
		t.Fatalf("Sample expect %t, but got %t", expect, res)
	}
}

func TestSample1(t *testing.T) {
	n, k, l := 2, 1, 1
	d := []int{1, 0}
	expect := true
	runSample(t, n, k, l, d, expect)
}

func TestSample2(t *testing.T) {
	n, k, l := 5, 2, 3
	d := []int{1, 2, 3, 2, 2}
	expect := false
	runSample(t, n, k, l, d, expect)
}

func TestSample3(t *testing.T) {
	n, k, l := 4, 3, 4
	d := []int{0, 2, 4, 3}
	expect := true
	runSample(t, n, k, l, d, expect)
}

func TestSample4(t *testing.T) {
	n, k, l := 2, 3, 5
	d := []int{3, 0}
	expect := true
	runSample(t, n, k, l, d, expect)
}

func TestSample5(t *testing.T) {
	n, k, l := 7, 2, 3
	d := []int{3, 0, 2, 1, 3, 0, 1}
	expect := true
	runSample(t, n, k, l, d, expect)
}

func TestSample6(t *testing.T) {
	n, k, l := 7, 1, 4
	d := []int{4, 4, 3, 0, 2, 4, 2}
	expect := false
	runSample(t, n, k, l, d, expect)
}

func TestSample7(t *testing.T) {
	n, k, l := 5, 2, 3
	d := []int{1, 2, 3, 2, 2}
	expect := false
	runSample(t, n, k, l, d, expect)
}

package main

import "testing"

func runSample(t *testing.T, n int, m int, x []int, expect int) {
	res := solve(n, m, x)

	if res != expect {
		t.Fatalf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	n, m := 3, 10
	x := []int{5, 5, 5}
	expect := 1
	runSample(t, n, m, x, expect)
}

func TestSample2(t *testing.T) {
	n, m := 4, 6
	x := []int{1, -2, -2, 1}
	expect := 3
	runSample(t, n, m, x, expect)
}

func TestSample3(t *testing.T) {
	n, m := 5, 7
	x := []int{-1, -1, 4, -2, -2}
	expect := 5
	runSample(t, n, m, x, expect)
}

func TestSample4(t *testing.T) {
	n, m := 6, 7
	x := []int{5, -1, -2, -2, -2, -2}
	expect := 5
	runSample(t, n, m, x, expect)
}

func TestSample5(t *testing.T) {
	n, m := 5, 2
	x := []int{1, -1, -2, -2, -1}
	expect := 2
	runSample(t, n, m, x, expect)
}

func TestSample6(t *testing.T) {
	n, m := 5, 4
	x := []int{2, 3, -2, 1, -2}
	expect := 4
	runSample(t, n, m, x, expect)
}

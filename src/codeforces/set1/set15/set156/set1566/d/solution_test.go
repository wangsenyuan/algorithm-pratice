package main

import "testing"

func runSample(t *testing.T, n int, m int, a []int, expect int) {
	res := solve(n, m, a)

	if res != expect {
		t.Fatalf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	a := []int{1, 2}
	n, m := 1, 2
	expect := 1
	runSample(t, n, m, a, expect)
}

func TestSample2(t *testing.T) {
	a := []int{1, 1, 2, 2, 3, 3}
	n, m := 3, 2
	expect := 0
	runSample(t, n, m, a, expect)
}

func TestSample3(t *testing.T) {
	a := []int{3, 4, 4, 1, 1, 1, 1, 1, 2}
	n, m := 3, 3
	expect := 4
	runSample(t, n, m, a, expect)
}

func TestSample4(t *testing.T) {
	a := []int{1, 1, 2, 1}
	n, m := 2, 2
	expect := 0
	runSample(t, n, m, a, expect)
}

func TestSample5(t *testing.T) {
	a := []int{50, 50, 50, 50, 3, 50, 50, 50}
	n, m := 4, 2
	expect := 0
	runSample(t, n, m, a, expect)
}

func TestSample6(t *testing.T) {
	a := []int{6, 6, 6, 6, 2, 2, 9, 6}
	n, m := 4, 2
	expect := 0
	runSample(t, n, m, a, expect)
}

func TestSample7(t *testing.T) {
	a := []int{1, 3, 3, 3, 3, 3, 1, 1, 3, 1, 3, 1, 1, 3, 3, 1, 1, 3}
	n, m := 2, 9
	expect := 1
	runSample(t, n, m, a, expect)
}

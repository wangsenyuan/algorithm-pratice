package main

import "testing"

func runSample(t *testing.T, a []int, b []int, k int, expect int) {
	res := solve(a, b, k)

	if res != expect {
		t.Fatalf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	a := []int{3, 3}
	b := []int{1, 1}
	k := 10
	expect := 16
	runSample(t, a, b, k, expect)
}

func TestSample2(t *testing.T) {
	a := []int{3, 3, 3}
	b := []int{0, 0, 0}
	k := 10
	expect := 6
	runSample(t, a, b, k, expect)
}

func TestSample3(t *testing.T) {
	a := []int{2, 1, 5, 1}
	b := []int{0, 1, 0, 1}
	k := 4
	expect := 8
	runSample(t, a, b, k, expect)
}

func TestSample4(t *testing.T) {
	a := []int{7, 5, 2, 5, 4}
	b := []int{0, 0, 1, 0, 1}
	k := 4
	expect := 13
	runSample(t, a, b, k, expect)
}

func TestSample5(t *testing.T) {
	a := []int{5, 15, 15, 2, 11}
	b := []int{1, 0, 0, 1, 1}
	k := 1
	expect := 21
	runSample(t, a, b, k, expect)
}

func TestSample6(t *testing.T) {
	a := []int{1, 1, 2, 5}
	b := []int{1, 1, 0, 0}
	k := 4
	expect := 8
	runSample(t, a, b, k, expect)
}

func TestSample7(t *testing.T) {
	a := []int{1000000000, 1000000000}
	b := []int{1, 1}
	k := 1000000000
	expect := 3000000000
	runSample(t, a, b, k, expect)
}

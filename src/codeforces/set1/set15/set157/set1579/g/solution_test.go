package main

import "testing"

func runSample(t *testing.T, a []int, expect int) {
	res := solve(a)

	if res != expect {
		t.Fatalf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	a := []int{1, 3}
	expect := 3
	runSample(t, a, expect)
}

func TestSample2(t *testing.T) {
	a := []int{1, 2, 3}
	expect := 3
	runSample(t, a, expect)
}

func TestSample3(t *testing.T) {
	a := []int{6, 2, 3, 9}
	expect := 9
	runSample(t, a, expect)
}

func TestSample4(t *testing.T) {
	a := []int{6, 8, 4, 5}
	expect := 9
	runSample(t, a, expect)
}

func TestSample5(t *testing.T) {
	a := []int{1, 2, 4, 6, 7, 7, 3}
	expect := 7
	runSample(t, a, expect)
}

func TestSample6(t *testing.T) {
	a := []int{8, 6, 5, 1, 2, 2, 3, 6}
	expect := 8
	runSample(t, a, expect)
}

package main

import "testing"

func runSample(t *testing.T, a []int, expect int) {
	res := solve(a)

	if res != expect {
		t.Fatalf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	a := []int{1}
	expect := 1
	runSample(t, a, expect)
}

func TestSample2(t *testing.T) {
	a := []int{2, 4}
	expect := 3
	runSample(t, a, expect)
}

func TestSample3(t *testing.T) {
	a := []int{2, 3}
	expect := 2
	runSample(t, a, expect)
}

func TestSample4(t *testing.T) {
	a := []int{2, 4, 7, 14}
	expect := 7
	runSample(t, a, expect)
}

func TestSample5(t *testing.T) {
	a := []int{16, 5, 18, 7, 7, 12, 14}
	expect := 10
	runSample(t, a, expect)
}

func TestSample6(t *testing.T) {
	a := []int{16, 14, 2, 6, 16, 2}
	expect := 19
	runSample(t, a, expect)
}

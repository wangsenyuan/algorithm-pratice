package main

import "testing"

func runSample(t *testing.T, a []int, expect int) {
	res := solve(a)

	if res != expect {
		t.Fatalf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	a := []int{1, 5, 2, 4, 6}
	expect := 3
	runSample(t, a, expect)
}

func TestSample2(t *testing.T) {
	a := []int{8, 2, 5, 10}
	expect := 3
	runSample(t, a, expect)
}

func TestSample3(t *testing.T) {
	a := []int{1000, 2000}
	expect := 2
	runSample(t, a, expect)
}

func TestSample4(t *testing.T) {
	a := []int{465, 55, 3, 54, 234, 12, 45, 78}
	expect := 6
	runSample(t, a, expect)
}

func TestSample5(t *testing.T) {
	a := []int{16, 15, 17, 8, 30, 23, 20, 28, 27, 6, 1, 18, 24, 2, 10, 5, 14, 29, 12, 7}
	expect := 4
	runSample(t, a, expect)
}

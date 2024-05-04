package main

import "testing"

func runSample(t *testing.T, d int, a []int, expect int) {
	res := solve(d, a)

	if res != expect {
		t.Fatalf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	d := 10
	a := []int{-1, 5, 0, -5, 3}
	expect := 0
	runSample(t, d, a, expect)
}

func TestSample2(t *testing.T) {
	d := 4
	a := []int{-10, 0, 20}
	expect := -1
	runSample(t, d, a, expect)
}

func TestSample3(t *testing.T) {
	d := 10
	a := []int{-5, 0, 10, -11, 0}
	expect := 2
	runSample(t, d, a, expect)
}

func TestSample4(t *testing.T) {
	d := 78701
	a := []int{1, 0, -1, 0, -1, -1, 0, 1, 0, -1, 1, 1, -1, 1, 0, 0, -1, 0, 0}
	expect := 1
	runSample(t, d, a, expect)
}

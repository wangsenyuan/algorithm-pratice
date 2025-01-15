package main

import "testing"

func runSample(t *testing.T, a []int, expect int) {
	res := solve(a)

	if res != expect {
		t.Fatalf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	a := []int{1, 2, 2, 4}
	expect := 2
	runSample(t, a, expect)
}

func TestSample2(t *testing.T) {
	a := []int{1, 2, 4, 4, 4, 9}
	expect := 2
	runSample(t, a, expect)
}

func TestSample3(t *testing.T) {
	a := []int{1, 2, 4, 4, 4, 9, 4}
	expect := 2
	runSample(t, a, expect)
}

func TestSample4(t *testing.T) {
	a := []int{2, 1, 1, 1, 2, 2, 2, 1, 1, 2}
	expect := 10
	runSample(t, a, expect)
}

func TestSample5(t *testing.T) {
	a := []int{50, 50, 50, 50, 46, 46, 21, 62, 48, 43, 43, 80, 80, 25, 53, 98, 98, 92, 44, 74, 74, 90, 62, 71, 71, 95, 91, 96, 96, 65, 65, 73, 73, 94, 66, 66, 89, 89, 68, 69, 69, 93, 93, 68, 94, 99, 99, 95, 87, 87}
	expect := 3
	runSample(t, a, expect)
}

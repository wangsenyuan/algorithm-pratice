package main

import "testing"

func runSample(t *testing.T, a []int, expect int) {
	res := solve(a)

	if res != expect {
		t.Fatalf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	a := []int{0}
	expect := 0
	runSample(t, a, expect)
}

func TestSample2(t *testing.T) {
	a := []int{2, 4, 4, 2}
	expect := 3
	runSample(t, a, expect)
}

func TestSample3(t *testing.T) {
	a := []int{3, 2, 1, 0}
	expect := 2
	runSample(t, a, expect)
}

func TestSample4(t *testing.T) {
	a := []int{0, 3, 0}
	expect := 1
	runSample(t, a, expect)
}

func TestSample5(t *testing.T) {
	a := []int{0, 1, 3}
	expect := 2
	runSample(t, a, expect)
}

func TestSample6(t *testing.T) {
	a := []int{3, 1, 0, 3}
	expect := 3
	runSample(t, a, expect)
}

func TestSample7(t *testing.T) {
	a := []int{0, 2, 2, 2}
	expect := 2
	runSample(t, a, expect)
}

func TestSample8(t *testing.T) {
	a := []int{1, 3, 4, 2, 0, 4}
	expect := 4
	runSample(t, a, expect)
}

func TestSample9(t *testing.T) {
	a := []int{2, 2, 5, 2, 3, 4, 2, 4}
	expect := 6
	runSample(t, a, expect)
}

func TestSample10(t *testing.T) {
	a := []int{1, 3, 1}
	expect := 3
	runSample(t, a, expect)
}

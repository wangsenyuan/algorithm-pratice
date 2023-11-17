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
	expect := 2
	runSample(t, a, expect)
}

func TestSample2(t *testing.T) {
	a := []int{-1, 0}
	expect := 8
	runSample(t, a, expect)
}

func TestSample3(t *testing.T) {
	a := []int{1, -1}
	expect := 6
	runSample(t, a, expect)
}

func TestSample4(t *testing.T) {
	a := []int{1, 1}
	expect := 7
	runSample(t, a, expect)
}

func TestSample5(t *testing.T) {
	a := []int{-1, -2, -1}
	expect := 20
	runSample(t, a, expect)
}

func TestSample6(t *testing.T) {
	a := []int{1, -2, -1}
	expect := 17
	runSample(t, a, expect)
}

func TestSample7(t *testing.T) {
	a := []int{-1, 4, -2, 1}
	expect := 34
	runSample(t, a, expect)
}

func TestSample8(t *testing.T) {
	a := []int{1, 1, 1, 1, -4}
	expect := 30
	runSample(t, a, expect)
}

func TestSample9(t *testing.T) {
	a := []int{1, 1, 1, 1, 1}
	expect := 40
	runSample(t, a, expect)
}

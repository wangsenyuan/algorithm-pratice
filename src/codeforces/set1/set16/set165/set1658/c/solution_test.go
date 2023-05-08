package main

import "testing"

func runSample(t *testing.T, C []int, expect bool) {
	res := solve(C)

	if res != expect {
		t.Fatalf("Sample expect %t, but got %t", expect, res)
	}
}

func TestSample1(t *testing.T) {
	C := []int{1}
	expect := true
	runSample(t, C, expect)
}

func TestSample2(t *testing.T) {
	C := []int{1, 2}
	expect := true
	runSample(t, C, expect)
}

func TestSample3(t *testing.T) {
	C := []int{2, 2}
	expect := false
	runSample(t, C, expect)
}

func TestSample4(t *testing.T) {
	C := []int{1, 2, 4, 6, 3, 5}
	expect := false
	runSample(t, C, expect)
}

func TestSample5(t *testing.T) {
	C := []int{2, 3, 1, 2, 3, 4}
	expect := true
	runSample(t, C, expect)
}

func TestSample6(t *testing.T) {
	C := []int{3, 2, 1}
	expect := false
	runSample(t, C, expect)
}

func TestSample7(t *testing.T) {
	// 1, 2, 3   3
	// 3, 1, 2   1
	// 2, 3, 1   2
	C := []int{3, 1, 2}
	expect := true
	runSample(t, C, expect)
}
